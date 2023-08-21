package triggers

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"k8s.io/client-go/dynamic"

	"github.com/google/go-cmp/cmp"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	appsinformerv1 "k8s.io/client-go/informers/apps/v1"
	"time"

	coreinformerv1 "k8s.io/client-go/informers/core/v1"
	networkinginformerv1 "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/kubeshop/testkube-operator/pkg/clientset/versioned"
	testkubeinformerv1 "github.com/kubeshop/testkube-operator/pkg/informers/externalversions/tests/v1"
	testkubeinformerv3 "github.com/kubeshop/testkube-operator/pkg/informers/externalversions/tests/v3"

	"k8s.io/client-go/tools/cache"

	testsv3 "github.com/kubeshop/testkube-operator/apis/tests/v3"
	testsuitev3 "github.com/kubeshop/testkube-operator/apis/testsuite/v3"
	testtriggersv1 "github.com/kubeshop/testkube-operator/apis/testtriggers/v1"
	"github.com/kubeshop/testkube-operator/pkg/informers/externalversions"
	"github.com/kubeshop/testkube-operator/pkg/validation/tests/v1/testtrigger"
)

type DynamicInformersEntry struct {
	informer *informers.GenericInformer
	cancel   context.CancelFunc
}

type k8sInformers struct {
	customResourceInformers []informers.GenericInformer
	podInformers            []coreinformerv1.PodInformer
	deploymentInformers     []appsinformerv1.DeploymentInformer
	daemonsetInformers      []appsinformerv1.DaemonSetInformer
	statefulsetInformers    []appsinformerv1.StatefulSetInformer
	serviceInformers        []coreinformerv1.ServiceInformer
	ingressInformers        []networkinginformerv1.IngressInformer
	clusterEventInformers   []coreinformerv1.EventInformer
	configMapInformers      []coreinformerv1.ConfigMapInformer

	testTriggerInformer testkubeinformerv1.TestTriggerInformer
	testSuiteInformer   testkubeinformerv3.TestSuiteInformer
	testInformer        testkubeinformerv3.TestInformer
	dynamicInformer     *DynamicInformerManager
}

func newK8sInformers(clientset kubernetes.Interface, dynamicClientset dynamic.Interface, testKubeClientset versioned.Interface, testkubeNamespace string, watcherNamespaces []string, logger *zap.SugaredLogger) *k8sInformers {

	var k8sInformers k8sInformers
	if len(watcherNamespaces) == 0 {
		watcherNamespaces = append(watcherNamespaces, v1.NamespaceAll)
	}

	for _, namespace := range watcherNamespaces {
		f := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithNamespace(namespace))
		k8sInformers.podInformers = append(k8sInformers.podInformers, f.Core().V1().Pods())
		k8sInformers.deploymentInformers = append(k8sInformers.deploymentInformers, f.Apps().V1().Deployments())
		k8sInformers.daemonsetInformers = append(k8sInformers.daemonsetInformers, f.Apps().V1().DaemonSets())
		k8sInformers.statefulsetInformers = append(k8sInformers.statefulsetInformers, f.Apps().V1().StatefulSets())
		k8sInformers.serviceInformers = append(k8sInformers.serviceInformers, f.Core().V1().Services())
		k8sInformers.ingressInformers = append(k8sInformers.ingressInformers, f.Networking().V1().Ingresses())
		k8sInformers.clusterEventInformers = append(k8sInformers.clusterEventInformers, f.Core().V1().Events())
		k8sInformers.configMapInformers = append(k8sInformers.configMapInformers, f.Core().V1().ConfigMaps())

	}

	var testkubeInformerFactory externalversions.SharedInformerFactory
	testkubeInformerFactory = externalversions.NewSharedInformerFactoryWithOptions(
		testKubeClientset, 0, externalversions.WithNamespace(testkubeNamespace))
	k8sInformers.testTriggerInformer = testkubeInformerFactory.Tests().V1().TestTriggers()
	k8sInformers.testSuiteInformer = testkubeInformerFactory.Tests().V3().TestSuites()
	k8sInformers.testInformer = testkubeInformerFactory.Tests().V3().Tests()
	k8sInformers.dynamicInformer = NewDynamicInformerManager(dynamicClientset, logger)
	return &k8sInformers
}

func (s *Service) runWatcher(ctx context.Context, leaseChan chan bool) {
	running := false
	var stopChan chan struct{}

	for {
		select {
		case <-ctx.Done():
			s.logger.Infof("trigger service: stopping watcher component: context finished")
			if _, ok := <-stopChan; ok {
				close(stopChan)
			}
			return
		case leased := <-leaseChan:
			if !leased {
				if running {
					s.logger.Infof("trigger service: instance %s in cluster %s lost lease", s.identifier, s.clusterID)
					close(stopChan)
					s.informers = nil
					running = false
				}
			} else {
				if !running {
					s.logger.Infof("trigger service: instance %s in cluster %s acquired lease", s.identifier, s.clusterID)
					s.informers = newK8sInformers(s.clientset, s.dynamicClientset, s.testKubeClientset, s.testkubeNamespace, s.watcherNamespaces, s.logger)
					stopChan = make(chan struct{})
					//TODO: clear the initialization
					s.informers.dynamicInformer.setStopCh(stopChan)
					s.runInformers(ctx, stopChan)
					running = true
				}
			}
		}
	}
}

