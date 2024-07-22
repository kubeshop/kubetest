package presets

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	testworkflowsv1 "github.com/kubeshop/testkube-operator/api/testworkflows/v1"
	constants2 "github.com/kubeshop/testkube/cmd/testworkflow-init/constants"
	"github.com/kubeshop/testkube/internal/common"
	"github.com/kubeshop/testkube/pkg/expressions"
	"github.com/kubeshop/testkube/pkg/imageinspector"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowprocessor/action/actiontypes"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowprocessor/action/actiontypes/lite"
	"github.com/kubeshop/testkube/pkg/testworkflows/testworkflowprocessor/constants"
)

type dummyInspector struct{}

func (*dummyInspector) Inspect(ctx context.Context, registry, image string, pullPolicy corev1.PullPolicy, pullSecretNames []string) (*imageinspector.Info, error) {
	return &imageinspector.Info{}, nil
}

func (*dummyInspector) ResolveName(registry, image string) string {
	return image
}

var (
	ins         = &dummyInspector{}
	proc        = NewPro(ins)
	execMachine = expressions.NewMachine().
			Register("resource.root", "dummy-id").
			Register("resource.id", "dummy-id-abc")
	envInstructions = lite.EnvVarFrom("01", false, constants2.EnvInstructions, corev1.EnvVarSource{
		FieldRef: &corev1.ObjectFieldSelector{FieldPath: constants.SpecAnnotationFieldPath},
	})
	envDebugNode = lite.EnvVarFrom("00", false, constants2.EnvNodeName, corev1.EnvVarSource{
		FieldRef: &corev1.ObjectFieldSelector{FieldPath: "spec.nodeName"},
	})
	envDebugPod = lite.EnvVarFrom("00", false, constants2.EnvPodName, corev1.EnvVarSource{
		FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.name"},
	})
	envDebugNamespace = lite.EnvVarFrom("00", false, constants2.EnvNamespaceName, corev1.EnvVarSource{
		FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.namespace"},
	})
	envDebugServiceAccount = lite.EnvVarFrom("00", false, constants2.EnvServiceAccountName, corev1.EnvVarSource{
		FieldRef: &corev1.ObjectFieldSelector{FieldPath: "spec.serviceAccountName"},
	})
)

func env(index int, computed bool, name, value string) corev1.EnvVar {
	return lite.EnvVar(fmt.Sprintf("%d", index), computed, name, value)
}

func cmd(values ...string) *[]string {
	return &values
}

func cmdShell(shell string) *[]string {
	args := []string{"-c", "set -e\n" + shell}
	return &args
}

func getSpec(actions actiontypes.ActionGroups) string {
	v, err := json.Marshal(actions)
	if err != nil {
		panic(err)
	}
	return string(v)
}

func TestProcessEmpty(t *testing.T) {
	wf := &testworkflowsv1.TestWorkflow{}

	_, err := proc.Bundle(context.Background(), wf, execMachine)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "has nothing to run")
}

func TestProcessBasic(t *testing.T) {
	wf := &testworkflowsv1.TestWorkflow{
		Spec: testworkflowsv1.TestWorkflowSpec{
			Steps: []testworkflowsv1.Step{
				{StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test"}},
			},
		},
	}

	res, err := proc.Bundle(context.Background(), wf, execMachine)
	assert.NoError(t, err)

	sig := res.Signature
	sigSerialized, _ := json.Marshal(sig)

	volumes := res.Job.Spec.Template.Spec.Volumes
	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts

	wantInstructions := actiontypes.NewActionGroups().
		Append(func(list actiontypes.ActionList) actiontypes.ActionList {
			return list.
				Setup(true, true).
				Declare(constants.RootOperationName, "true").
				Declare(sig[0].Ref(), "true", constants.RootOperationName).
				Result(constants.RootOperationName, sig[0].Ref()).
				Result("", constants.RootOperationName).
				Start("").
				CurrentStatus("true").
				Start(constants.RootOperationName).
				CurrentStatus(constants.RootOperationName)
		}).
		Append(func(list actiontypes.ActionList) actiontypes.ActionList {
			return list.
				MutateContainer(sig[0].Ref(), testworkflowsv1.ContainerConfig{
					Command: cmd("/.tktw/bin/sh"),
					Args:    cmdShell("shell-test"),
				}).
				Start(sig[0].Ref()).
				Execute(sig[0].Ref(), false).
				End(sig[0].Ref()).
				End(constants.RootOperationName).
				End("")
		})

	want := batchv1.Job{
		TypeMeta: metav1.TypeMeta{Kind: "Job", APIVersion: "batch/v1"},
		ObjectMeta: metav1.ObjectMeta{
			Name: "dummy-id-abc",
			Labels: map[string]string{
				constants.ResourceIdLabelName:     "dummy-id-abc",
				constants.RootResourceIdLabelName: "dummy-id",
			},
			Annotations: map[string]string{
				constants.SignatureAnnotationName: string(sigSerialized),
			},
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: common.Ptr(int32(0)),
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						constants.ResourceIdLabelName:     "dummy-id-abc",
						constants.RootResourceIdLabelName: "dummy-id",
					},
					Annotations: map[string]string{
						constants.SpecAnnotationName: getSpec(wantInstructions),
					},
				},
				Spec: corev1.PodSpec{
					RestartPolicy:      corev1.RestartPolicyNever,
					EnableServiceLinks: common.Ptr(false),
					Volumes:            volumes,
					InitContainers: []corev1.Container{
						{
							Name:            "1",
							Image:           constants.DefaultInitImage,
							ImagePullPolicy: corev1.PullIfNotPresent,
							Command:         []string{"/init", "0"},
							Env: []corev1.EnvVar{
								envDebugNode,
								envDebugPod,
								envDebugNamespace,
								envDebugServiceAccount,
								envInstructions,
							},
							VolumeMounts: volumeMounts,
							SecurityContext: &corev1.SecurityContext{
								RunAsGroup: common.Ptr(constants.DefaultFsGroup),
							},
						},
					},
					Containers: []corev1.Container{
						{
							Name:            "2",
							Image:           constants.DefaultInitImage,
							ImagePullPolicy: corev1.PullIfNotPresent,
							Command:         []string{"/init", "1"},
							Env: []corev1.EnvVar{
								env(0, false, "CI", "1"),
							},
							VolumeMounts: volumeMounts,
							SecurityContext: &corev1.SecurityContext{
								RunAsGroup: common.Ptr(constants.DefaultFsGroup),
							},
						},
					},
					SecurityContext: &corev1.PodSecurityContext{
						FSGroup: common.Ptr(constants.DefaultFsGroup),
					},
				},
			},
		},
	}

	assert.Equal(t, want, res.Job)

	assert.Equal(t, 2, len(volumeMounts))
	assert.Equal(t, 2, len(volumes))
	assert.Equal(t, constants.DefaultInternalPath, volumeMounts[0].MountPath)
	assert.Equal(t, constants.DefaultDataPath, volumeMounts[1].MountPath)
	assert.True(t, volumeMounts[0].Name == volumes[0].Name)
	assert.True(t, volumeMounts[1].Name == volumes[1].Name)
}

