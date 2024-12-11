package testworkflowexecutor

import (
	"fmt"
	"maps"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	testworkflowsv1 "github.com/kubeshop/testkube-operator/api/testworkflows/v1"
	testworkflowsclientv1 "github.com/kubeshop/testkube-operator/pkg/client/testworkflows/v1"
	"github.com/kubeshop/testkube/pkg/cloud"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowresolver"
)

type testWorkflowFetcher struct {
	client           testworkflowsclientv1.Interface
	cache            map[string]*testworkflowsv1.TestWorkflow
	prefetchedLabels []map[string]string
}

func NewTestWorkflowFetcher(
	client testworkflowsclientv1.Interface,
) *testWorkflowFetcher {
	return &testWorkflowFetcher{
		client: client,
		cache:  make(map[string]*testworkflowsv1.TestWorkflow),
	}
}

func (r *testWorkflowFetcher) PrefetchByLabelSelector(labels map[string]string) error {
	if containsSameMap(r.prefetchedLabels, labels) {
		return nil
	}
	selectors := make([]string, 0, len(labels))
	for k := range labels {
		selectors = append(selectors, fmt.Sprintf("%s=%s", k, labels[k]))
	}
	workflows, err := r.client.List(strings.Join(selectors, ","))
	if err != nil {
		return errors.Wrapf(err, "cannot fetch Test Workflows by label selector: %v", labels)
	}
	for i := range workflows.Items {
		r.cache[workflows.Items[i].Name] = &workflows.Items[i]
	}
	r.prefetchedLabels = append(r.prefetchedLabels, labels)
	return nil
}

func (r *testWorkflowFetcher) PrefetchByName(name string) error {
	if _, ok := r.cache[name]; ok {
		return nil
	}
	workflow, err := r.client.Get(name)
	if err != nil {
		return errors.Wrapf(err, "cannot fetch Test Workflow by name: %s", name)
	}
	r.cache[name] = workflow
	return nil
}

func (r *testWorkflowFetcher) PrefetchMany(selectors []*cloud.ScheduleSelector) error {
	// Categorize selectors
	names := make(map[string]struct{})
	labels := make([]map[string]string, 0)
	for i := range selectors {
		if selectors[i].Name == "" {
			if !containsSameMap(labels, selectors[i].LabelSelector) {
				labels = append(labels, selectors[i].LabelSelector)
			}
		} else {
			names[selectors[i].Name] = struct{}{}
		}
	}

	// Fetch firstly by the label selector, as it is more likely to conflict with others
	g := errgroup.Group{}
	g.SetLimit(10)
	for i := range labels {
		func(m map[string]string) {
			g.Go(func() error {
				return r.PrefetchByLabelSelector(labels[i])
			})
		}(labels[i])
	}
	err := g.Wait()
	if err != nil {
		return err
	}

	// Fetch the rest by name
	g = errgroup.Group{}
	g.SetLimit(10)
	for name := range names {
		func(n string) {
			g.Go(func() error {
				return r.PrefetchByName(n)
			})
		}(name)
	}
	return g.Wait()
}

func (r *testWorkflowFetcher) GetByName(name string) (*testworkflowsv1.TestWorkflow, error) {
	if r.cache[name] == nil {
		err := r.PrefetchByName(name)
		if err != nil {
			return nil, err
		}
	}
	return r.cache[name], nil
}

func (r *testWorkflowFetcher) GetByLabelSelector(labels map[string]string) ([]*testworkflowsv1.TestWorkflow, error) {
	if !containsSameMap(r.prefetchedLabels, labels) {
		err := r.PrefetchByLabelSelector(labels)
		if err != nil {
			return nil, err
		}
	}
	result := make([]*testworkflowsv1.TestWorkflow, 0)
loop:
	for name := range r.cache {
		for k := range labels {
			if r.cache[name].Labels[k] != labels[k] {
				continue loop
			}
		}
		result = append(result, r.cache[name])
	}
	return result, nil
}

func (r *testWorkflowFetcher) Get(selector *cloud.ScheduleSelector) ([]*testworkflowsv1.TestWorkflow, error) {
	if selector.Name == "" {
		return r.GetByLabelSelector(selector.LabelSelector)
	}
	v, err := r.GetByName(selector.Name)
	if err != nil {
		return nil, err
	}
	return []*testworkflowsv1.TestWorkflow{v}, nil
}

func (r *testWorkflowFetcher) Names() []string {
	names := make([]string, 0, len(r.cache))
	for k := range r.cache {
		names = append(names, k)
	}
	return names
}

func (r *testWorkflowFetcher) TemplateNames() map[string]struct{} {
	result := make(map[string]struct{})
	for k := range r.cache {
		maps.Copy(result, testworkflowresolver.ListTemplates(r.cache[k]))
	}
	return result
}

func containsSameMap[T comparable, U comparable](s []map[T]U, v map[T]U) bool {
	for i := range s {
		if len(s[i]) != len(v) {
			continue
		}
		for k := range s[i] {
			if x, ok := v[k]; !ok || x != s[i][k] {
				return true
			}
		}
	}
	return false
}