func (s *Service) runInformers(ctx context.Context, stop <-chan struct{}) {
	if s.informers == nil {
		s.logger.Errorf("trigger service: error running k8s informers: informers are nil")
		return
	}

	for i := range s.informers.podInformers {
		s.informers.podInformers[i].Informer().AddEventHandler(s.podEventHandler(ctx))
	}

	for i := range s.informers.deploymentInformers {
		s.informers.deploymentInformers[i].Informer().AddEventHandler(s.deploymentEventHandler(ctx))
	}

	for i := range s.informers.daemonsetInformers {
		s.informers.daemonsetInformers[i].Informer().AddEventHandler(s.daemonSetEventHandler(ctx))
	}

	for i := range s.informers.statefulsetInformers {
		s.informers.statefulsetInformers[i].Informer().AddEventHandler(s.statefulSetEventHandler(ctx))
	}

	for i := range s.informers.serviceInformers {
		s.informers.serviceInformers[i].Informer().AddEventHandler(s.serviceEventHandler(ctx))
	}

	for i := range s.informers.ingressInformers {
		s.informers.ingressInformers[i].Informer().AddEventHandler(s.ingressEventHandler(ctx))
	}

	for i := range s.informers.clusterEventInformers {
		s.informers.clusterEventInformers[i].Informer().AddEventHandler(s.clusterEventEventHandler(ctx))
	}

	for i := range s.informers.configMapInformers {
		s.informers.configMapInformers[i].Informer().AddEventHandler(s.configMapEventHandler(ctx))
	}

	s.informers.testTriggerInformer.Informer().AddEventHandler(s.testTriggerEventHandler(ctx, stop))
	s.informers.testSuiteInformer.Informer().AddEventHandler(s.testSuiteEventHandler())
	s.informers.testInformer.Informer().AddEventHandler(s.testEventHandler())

	s.logger.Debugf("trigger service: starting pod informers")
	for i := range s.informers.podInformers {
		go s.informers.podInformers[i].Informer().Run(stop)
	}

	s.logger.Debugf("trigger service: starting deployment informers")
	for i := range s.informers.deploymentInformers {
		go s.informers.deploymentInformers[i].Informer().Run(stop)
	}

	s.logger.Debugf("trigger service: starting daemonset informers")
	for i := range s.informers.daemonsetInformers {
		go s.informers.daemonsetInformers[i].Informer().Run(stop)
	}

	s.logger.Debugf("trigger service: starting statefulset informers")
	for i := range s.informers.statefulsetInformers {
		go s.informers.statefulsetInformers[i].Informer().Run(stop)
	}

	s.logger.Debugf("trigger service: starting service informers")
	for i := range s.informers.serviceInformers {
		go s.informers.serviceInformers[i].Informer().Run(stop)
	}

	s.logger.Debugf("trigger service: starting ingress informers")
	for i := range s.informers.ingressInformers {
		go s.informers.ingressInformers[i].Informer().Run(stop)
	}

	s.logger.Debugf("trigger service: starting cluster event informers")
	for i := range s.informers.clusterEventInformers {
		go s.informers.clusterEventInformers[i].Informer().Run(stop)
	}

	s.logger.Debugf("trigger service: starting config map informers")
	for i := range s.informers.configMapInformers {
		go s.informers.configMapInformers[i].Informer().Run(stop)
	}

	s.logger.Debugf("trigger service: starting test trigger informer")
	go s.informers.testTriggerInformer.Informer().Run(stop)
	s.logger.Debugf("trigger service: starting test suite informer")
	go s.informers.testSuiteInformer.Informer().Run(stop)
	s.logger.Debugf("trigger service: starting test informer")
	go s.informers.testInformer.Informer().Run(stop)
}

func (s *Service) podEventHandler(ctx context.Context) cache.ResourceEventHandlerFuncs {
	getConditions := func(object metav1.Object) func() ([]testtriggersv1.TestTriggerCondition, error) {
		return func() ([]testtriggersv1.TestTriggerCondition, error) {
			return getPodConditions(ctx, s.clientset, object)
		}
	}
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				s.logger.Errorf("failed to process create pod event due to it being an unexpected type, received type %+v", obj)
				return
			}
			if inPast(pod.CreationTimestamp.Time, s.watchFromDate) {
				s.logger.Debugf(
					"trigger service: watcher component: no-op create trigger: pod %s/%s was created in the past",
					pod.Namespace, pod.Name,
				)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: pod %s/%s created", pod.Namespace, pod.Name)
			event := newWatcherEvent(testtrigger.EventCreated, pod, testtrigger.ResourcePod, withConditionsGetter(getConditions(pod)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching create pod event: %v", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				s.logger.Errorf("failed to process delete pod event due to it being an unexpected type, received type %+v", obj)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: pod %s/%s deleted", pod.Namespace, pod.Name)
			event := newWatcherEvent(testtrigger.EventDeleted, pod, testtrigger.ResourcePod, withConditionsGetter(getConditions(pod)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching delete pod event: %v", err)
			}
		},
	}
}