func TestProcessBasicEnvReference(t *testing.T) {
	wf := &testworkflowsv1.TestWorkflow{
		Spec: testworkflowsv1.TestWorkflowSpec{
			Steps: []testworkflowsv1.Step{
				{StepDefaults: testworkflowsv1.StepDefaults{
					Container: &testworkflowsv1.ContainerConfig{
						Env: []corev1.EnvVar{
							{Name: "ZERO", Value: "foo"},
							{Name: "UNDETERMINED", Value: "{{call(abc)}}xxx"},
							{Name: "INPUT", Value: "{{env.ZERO}}bar"},
							{Name: "NEXT", Value: "foo{{env.UNDETERMINED}}{{env.LAST}}"},
							{Name: "LAST", Value: "foo{{env.INPUT}}bar"},
						},
					},
				}, StepOperations: testworkflowsv1.StepOperations{
					Shell: "shell-test",
				}},
			},
		},
	}

	res, err := proc.Bundle(context.Background(), wf, execMachine)
	sig := res.Signature

	volumes := res.Job.Spec.Template.Spec.Volumes
	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts

	wantInstructions := lite.NewLiteActionGroups().
		Append(func(list lite.LiteActionList) lite.LiteActionList {
			return list.
				Setup(true, true).
				Declare(constants.RootOperationName, "true").
				Declare(sig[0].Ref(), "true", constants.RootOperationName).
				Result(constants.RootOperationName, sig[0].Ref()).
				Result("", constants.RootOperationName).
				Start("").
				CurrentStatus("true").
				Start(constants.RootOperationName).
				CurrentStatus(constants.RootOperationName)
		}).
		Append(func(list lite.LiteActionList) lite.LiteActionList {
			return list.
				MutateContainer(lite.LiteContainerConfig{
					Command: cmd("/.tktw/bin/sh"),
					Args:    cmdShell("shell-test"),
				}).
				Start(sig[0].Ref()).
				Execute(sig[0].Ref(), false).
				End(sig[0].Ref()).
				End(constants.RootOperationName).
				End("")
		})

	wantPod := corev1.PodSpec{
		RestartPolicy:      corev1.RestartPolicyNever,
		EnableServiceLinks: common.Ptr(false),
		Volumes:            volumes,
		InitContainers: []corev1.Container{
			{
				Name:            "1",
				Image:           constants.DefaultInitImage,
				ImagePullPolicy: corev1.PullIfNotPresent,
				Command:         []string{"/init", "0"},
				Env: []corev1.EnvVar{
					envDebugNode,
					envDebugPod,
					envDebugNamespace,
					envDebugServiceAccount,
					envInstructions,
				},
				VolumeMounts: volumeMounts,
				SecurityContext: &corev1.SecurityContext{
					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
				},
			},
		},
		Containers: []corev1.Container{
			{
				Name:            "2",
				Image:           constants.DefaultInitImage,
				ImagePullPolicy: corev1.PullIfNotPresent,
				Command:         []string{"/init", "1"},
				Env: []corev1.EnvVar{
					env(0, false, "CI", "1"),
					env(0, false, "ZERO", "foo"),
					env(0, true, "UNDETERMINED", "{{call(abc)}}xxx"),
					env(0, false, "INPUT", "foobar"),
					env(0, true, "NEXT", "foo{{env.UNDETERMINED}}foofoobarbar"),
					env(0, false, "LAST", "foofoobarbar"),
				},
				VolumeMounts: volumeMounts,
				SecurityContext: &corev1.SecurityContext{
					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
				},
			},
		},
		SecurityContext: &corev1.PodSecurityContext{
			FSGroup: common.Ptr(constants.DefaultFsGroup),
		},
	}

	var gotInstructions lite.LiteActionGroups
	instructionsErr := json.Unmarshal([]byte(res.Job.Spec.Template.Annotations[constants.SpecAnnotationName]), &gotInstructions)

	assert.NoError(t, instructionsErr)
	assert.Equal(t, wantInstructions, gotInstructions)

	assert.NoError(t, err)
	assert.Equal(t, wantPod, res.Job.Spec.Template.Spec)
}

//func TestProcessMultipleSteps(t *testing.T) {
//	wf := &testworkflowsv1.TestWorkflow{
//		Spec: testworkflowsv1.TestWorkflowSpec{
//			Steps: []testworkflowsv1.Step{
//				{StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test"}},
//				{StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-2"}},
//			},
//		},
//	}
//
//	res, err := proc.Bundle(context.Background(), wf, execMachine)
//	sig := res.Signature
//
//	volumes := res.Job.Spec.Template.Spec.Volumes
//	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts
//
//	want := corev1.PodSpec{
//		RestartPolicy:      corev1.RestartPolicyNever,
//		EnableServiceLinks: common.Ptr(false),
//		Volumes:            volumes,
//		InitContainers: []corev1.Container{
//			{
//				Name:            "tktw-init",
//				Image:           constants.DefaultInitImage,
//				ImagePullPolicy: corev1.PullIfNotPresent,
//				Command:         []string{"/bin/sh", "-c"},
//				Env:             initEnvs,
//				VolumeMounts:    volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[0].Ref(),
//					"-c", fmt.Sprintf("%s,%s=passed", sig[0].Ref(), sig[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s", sig[0].Ref(), sig[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		Containers: []corev1.Container{
//			{
//				Name:            sig[1].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Ref(),
//					"-c", fmt.Sprintf("%s=passed", sig[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s", sig[0].Ref(), sig[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-2"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		SecurityContext: &corev1.PodSecurityContext{
//			FSGroup: common.Ptr(constants.DefaultFsGroup),
//		},
//	}
//
//	assert.NoError(t, err)
//	assert.Equal(t, want, res.Job.Spec.Template.Spec)
//}
//
//func TestProcessNestedSteps(t *testing.T) {
//	wf := &testworkflowsv1.TestWorkflow{
//		Spec: testworkflowsv1.TestWorkflowSpec{
//			Steps: []testworkflowsv1.Step{
//				{StepMeta: testworkflowsv1.StepMeta{Name: "A"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test"}},
//				{
//					StepMeta: testworkflowsv1.StepMeta{Name: "B"},
//					Steps: []testworkflowsv1.Step{
//						{StepMeta: testworkflowsv1.StepMeta{Name: "C"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-2"}},
//						{StepMeta: testworkflowsv1.StepMeta{Name: "D"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-3"}},
//					},
//				},
//				{StepMeta: testworkflowsv1.StepMeta{Name: "E"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-4"}},
//			},
//		},
//	}
//
//	res, err := proc.Bundle(context.Background(), wf, execMachine)
//	sig := res.Signature
//
//	volumes := res.Job.Spec.Template.Spec.Volumes
//	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts
//
//	want := corev1.PodSpec{
//		RestartPolicy:      corev1.RestartPolicyNever,
//		EnableServiceLinks: common.Ptr(false),
//		Volumes:            volumes,
//		InitContainers: []corev1.Container{
//			{
//				Name:            "tktw-init",
//				Image:           constants.DefaultInitImage,
//				ImagePullPolicy: corev1.PullIfNotPresent,
//				Command:         []string{"/bin/sh", "-c"},
//				Env:             initEnvs,
//				VolumeMounts:    volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[0].Ref(),
//					"-c", fmt.Sprintf("%s,%s,%s,%s=passed", sig[0].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref(), sig[2].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s&&%s", sig[0].Ref(), sig[1].Ref(), sig[2].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[1].Children()[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Children()[0].Ref(),
//					"-i", fmt.Sprintf("%s", sig[1].Ref()),
//					"-c", fmt.Sprintf("%s,%s,%s=passed", sig[1].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s&&%s", sig[0].Ref(), sig[1].Ref(), sig[2].Ref()),
//					"-r", fmt.Sprintf("%s=%s&&%s", sig[1].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-2"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[1].Children()[1].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Children()[1].Ref(),
//					"-i", fmt.Sprintf("%s", sig[1].Ref()),
//					"-c", fmt.Sprintf("%s=passed", sig[1].Children()[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s&&%s", sig[0].Ref(), sig[1].Ref(), sig[2].Ref()),
//					"-r", fmt.Sprintf("%s=%s&&%s", sig[1].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-3"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		Containers: []corev1.Container{
//			{
//				Name:            sig[2].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[2].Ref(),
//					"-c", fmt.Sprintf("%s=passed", sig[2].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s&&%s", sig[0].Ref(), sig[1].Ref(), sig[2].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-4"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		SecurityContext: &corev1.PodSecurityContext{
//			FSGroup: common.Ptr(constants.DefaultFsGroup),
//		},
//	}
//
//	assert.NoError(t, err)
//	assert.Equal(t, want, res.Job.Spec.Template.Spec)
//}
//
//func TestProcessOptionalSteps(t *testing.T) {
//	wf := &testworkflowsv1.TestWorkflow{
//		Spec: testworkflowsv1.TestWorkflowSpec{
//			Steps: []testworkflowsv1.Step{
//				{StepMeta: testworkflowsv1.StepMeta{Name: "A"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test"}},
//				{
//					StepMeta:    testworkflowsv1.StepMeta{Name: "B"},
//					StepControl: testworkflowsv1.StepControl{Optional: true},
//					Steps: []testworkflowsv1.Step{
//						{StepMeta: testworkflowsv1.StepMeta{Name: "C"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-2"}},
//						{StepMeta: testworkflowsv1.StepMeta{Name: "D"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-3"}},
//					},
//				},
//				{StepMeta: testworkflowsv1.StepMeta{Name: "E"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-4"}},
//			},
//		},
//	}
//
//	res, err := proc.Bundle(context.Background(), wf, execMachine)
//	sig := res.Signature
//
//	volumes := res.Job.Spec.Template.Spec.Volumes
//	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts
//
//	want := corev1.PodSpec{
//		RestartPolicy:      corev1.RestartPolicyNever,
//		EnableServiceLinks: common.Ptr(false),
//		Volumes:            volumes,
//		InitContainers: []corev1.Container{
//			{
//				Name:            "tktw-init",
//				Image:           constants.DefaultInitImage,
//				ImagePullPolicy: corev1.PullIfNotPresent,
//				Command:         []string{"/bin/sh", "-c"},
//				Env:             initEnvs,
//				VolumeMounts:    volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[0].Ref(),
//					"-c", fmt.Sprintf("%s,%s,%s,%s=passed", sig[0].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref(), sig[2].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s", sig[0].Ref(), sig[2].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[1].Children()[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Children()[0].Ref(),
//					"-i", fmt.Sprintf("%s", sig[1].Ref()),
//					"-c", fmt.Sprintf("%s,%s,%s=passed", sig[1].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref()),
//					"-r", fmt.Sprintf("%s=%s&&%s", sig[1].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-2"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[1].Children()[1].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Children()[1].Ref(),
//					"-i", fmt.Sprintf("%s", sig[1].Ref()),
//					"-c", fmt.Sprintf("%s=passed", sig[1].Children()[1].Ref()),
//					"-r", fmt.Sprintf("%s=%s&&%s", sig[1].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-3"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		Containers: []corev1.Container{
//			{
//				Name:            sig[2].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[2].Ref(),
//					"-c", fmt.Sprintf("%s=passed", sig[2].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s", sig[0].Ref(), sig[2].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-4"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		SecurityContext: &corev1.PodSecurityContext{
//			FSGroup: common.Ptr(constants.DefaultFsGroup),
//		},
//	}
//
//	assert.NoError(t, err)
//	assert.Equal(t, want, res.Job.Spec.Template.Spec)
//}
//
//func TestProcessNegativeSteps(t *testing.T) {
//	wf := &testworkflowsv1.TestWorkflow{
//		Spec: testworkflowsv1.TestWorkflowSpec{
//			Steps: []testworkflowsv1.Step{
//				{StepMeta: testworkflowsv1.StepMeta{Name: "A"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test"}},
//				{
//					StepMeta:    testworkflowsv1.StepMeta{Name: "B"},
//					StepControl: testworkflowsv1.StepControl{Negative: true},
//					Steps: []testworkflowsv1.Step{
//						{StepMeta: testworkflowsv1.StepMeta{Name: "C"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-2"}},
//						{StepMeta: testworkflowsv1.StepMeta{Name: "D"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-3"}},
//					},
//				},
//				{StepMeta: testworkflowsv1.StepMeta{Name: "E"}, StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-4"}},
//			},
//		},
//	}
//
//	res, err := proc.Bundle(context.Background(), wf, execMachine)
//	sig := res.Signature
//
//	volumes := res.Job.Spec.Template.Spec.Volumes
//	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts
//
//	want := corev1.PodSpec{
//		RestartPolicy:      corev1.RestartPolicyNever,
//		EnableServiceLinks: common.Ptr(false),
//		Volumes:            volumes,
//		InitContainers: []corev1.Container{
//			{
//				Name:            "tktw-init",
//				Image:           constants.DefaultInitImage,
//				ImagePullPolicy: corev1.PullIfNotPresent,
//				Command:         []string{"/bin/sh", "-c"},
//				Env:             initEnvs,
//				VolumeMounts:    volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[0].Ref(),
//					"-c", fmt.Sprintf("%s,%s,%s,%s=passed", sig[0].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref(), sig[2].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s&&%s", sig[0].Ref(), sig[1].Ref(), sig[2].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[1].Children()[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Children()[0].Ref(),
//					"-i", fmt.Sprintf("%s.v", sig[1].Ref()),
//					"-c", fmt.Sprintf("%s,%s,%s,%s.v=passed", sig[1].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref(), sig[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s&&%s", sig[0].Ref(), sig[1].Ref(), sig[2].Ref()),
//					"-r", fmt.Sprintf("%s=!%s.v", sig[1].Ref(), sig[1].Ref()),
//					"-r", fmt.Sprintf("%s.v=%s&&%s", sig[1].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-2"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[1].Children()[1].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Children()[1].Ref(),
//					"-i", fmt.Sprintf("%s.v", sig[1].Ref()),
//					"-c", fmt.Sprintf("%s=passed", sig[1].Children()[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s&&%s", sig[0].Ref(), sig[1].Ref(), sig[2].Ref()),
//					"-r", fmt.Sprintf("%s=!%s.v", sig[1].Ref(), sig[1].Ref()),
//					"-r", fmt.Sprintf("%s.v=%s&&%s", sig[1].Ref(), sig[1].Children()[0].Ref(), sig[1].Children()[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-3"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		Containers: []corev1.Container{
//			{
//				Name:            sig[2].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[2].Ref(),
//					"-c", fmt.Sprintf("%s=passed", sig[2].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s&&%s", sig[0].Ref(), sig[1].Ref(), sig[2].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-4"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		SecurityContext: &corev1.PodSecurityContext{
//			FSGroup: common.Ptr(constants.DefaultFsGroup),
//		},
//	}
//
//	assert.NoError(t, err)
//	assert.Equal(t, want, res.Job.Spec.Template.Spec)
//}
//
//func TestProcessNegativeContainerStep(t *testing.T) {
//	wf := &testworkflowsv1.TestWorkflow{
//		Spec: testworkflowsv1.TestWorkflowSpec{
//			Steps: []testworkflowsv1.Step{
//				{StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test"}},
//				{StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-2"}, StepControl: testworkflowsv1.StepControl{Negative: true}},
//			},
//		},
//	}
//
//	res, err := proc.Bundle(context.Background(), wf, execMachine)
//	sig := res.Signature
//
//	volumes := res.Job.Spec.Template.Spec.Volumes
//	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts
//
//	want := corev1.PodSpec{
//		RestartPolicy:      corev1.RestartPolicyNever,
//		EnableServiceLinks: common.Ptr(false),
//		Volumes:            volumes,
//		InitContainers: []corev1.Container{
//			{
//				Name:            "tktw-init",
//				Image:           constants.DefaultInitImage,
//				ImagePullPolicy: corev1.PullIfNotPresent,
//				Command:         []string{"/bin/sh", "-c"},
//				Env:             initEnvs,
//				VolumeMounts:    volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[0].Ref(),
//					"-c", fmt.Sprintf("%s,%s=passed", sig[0].Ref(), sig[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s", sig[0].Ref(), sig[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		Containers: []corev1.Container{
//			{
//				Name:            sig[1].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Ref(),
//					"-n", "true",
//					"-c", fmt.Sprintf("%s=passed", sig[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s", sig[0].Ref(), sig[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-2"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//
//		SecurityContext: &corev1.PodSecurityContext{
//			FSGroup: common.Ptr(constants.DefaultFsGroup),
//		},
//	}
//
//	assert.NoError(t, err)
//	assert.Equal(t, want, res.Job.Spec.Template.Spec)
//}
//
//func TestProcessOptionalContainerStep(t *testing.T) {
//	wf := &testworkflowsv1.TestWorkflow{
//		Spec: testworkflowsv1.TestWorkflowSpec{
//			Steps: []testworkflowsv1.Step{
//				{StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test"}},
//				{StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-2"}, StepControl: testworkflowsv1.StepControl{Optional: true}},
//			},
//		},
//	}
//
//	res, err := proc.Bundle(context.Background(), wf, execMachine)
//	sig := res.Signature
//
//	volumes := res.Job.Spec.Template.Spec.Volumes
//	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts
//
//	want := corev1.PodSpec{
//		RestartPolicy:      corev1.RestartPolicyNever,
//		EnableServiceLinks: common.Ptr(false),
//		Volumes:            volumes,
//		InitContainers: []corev1.Container{
//			{
//				Name:            "tktw-init",
//				Image:           constants.DefaultInitImage,
//				ImagePullPolicy: corev1.PullIfNotPresent,
//				Command:         []string{"/bin/sh", "-c"},
//				Env:             initEnvs,
//				VolumeMounts:    volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[0].Ref(),
//					"-c", fmt.Sprintf("%s,%s=passed", sig[0].Ref(), sig[1].Ref()),
//					"-r", fmt.Sprintf("=%s", sig[0].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		Containers: []corev1.Container{
//			{
//				Name:            sig[1].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Ref(),
//					"-c", fmt.Sprintf("%s=passed", sig[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-2"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		SecurityContext: &corev1.PodSecurityContext{
//			FSGroup: common.Ptr(constants.DefaultFsGroup),
//		},
//	}
//
//	assert.NoError(t, err)
//	assert.Equal(t, want, res.Job.Spec.Template.Spec)
//}
//
//func TestProcessLocalContent(t *testing.T) {
//	wf := &testworkflowsv1.TestWorkflow{
//		Spec: testworkflowsv1.TestWorkflowSpec{
//			Steps: []testworkflowsv1.Step{
//				{StepOperations: testworkflowsv1.StepOperations{
//					Shell: "shell-test",
//				}, StepSource: testworkflowsv1.StepSource{
//					Content: &testworkflowsv1.Content{
//						Files: []testworkflowsv1.ContentFile{{
//							Path:    "/some/path",
//							Content: `some-{{"{{"}}content`,
//						}},
//					},
//				}},
//				{StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-2"}},
//			},
//		},
//	}
//
//	res, err := proc.Bundle(context.Background(), wf, execMachine)
//	assert.NoError(t, err)
//
//	sig := res.Signature
//
//	volumes := res.Job.Spec.Template.Spec.Volumes
//	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts
//	volumeMountsWithContent := res.Job.Spec.Template.Spec.InitContainers[1].VolumeMounts
//
//	want := corev1.PodSpec{
//		RestartPolicy:      corev1.RestartPolicyNever,
//		EnableServiceLinks: common.Ptr(false),
//		Volumes:            volumes,
//		InitContainers: []corev1.Container{
//			{
//				Name:            "tktw-init",
//				Image:           constants.DefaultInitImage,
//				ImagePullPolicy: corev1.PullIfNotPresent,
//				Command:         []string{"/bin/sh", "-c"},
//				Env:             initEnvs,
//				VolumeMounts:    volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[0].Ref(),
//					"-c", fmt.Sprintf("%s,%s=passed", sig[0].Ref(), sig[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s", sig[0].Ref(), sig[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMountsWithContent,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		Containers: []corev1.Container{
//			{
//				Name:            sig[1].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Ref(),
//					"-c", fmt.Sprintf("%s=passed", sig[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s", sig[0].Ref(), sig[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-2"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		SecurityContext: &corev1.PodSecurityContext{
//			FSGroup: common.Ptr(constants.DefaultFsGroup),
//		},
//	}
//
//	assert.Equal(t, want, res.Job.Spec.Template.Spec)
//	assert.Equal(t, 2, len(volumeMounts))
//	assert.Equal(t, 3, len(volumeMountsWithContent))
//	assert.Equal(t, volumeMounts, volumeMountsWithContent[:2])
//	assert.Equal(t, "/some/path", volumeMountsWithContent[2].MountPath)
//	assert.Equal(t, 1, len(res.ConfigMaps))
//	assert.Equal(t, volumeMountsWithContent[2].Name, volumes[2].Name)
//	assert.Equal(t, volumes[2].ConfigMap.Name, res.ConfigMaps[0].Name)
//	assert.Equal(t, "some-{{content", res.ConfigMaps[0].Data[volumeMountsWithContent[2].SubPath])
//}
//
//func TestProcessGlobalContent(t *testing.T) {
//	wf := &testworkflowsv1.TestWorkflow{
//		Spec: testworkflowsv1.TestWorkflowSpec{
//			TestWorkflowSpecBase: testworkflowsv1.TestWorkflowSpecBase{
//				Content: &testworkflowsv1.Content{
//					Files: []testworkflowsv1.ContentFile{{
//						Path:    "/some/path",
//						Content: `some-{{"{{"}}content`,
//					}},
//				},
//			},
//			Steps: []testworkflowsv1.Step{
//				{StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test"}},
//				{StepOperations: testworkflowsv1.StepOperations{Shell: "shell-test-2"}},
//			},
//		},
//	}
//
//	res, err := proc.Bundle(context.Background(), wf, execMachine)
//	assert.NoError(t, err)
//
//	sig := res.Signature
//
//	volumes := res.Job.Spec.Template.Spec.Volumes
//	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts
//
//	want := corev1.PodSpec{
//		RestartPolicy:      corev1.RestartPolicyNever,
//		EnableServiceLinks: common.Ptr(false),
//		Volumes:            volumes,
//		InitContainers: []corev1.Container{
//			{
//				Name:            "tktw-init",
//				Image:           constants.DefaultInitImage,
//				ImagePullPolicy: corev1.PullIfNotPresent,
//				Command:         []string{"/bin/sh", "-c"},
//				Env:             initEnvs,
//				VolumeMounts:    volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//			{
//				Name:            sig[0].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[0].Ref(),
//					"-c", fmt.Sprintf("%s,%s=passed", sig[0].Ref(), sig[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s", sig[0].Ref(), sig[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		Containers: []corev1.Container{
//			{
//				Name:            sig[1].Ref(),
//				ImagePullPolicy: "",
//				Image:           constants.DefaultInitImage,
//				Command: []string{
//					"/.tktw/init",
//					sig[1].Ref(),
//					"-c", fmt.Sprintf("%s=passed", sig[1].Ref()),
//					"-r", fmt.Sprintf("=%s&&%s", sig[0].Ref(), sig[1].Ref()),
//					"--",
//				},
//				Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test-2"},
//				WorkingDir:   "",
//				EnvFrom:      []corev1.EnvFromSource(nil),
//				Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//				Resources:    corev1.ResourceRequirements{},
//				VolumeMounts: volumeMounts,
//				SecurityContext: &corev1.SecurityContext{
//					RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//				},
//			},
//		},
//		SecurityContext: &corev1.PodSecurityContext{
//			FSGroup: common.Ptr(constants.DefaultFsGroup),
//		},
//	}
//
//	assert.Equal(t, want, res.Job.Spec.Template.Spec)
//	assert.Equal(t, 3, len(volumeMounts))
//	assert.Equal(t, "/some/path", volumeMounts[2].MountPath)
//	assert.Equal(t, 1, len(res.ConfigMaps))
//	assert.Equal(t, volumeMounts[2].Name, volumes[2].Name)
//	assert.Equal(t, volumes[2].ConfigMap.Name, res.ConfigMaps[0].Name)
//	assert.Equal(t, "some-{{content", res.ConfigMaps[0].Data[volumeMounts[2].SubPath])
//}
//
//func TestProcessRunShell(t *testing.T) {
//	wf := &testworkflowsv1.TestWorkflow{
//		Spec: testworkflowsv1.TestWorkflowSpec{
//			Steps: []testworkflowsv1.Step{
//				{StepOperations: testworkflowsv1.StepOperations{Run: &testworkflowsv1.StepRun{Shell: common.Ptr("shell-test")}}},
//			},
//		},
//	}
//
//	res, err := proc.Bundle(context.Background(), wf, execMachine)
//	assert.NoError(t, err)
//
//	sig := res.Signature
//	sigSerialized, _ := json.Marshal(sig)
//
//	volumes := res.Job.Spec.Template.Spec.Volumes
//	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts
//
//	want := batchv1.Job{
//		TypeMeta: metav1.TypeMeta{Kind: "Job", APIVersion: "batch/v1"},
//		ObjectMeta: metav1.ObjectMeta{
//			Name: "dummy-id-abc",
//			Labels: map[string]string{
//				constants.RootResourceIdLabelName: "dummy-id",
//				constants.ResourceIdLabelName:     "dummy-id-abc",
//			},
//			Annotations: map[string]string{
//				constants.SignatureAnnotationName: string(sigSerialized),
//			},
//		},
//		Spec: batchv1.JobSpec{
//			BackoffLimit: common.Ptr(int32(0)),
//			Template: corev1.PodTemplateSpec{
//				ObjectMeta: metav1.ObjectMeta{
//					Labels: map[string]string{
//						constants.RootResourceIdLabelName: "dummy-id",
//						constants.ResourceIdLabelName:     "dummy-id-abc",
//					},
//					Annotations: map[string]string(nil),
//				},
//				Spec: corev1.PodSpec{
//					RestartPolicy:      corev1.RestartPolicyNever,
//					EnableServiceLinks: common.Ptr(false),
//					Volumes:            volumes,
//					InitContainers: []corev1.Container{
//						{
//							Name:            "tktw-init",
//							Image:           constants.DefaultInitImage,
//							ImagePullPolicy: corev1.PullIfNotPresent,
//							Command:         []string{"/bin/sh", "-c"},
//							Env:             initEnvs,
//							VolumeMounts:    volumeMounts,
//							SecurityContext: &corev1.SecurityContext{
//								RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//							},
//						},
//					},
//					Containers: []corev1.Container{
//						{
//							Name:            sig[0].Ref(),
//							ImagePullPolicy: "",
//							Image:           constants.DefaultInitImage,
//							Command: []string{
//								"/.tktw/init",
//								sig[0].Ref(),
//								"-c", fmt.Sprintf("%s=passed", sig[0].Ref()),
//								"-r", fmt.Sprintf("=%s", sig[0].Ref()),
//								"--",
//							},
//							Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test"},
//							WorkingDir:   "",
//							EnvFrom:      []corev1.EnvFromSource(nil),
//							Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//							Resources:    corev1.ResourceRequirements{},
//							VolumeMounts: volumeMounts,
//							SecurityContext: &corev1.SecurityContext{
//								RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//							},
//						},
//					},
//					SecurityContext: &corev1.PodSecurityContext{
//						FSGroup: common.Ptr(constants.DefaultFsGroup),
//					},
//				},
//			},
//		},
//	}
//
//	assert.Equal(t, want, res.Job)
//
//	assert.Equal(t, 2, len(volumeMounts))
//	assert.Equal(t, 2, len(volumes))
//	assert.Equal(t, constants.DefaultInternalPath, volumeMounts[0].MountPath)
//	assert.Equal(t, constants.DefaultDataPath, volumeMounts[1].MountPath)
//	assert.True(t, volumeMounts[0].Name == volumes[0].Name)
//	assert.True(t, volumeMounts[1].Name == volumes[1].Name)
//}
//
//func TestProcessEscapedAnnotations(t *testing.T) {
//	wf := &testworkflowsv1.TestWorkflow{
//		Spec: testworkflowsv1.TestWorkflowSpec{
//			TestWorkflowSpecBase: testworkflowsv1.TestWorkflowSpecBase{
//				Pod: &testworkflowsv1.PodConfig{
//					Annotations: map[string]string{
//						"vault.hashicorp.com/agent-inject-template-database-config.txt": `{{"{{"}}- with secret "internal/data/database/config" -}}{{"{{"}} .Data.data.username }}@{{"{{"}} .Data.data.password }}{{"{{"}}- end -}}`,
//					},
//				},
//			},
//			Steps: []testworkflowsv1.Step{
//				{StepOperations: testworkflowsv1.StepOperations{Run: &testworkflowsv1.StepRun{Shell: common.Ptr("shell-test")}}},
//			},
//		},
//	}
//
//	res, err := proc.Bundle(context.Background(), wf, execMachine)
//	assert.NoError(t, err)
//
//	sig := res.Signature
//	sigSerialized, _ := json.Marshal(sig)
//
//	volumes := res.Job.Spec.Template.Spec.Volumes
//	volumeMounts := res.Job.Spec.Template.Spec.InitContainers[0].VolumeMounts
//
//	want := batchv1.Job{
//		TypeMeta: metav1.TypeMeta{Kind: "Job", APIVersion: "batch/v1"},
//		ObjectMeta: metav1.ObjectMeta{
//			Name: "dummy-id-abc",
//			Labels: map[string]string{
//				constants.RootResourceIdLabelName: "dummy-id",
//				constants.ResourceIdLabelName:     "dummy-id-abc",
//			},
//			Annotations: map[string]string{
//				constants.SignatureAnnotationName: string(sigSerialized),
//			},
//		},
//		Spec: batchv1.JobSpec{
//			BackoffLimit: common.Ptr(int32(0)),
//			Template: corev1.PodTemplateSpec{
//				ObjectMeta: metav1.ObjectMeta{
//					Labels: map[string]string{
//						constants.RootResourceIdLabelName: "dummy-id",
//						constants.ResourceIdLabelName:     "dummy-id-abc",
//					},
//					Annotations: map[string]string{
//						"vault.hashicorp.com/agent-inject-template-database-config.txt": `{{- with secret "internal/data/database/config" -}}{{ .Data.data.username }}@{{ .Data.data.password }}{{- end -}}`,
//					},
//				},
//				Spec: corev1.PodSpec{
//					RestartPolicy:      corev1.RestartPolicyNever,
//					EnableServiceLinks: common.Ptr(false),
//					Volumes:            volumes,
//					InitContainers: []corev1.Container{
//						{
//							Name:            "tktw-init",
//							Image:           constants.DefaultInitImage,
//							ImagePullPolicy: corev1.PullIfNotPresent,
//							Command:         []string{"/bin/sh", "-c"},
//							Env:             initEnvs,
//							VolumeMounts:    volumeMounts,
//							SecurityContext: &corev1.SecurityContext{
//								RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//							},
//						},
//					},
//					Containers: []corev1.Container{
//						{
//							Name:            sig[0].Ref(),
//							ImagePullPolicy: "",
//							Image:           constants.DefaultInitImage,
//							Command: []string{
//								"/.tktw/init",
//								sig[0].Ref(),
//								"-c", fmt.Sprintf("%s=passed", sig[0].Ref()),
//								"-r", fmt.Sprintf("=%s", sig[0].Ref()),
//								"--",
//							},
//							Args:         []string{constants.DefaultShellPath, "-c", constants.DefaultShellHeader + "shell-test"},
//							WorkingDir:   "",
//							EnvFrom:      []corev1.EnvFromSource(nil),
//							Env:          []corev1.EnvVar{{Name: "CI", Value: "1"}},
//							Resources:    corev1.ResourceRequirements{},
//							VolumeMounts: volumeMounts,
//							SecurityContext: &corev1.SecurityContext{
//								RunAsGroup: common.Ptr(constants.DefaultFsGroup),
//							},
//						},
//					},
//					SecurityContext: &corev1.PodSecurityContext{
//						FSGroup: common.Ptr(constants.DefaultFsGroup),
//					},
//				},
//			},
//		},
//	}
//
//	assert.Equal(t, want, res.Job)
//
//	assert.Equal(t, 2, len(volumeMounts))
//	assert.Equal(t, 2, len(volumes))
//	assert.Equal(t, constants.DefaultInternalPath, volumeMounts[0].MountPath)
//	assert.Equal(t, constants.DefaultDataPath, volumeMounts[1].MountPath)
//	assert.True(t, volumeMounts[0].Name == volumes[0].Name)
//	assert.True(t, volumeMounts[1].Name == volumes[1].Name)
//}