func (s *Service) deploymentEventHandler(ctx context.Context) cache.ResourceEventHandlerFuncs {
	getConditions := func(object metav1.Object) func() ([]testtriggersv1.TestTriggerCondition, error) {
		return func() ([]testtriggersv1.TestTriggerCondition, error) {
			return getDeploymentConditions(ctx, s.clientset, object)
		}
	}
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			deployment, ok := obj.(*appsv1.Deployment)
			if !ok {
				s.logger.Errorf("failed to process create deployment event due to it being an unexpected type, received type %+v", obj)
				return
			}
			if inPast(deployment.CreationTimestamp.Time, s.watchFromDate) {
				s.logger.Debugf(
					"trigger service: watcher component: no-op create trigger: deployment %s/%s was created in the past",
					deployment.Namespace, deployment.Name,
				)
				return
			}
			s.logger.Debugf("emiting event: deployment %s/%s created", deployment.Namespace, deployment.Name)
			event := newWatcherEvent(testtrigger.EventCreated, deployment, testtrigger.ResourceDeployment, withConditionsGetter(getConditions(deployment)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching create deployment event: %v", err)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldDeployment, ok := oldObj.(*appsv1.Deployment)
			if !ok {
				s.logger.Errorf(
					"failed to process update deployment event for old object due to it being an unexpected type, received type %+v",
					oldDeployment,
				)
				return
			}
			newDeployment, ok := newObj.(*appsv1.Deployment)
			if !ok {
				s.logger.Errorf(
					"failed to process update deployment event for new object due to it being an unexpected type, received type %+v",
					newDeployment,
				)
				return
			}
			if cmp.Equal(oldDeployment.Spec, newDeployment.Spec) {
				s.logger.Debugf("trigger service: watcher component: no-op update trigger: deployment specs are equal")
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: emiting event: deployment %s/%s updated",
				newDeployment.Namespace, newDeployment.Name,
			)
			causes := diffDeployments(oldDeployment, newDeployment)
			event := newWatcherEvent(testtrigger.EventModified, newDeployment, testtrigger.ResourceDeployment, withCauses(causes), withConditionsGetter(getConditions(newDeployment)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching update deployment event: %v", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			deployment, ok := obj.(*appsv1.Deployment)
			if !ok {
				s.logger.Errorf("failed to process create deployment event due to it being an unexpected type, received type %+v", obj)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: deployment %s/%s deleted", deployment.Namespace, deployment.Name)
			event := newWatcherEvent(testtrigger.EventDeleted, deployment, testtrigger.ResourceDeployment, withConditionsGetter(getConditions(deployment)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching delete deployment event: %v", err)
			}
		},
	}
}

func (s *Service) statefulSetEventHandler(ctx context.Context) cache.ResourceEventHandlerFuncs {
	getConditions := func(object metav1.Object) func() ([]testtriggersv1.TestTriggerCondition, error) {
		return func() ([]testtriggersv1.TestTriggerCondition, error) {
			return getStatefulSetConditions(ctx, s.clientset, object)
		}
	}
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			statefulset, ok := obj.(*appsv1.StatefulSet)
			if !ok {
				s.logger.Errorf("failed to process create statefulset event due to it being an unexpected type, received type %+v", obj)
				return
			}
			if inPast(statefulset.CreationTimestamp.Time, s.watchFromDate) {
				s.logger.Debugf(
					"trigger service: watcher component: no-op create trigger: statefulset %s/%s was created in the past",
					statefulset.Namespace, statefulset.Name,
				)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: statefulset %s/%s created", statefulset.Namespace, statefulset.Name)
			event := newWatcherEvent(testtrigger.EventCreated, statefulset, testtrigger.ResourceStatefulSet, withConditionsGetter(getConditions(statefulset)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching create statefulset event: %v", err)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldStatefulSet, ok := oldObj.(*appsv1.StatefulSet)
			if !ok {
				s.logger.Errorf(
					"failed to process update statefulset event for old object due to it being an unexpected type, received type %+v",
					oldStatefulSet,
				)
				return
			}
			newStatefulSet, ok := newObj.(*appsv1.StatefulSet)
			if !ok {
				s.logger.Errorf(
					"failed to process update statefulset event for new object due to it being an unexpected type, received type %+v",
					newStatefulSet,
				)
				return
			}
			if cmp.Equal(oldStatefulSet.Spec, newStatefulSet.Spec) {
				s.logger.Debugf("trigger service: watcher component: no-op update trigger: statefulset specs are equal")
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: emiting event: statefulset %s/%s updated",
				newStatefulSet.Namespace, newStatefulSet.Name,
			)
			event := newWatcherEvent(testtrigger.EventModified, newStatefulSet, testtrigger.ResourceStatefulSet, withConditionsGetter(getConditions(newStatefulSet)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching update statefulset event: %v", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			statefulset, ok := obj.(*appsv1.StatefulSet)
			if !ok {
				s.logger.Errorf("failed to process delete statefulset event due to it being an unexpected type, received type %+v", obj)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: statefulset %s/%s deleted", statefulset.Namespace, statefulset.Name)
			event := newWatcherEvent(testtrigger.EventDeleted, statefulset, testtrigger.ResourceStatefulSet, withConditionsGetter(getConditions(statefulset)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching delete statefulset event: %v", err)
			}
		},
	}
}

func (s *Service) daemonSetEventHandler(ctx context.Context) cache.ResourceEventHandlerFuncs {
	getConditions := func(object metav1.Object) func() ([]testtriggersv1.TestTriggerCondition, error) {
		return func() ([]testtriggersv1.TestTriggerCondition, error) {
			return getDaemonSetConditions(ctx, s.clientset, object)
		}
	}
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			daemonset, ok := obj.(*appsv1.DaemonSet)
			if !ok {
				s.logger.Errorf("failed to process create daemonset event due to it being an unexpected type, received type %+v", obj)
				return
			}
			if inPast(daemonset.CreationTimestamp.Time, s.watchFromDate) {
				s.logger.Debugf(
					"trigger service: watcher component: no-op create trigger: daemonset %s/%s was created in the past",
					daemonset.Namespace, daemonset.Name,
				)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: daemonset %s/%s created", daemonset.Namespace, daemonset.Name)
			event := newWatcherEvent(testtrigger.EventCreated, daemonset, testtrigger.ResourceDaemonSet, withConditionsGetter(getConditions(daemonset)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching create daemonset event: %v", err)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldDaemonSet, ok := oldObj.(*appsv1.DaemonSet)
			if !ok {
				s.logger.Errorf(
					"failed to process update daemonset event for old object due to it being an unexpected type, received type %+v",
					oldDaemonSet,
				)
				return
			}
			newDaemonSet, ok := newObj.(*appsv1.DaemonSet)
			if !ok {
				s.logger.Errorf(
					"failed to process update daemonset event for new object due to it being an unexpected type, received type %+v",
					newDaemonSet,
				)
				return
			}
			if cmp.Equal(oldDaemonSet.Spec, newDaemonSet.Spec) {
				s.logger.Debugf("trigger service: watcher component: no-op update trigger: daemonset specs are equal")
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: emiting event: daemonset %s/%s updated",
				newDaemonSet.Namespace, newDaemonSet.Name,
			)
			event := newWatcherEvent(testtrigger.EventModified, newDaemonSet, testtrigger.ResourceDaemonSet, withConditionsGetter(getConditions(newDaemonSet)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching update daemonset event: %v", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			daemonset, ok := obj.(*appsv1.DaemonSet)
			if !ok {
				s.logger.Errorf("failed to process delete daemonset event due to it being an unexpected type, received type %+v", obj)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: daemonset %s/%s deleted", daemonset.Namespace, daemonset.Name)
			event := newWatcherEvent(testtrigger.EventDeleted, daemonset, testtrigger.ResourceDaemonSet, withConditionsGetter(getConditions(daemonset)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching delete daemonset event: %v", err)
			}
		},
	}
}

func (s *Service) serviceEventHandler(ctx context.Context) cache.ResourceEventHandlerFuncs {
	getConditions := func(object metav1.Object) func() ([]testtriggersv1.TestTriggerCondition, error) {
		return func() ([]testtriggersv1.TestTriggerCondition, error) {
			return getServiceConditions(ctx, s.clientset, object)
		}
	}
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			service, ok := obj.(*corev1.Service)
			if !ok {
				s.logger.Errorf("failed to process create service event due to it being an unexpected type, received type %+v", obj)
				return
			}
			if inPast(service.CreationTimestamp.Time, s.watchFromDate) {
				s.logger.Debugf(
					"trigger service: watcher component: no-op create trigger: service %s/%s was created in the past",
					service.Namespace, service.Name,
				)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: service %s/%s created", service.Namespace, service.Name)
			event := newWatcherEvent(testtrigger.EventCreated, service, testtrigger.ResourceService, withConditionsGetter(getConditions(service)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching create service event: %v", err)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldService, ok := oldObj.(*corev1.Service)
			if !ok {
				s.logger.Errorf(
					"failed to process update service event for old object due to it being an unexpected type, received type %+v",
					oldService,
				)
				return
			}
			newService, ok := newObj.(*corev1.Service)
			if !ok {
				s.logger.Errorf(
					"failed to process update service event for new object due to it being an unexpected type, received type %+v",
					newService,
				)
				return
			}
			if cmp.Equal(oldService.Spec, newService.Spec) {
				s.logger.Debugf("trigger service: watcher component: no-op update trigger: service specs are equal")
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: emiting event: service %s/%s updated",
				newService.Namespace, newService.Name,
			)
			event := newWatcherEvent(testtrigger.EventModified, newService, testtrigger.ResourceService, withConditionsGetter(getConditions(newService)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching update service event: %v", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			service, ok := obj.(*corev1.Service)
			if !ok {
				s.logger.Errorf("failed to process delete service event due to it being an unexpected type, received type %+v", obj)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: service %s/%s deleted", service.Namespace, service.Name)
			event := newWatcherEvent(testtrigger.EventDeleted, service, testtrigger.ResourceService, withConditionsGetter(getConditions(service)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching delete service event: %v", err)
			}
		},
	}
}

func (s *Service) ingressEventHandler(ctx context.Context) cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			ingress, ok := obj.(*networkingv1.Ingress)
			if !ok {
				s.logger.Errorf("failed to process create ingress event due to it being an unexpected type, received type %+v", obj)
				return
			}
			if inPast(ingress.CreationTimestamp.Time, s.watchFromDate) {
				s.logger.Debugf(
					"trigger service: watcher component: no-op create trigger: ingress %s/%s was created in the past",
					ingress.Namespace, ingress.Name,
				)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: ingress %s/%s created", ingress.Namespace, ingress.Name)
			event := newWatcherEvent(testtrigger.EventCreated, ingress, testtrigger.ResourceIngress)
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching create ingress event: %v", err)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldIngress, ok := oldObj.(*networkingv1.Ingress)
			if !ok {
				s.logger.Errorf(
					"failed to process update ingress event for old object due to it being an unexpected type, received type %+v",
					oldIngress,
				)
				return
			}
			newIngress, ok := newObj.(*networkingv1.Ingress)
			if !ok {
				s.logger.Errorf(
					"failed to process update ingress event for new object due to it being an unexpected type, received type %+v",
					newIngress,
				)
				return
			}
			if cmp.Equal(oldIngress.Spec, newIngress.Spec) {
				s.logger.Debugf("trigger service: watcher component: no-op update trigger: ingress specs are equal")
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: emiting event: ingress %s/%s updated",
				oldIngress.Namespace, newIngress.Name,
			)
			event := newWatcherEvent(testtrigger.EventModified, newIngress, testtrigger.ResourceIngress)
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching update ingress event: %v", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			ingress, ok := obj.(*networkingv1.Ingress)
			if !ok {
				s.logger.Errorf("failed to process delete ingress event due to it being an unexpected type, received type %+v", obj)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: ingress %s/%s deleted", ingress.Namespace, ingress.Name)
			event := newWatcherEvent(testtrigger.EventDeleted, ingress, testtrigger.ResourceIngress)
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching delete ingress event: %v", err)
			}
		},
	}
}

func (s *Service) clusterEventEventHandler(ctx context.Context) cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			clusterEvent, ok := obj.(*corev1.Event)
			if !ok {
				s.logger.Errorf("failed to process create cluster event event due to it being an unexpected type, received type %+v", obj)
				return
			}
			if inPast(clusterEvent.CreationTimestamp.Time, s.watchFromDate) {
				s.logger.Debugf(
					"trigger service: watcher component: no-op create trigger: cluster event %s/%s was created in the past",
					clusterEvent.Namespace, clusterEvent.Name,
				)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: cluster event %s/%s created", clusterEvent.Namespace, clusterEvent.Name)
			event := newWatcherEvent(testtrigger.EventCreated, clusterEvent, testtrigger.ResourceEvent)
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching create cluster event event: %v", err)
			}
		},
	}
}

type DynamicInformerManager struct {
	dynamicClient      dynamic.Interface
	logger             *zap.SugaredLogger
	stop               <-chan struct{}
	triggerCRInformers map[string]*DynamicInformersEntry
	crInformers        map[string]*DynamicInformersEntry
}

func NewDynamicInformerManager(client dynamic.Interface, logger *zap.SugaredLogger) *DynamicInformerManager {
	return &DynamicInformerManager{
		dynamicClient:      client,
		logger:             logger,
		triggerCRInformers: make(map[string]*DynamicInformersEntry),
		crInformers:        make(map[string]*DynamicInformersEntry),
	}
}

func (m *DynamicInformerManager) NewInformer(t *testtriggersv1.TestTrigger, addHandler cache.ResourceEventHandlerFuncs, modifyHandler cache.ResourceEventHandlerFuncs, deleteHandler cache.ResourceEventHandlerFuncs) {
	if _, found := m.triggerCRInformers[t.Name+t.Namespace]; found {
		m.logger.Debugf("trigger service: informer already exist for %s resource in namespace %s", t.Name, t.Namespace)
		return
	}

	gvr := schema.GroupVersionResource{
		Group:    t.Spec.CustomResource.Group,
		Version:  t.Spec.CustomResource.Version,
		Resource: t.Spec.CustomResource.Resource,
	}

	df := dynamicinformer.NewFilteredDynamicSharedInformerFactory(m.dynamicClient, 0, t.Spec.ResourceSelector.Namespace, nil)
	customResourceInformer := df.ForResource(gvr)
	if string(t.Spec.Event) == string(testtrigger.EventCreated) {
		customResourceInformer.Informer().AddEventHandler(addHandler)
	}
	if string(t.Spec.Event) == string(testtrigger.EventModified) {
		customResourceInformer.Informer().AddEventHandler(modifyHandler)
	}
	if string(t.Spec.Event) == string(testtrigger.EventDeleted) {
		customResourceInformer.Informer().AddEventHandler(deleteHandler)
	}

	stopCtx, cancel := context.WithCancel(context.Background())
	go m.WaitForTermination(cancel, stopCtx.Done())

	ifd := &DynamicInformersEntry{
		informer: &customResourceInformer,
		cancel:   cancel,
	}
	m.triggerCRInformers[t.Name+t.Namespace] = ifd
	m.crInformers[gvr.String()+serializeSelector(t.Spec.ResourceSelector)] = ifd
	go customResourceInformer.Informer().Run(stopCtx.Done())
	m.logger.Debugf("trigger service: started a new custom resource informers for %s resource and %s test trigger", gvr.String(), t.Name)
}

func (m *DynamicInformerManager) setStopCh(stopChan chan struct{}) {
	m.stop = stopChan
}

func (m *DynamicInformerManager) TearDownInformer(t *testtriggersv1.TestTrigger) {
	if ifd, found := m.triggerCRInformers[t.Name+t.Namespace]; found {
		if ifd.informer != nil && ifd.cancel != nil {
			ifd.cancel()
			ifd.informer = nil
			delete(m.triggerCRInformers, t.Name+t.Namespace)
			m.logger.Debugf("trigger service: tearing down for %v resource and %s test trigger", t.Spec.CustomResource, t.Name)
		}
		m.logger.Debugf("trigger service: ignoring tear down for %v resource and %s test trigger", t.Spec.CustomResource, t.Name)
	}
	m.logger.Debugf("trigger service: could not found infomer for %v resource and %s test trigger to teardown", t.Spec.CustomResource, t.Name)
}

func (s *Service) testTriggerEventHandler(ctx context.Context, stop <-chan struct{}) cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			t, ok := obj.(*testtriggersv1.TestTrigger)
			if !ok {
				s.logger.Errorf("failed to process create testtrigger event due to it being an unexpected type, received type %+v", obj)
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: adding testtrigger %s/%s for resource %s on event %s",
				t.Namespace, t.Name, t.Spec.Resource, t.Spec.Event,
			)
			s.addTrigger(t)

			if t.Spec.Resource == testtrigger.ResourceCustomResource {
				gvr := schema.GroupVersionResource{
					Group:    t.Spec.CustomResource.Group,
					Version:  t.Spec.CustomResource.Version,
					Resource: t.Spec.CustomResource.Resource,
				}
				s.informers.dynamicInformer.NewInformer(t,
					s.customResourceAddEventHandler(ctx, gvr),
					s.customResourceModifyEventHandler(ctx, gvr),
					s.customResourceDeleteEventHandler(ctx, gvr),
				)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			t, ok := newObj.(*testtriggersv1.TestTrigger)
			if !ok {
				s.logger.Errorf(
					"failed to process update testtrigger event for new testtrigger due to it being an unexpected type, received type %+v",
					newObj,
				)
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: updating testtrigger %s/%s for resource %s on event %s",
				t.Namespace, t.Name, t.Spec.Resource, t.Spec.Event,
			)
			s.updateTrigger(t)
		},
		DeleteFunc: func(obj interface{}) {
			t, ok := obj.(*testtriggersv1.TestTrigger)
			if !ok {
				s.logger.Errorf("failed to process delete testtrigger event due to it being an unexpected type, received type %+v", obj)
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: deleting testtrigger %s/%s for resource %s on event %s",
				t.Namespace, t.Name, t.Spec.Resource, t.Spec.Event,
			)
			s.removeTrigger(t)
			s.informers.dynamicInformer.TearDownInformer(t)
		},
	}
}

// TODO: fix the formation
func serializeSelector(selector testtriggersv1.TestTriggerSelector) string {
	return selector.Name + selector.Namespace
}

func (m *DynamicInformerManager) WaitForTermination(cancel context.CancelFunc, done <-chan struct{}) {
	for {
		select {
		case <-m.stop:
			// Context is done, so we should stop the worker
			m.logger.Debug("Work is stopped")
			cancel()
			return
		case <-done:
			// Context is done, so we should stop the worker
			m.logger.Debug("Work is done")
			return
		default:
			// Do some work
			//m.logger.Debug("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func (s *Service) customResourceAddEventHandler(ctx context.Context, gvr schema.GroupVersionResource) cache.ResourceEventHandlerFuncs {
	getConditions := func(object metav1.Object) func() ([]testtriggersv1.TestTriggerCondition, error) {
		return func() ([]testtriggersv1.TestTriggerCondition, error) {
			return getCustomResourceConditions(ctx, s.dynamicClientset, object, gvr)
		}
	}
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			customResource, ok := obj.(*unstructured.Unstructured)
			if !ok {
				s.logger.Errorf("failed to process create customResource event due to it being an unexpected type, received type %+v", obj)
				return
			}
			obsGeneration, _, err := unstructured.NestedInt64(customResource.Object, "status", "observedGeneration")
			if err != nil {
				s.logger.Debugf("trigger service: watcher component: error when unstructuring custom resource: %s/%s ", customResource.GetKind(), customResource.GetName())
				return
			}
			generation := customResource.GetGeneration()
			if generation == obsGeneration {
				s.logger.Debugf("trigger service: watcher component: no-op update trigger: %s/%s new and old specs are equal", customResource.GetKind(), customResource.GetName())
				return
			}
			s.logger.Debugf("trigger customResource: watcher component: emiting event: customResource %s/%s created", customResource.GetNamespace(), customResource.GetName())
			event := newWatcherEvent(testtrigger.EventCreated, customResource, testtrigger.ResourceCustomResource, withConditionsGetter(getConditions(customResource)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching create customResource event: %v", err)
			}
		},
	}
}

func (s *Service) customResourceModifyEventHandler(ctx context.Context, gvr schema.GroupVersionResource) cache.ResourceEventHandlerFuncs {
	getConditions := func(object metav1.Object) func() ([]testtriggersv1.TestTriggerCondition, error) {
		return func() ([]testtriggersv1.TestTriggerCondition, error) {
			return getCustomResourceConditions(ctx, s.dynamicClientset, object, gvr)
		}
	}
	return cache.ResourceEventHandlerFuncs{
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldCustomResource, ok := oldObj.(*unstructured.Unstructured)
			if !ok {
				s.logger.Errorf(
					"failed to process update service event for old object due to it being an unexpected type, received type %+v",
					oldCustomResource,
				)
				return
			}
			newCustomResource, ok := newObj.(*unstructured.Unstructured)
			if !ok {
				s.logger.Errorf(
					"failed to process update service event for new object due to it being an unexpected type, received type %+v",
					newCustomResource,
				)
				return
			}

			newObsGeneration, _, err := unstructured.NestedInt64(newCustomResource.Object, "status", "observedGeneration")
			if err != nil {
				s.logger.Debugf("trigger service: watcher component: error when unstructuring custom resource: %s/%s ", oldCustomResource.GetKind(), oldCustomResource.GetName())
				return
			}

			oldObsGeneration, _, err := unstructured.NestedInt64(oldCustomResource.Object, "status", "observedGeneration")
			if err != nil {
				s.logger.Debugf("trigger service: watcher component: error when unstructuring custom resource: %s/%s ", oldCustomResource.GetKind(), oldCustomResource.GetName())
				return
			}

			oldGeneration := oldCustomResource.GetGeneration()
			newGeneration := newCustomResource.GetGeneration()
			fmt.Printf("old gen: %v, old obs gen: %v, new gen: %v, new obs gen: %v\n", oldGeneration, oldObsGeneration, newGeneration, newObsGeneration)

			if newGeneration == newObsGeneration {
				s.logger.Debugf("trigger service: watcher component: no-op update trigger: %s/%s new and old specs are equal", oldCustomResource.GetKind(), oldCustomResource.GetName())
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: emiting event: service %s/%s updated",
				newCustomResource.GetNamespace(), newCustomResource.GetName(),
			)
			event := newWatcherEvent(testtrigger.EventModified, newCustomResource, testtrigger.ResourceCustomResource, withConditionsGetter(getConditions(newCustomResource)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching update service event: %v", err)
			}
		},
	}
}

func (s *Service) customResourceDeleteEventHandler(ctx context.Context, gvr schema.GroupVersionResource) cache.ResourceEventHandlerFuncs {
	getConditions := func(object metav1.Object) func() ([]testtriggersv1.TestTriggerCondition, error) {
		return func() ([]testtriggersv1.TestTriggerCondition, error) {
			return getCustomResourceConditions(ctx, s.dynamicClientset, object, gvr)
		}
	}
	return cache.ResourceEventHandlerFuncs{
		DeleteFunc: func(obj interface{}) {
			customResource, ok := obj.(*unstructured.Unstructured)
			if !ok {
				s.logger.Errorf("failed to process delete customResource event due to it being an unexpected type, received type %+v", obj)
				return
			}
			obsGeneration, _, err := unstructured.NestedInt64(customResource.Object, "status", "observedGeneration")
			if err != nil {
				s.logger.Debugf("trigger service: watcher component: error when unstructuring custom resource: %s/%s ", customResource.GetKind(), customResource.GetName())
				return
			}
			generation := customResource.GetGeneration()
			if generation == obsGeneration {
				s.logger.Debugf("trigger service: watcher component: no-op update trigger: %s/%s new and old specs are equal", customResource.GetKind(), customResource.GetName())
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: customResource %s/%s deleted", customResource.GetNamespace(), customResource.GetName())
			event := newWatcherEvent(testtrigger.EventDeleted, customResource, testtrigger.ResourceCustomResource, withConditionsGetter(getConditions(customResource)))
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching delete customResource event: %v", err)
			}
		},
	}
}

func (s *Service) testSuiteEventHandler() cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			testSuite, ok := obj.(*testsuitev3.TestSuite)
			if !ok {
				s.logger.Errorf("failed to process create testsuite event due to it being an unexpected type, received type %+v", obj)
				return
			}
			if inPast(testSuite.CreationTimestamp.Time, s.watchFromDate) {
				s.logger.Debugf(
					"trigger service: watcher component: no-op create test suite: test suite %s/%s was created in the past",
					testSuite.Namespace, testSuite.Name,
				)
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: adding testsuite %s/%s",
				testSuite.Namespace, testSuite.Name,
			)
			s.addTestSuite(testSuite)
		},
	}
}

func (s *Service) testEventHandler() cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			test, ok := obj.(*testsv3.Test)
			if !ok {
				s.logger.Errorf("failed to process create test event due to it being an unexpected type, received type %+v", obj)
				return
			}
			if inPast(test.CreationTimestamp.Time, s.watchFromDate) {
				s.logger.Debugf(
					"trigger service: watcher component: no-op create test: test %s/%s was created in the past",
					test.Namespace, test.Name,
				)
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: adding test %s/%s",
				test.Namespace, test.Name,
			)
			s.addTest(test)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			test, ok := newObj.(*testsv3.Test)
			if !ok {
				s.logger.Errorf("failed to process update test event due to it being an unexpected type, received type %+v", newObj)
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: updating test %s/%s",
				test.Namespace, test.Name,
			)
			s.updateTest(test)
		},
	}
}

func (s *Service) configMapEventHandler(ctx context.Context) cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			configMap, ok := obj.(*corev1.ConfigMap)
			if !ok {
				s.logger.Errorf("failed to process create config map event due to it being an unexpected type, received type %+v", obj)
				return
			}
			if inPast(configMap.CreationTimestamp.Time, s.watchFromDate) {
				s.logger.Debugf(
					"trigger service: watcher component: no-op create trigger: config map %s/%s was created in the past",
					configMap.Namespace, configMap.Name,
				)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: config map %s/%s created", configMap.Namespace, configMap.Name)
			event := newWatcherEvent(testtrigger.EventCreated, configMap, testtrigger.ResourceConfigMap)
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching create config map event: %v", err)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldConfigMap, ok := oldObj.(*corev1.ConfigMap)
			if !ok {
				s.logger.Errorf(
					"failed to process update config map event for old object due to it being an unexpected type, received type %+v",
					oldConfigMap,
				)
				return
			}
			newConfigMap, ok := newObj.(*corev1.ConfigMap)
			if !ok {
				s.logger.Errorf(
					"failed to process update config map event for new object due to it being an unexpected type, received type %+v",
					newConfigMap,
				)
				return
			}
			if cmp.Equal(oldConfigMap.Data, newConfigMap.Data) && cmp.Equal(oldConfigMap.BinaryData, newConfigMap.BinaryData) {
				s.logger.Debugf("trigger service: watcher component: no-op update trigger: config map data and binary data are equal")
				return
			}
			s.logger.Debugf(
				"trigger service: watcher component: emiting event: config map %s/%s updated",
				oldConfigMap.Namespace, newConfigMap.Name,
			)
			event := newWatcherEvent(testtrigger.EventModified, newConfigMap, testtrigger.ResourceConfigMap)
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching update config map event: %v", err)
			}
		},
		DeleteFunc: func(obj interface{}) {
			configMap, ok := obj.(*corev1.ConfigMap)
			if !ok {
				s.logger.Errorf("failed to process delete config map event due to it being an unexpected type, received type %+v", obj)
				return
			}
			s.logger.Debugf("trigger service: watcher component: emiting event: config map %s/%s deleted", configMap.Namespace, configMap.Name)
			event := newWatcherEvent(testtrigger.EventDeleted, configMap, testtrigger.ResourceConfigMap)
			if err := s.match(ctx, event); err != nil {
				s.logger.Errorf("event matcher returned an error while matching delete config map event: %v", err)
			}
		},
	}
}
