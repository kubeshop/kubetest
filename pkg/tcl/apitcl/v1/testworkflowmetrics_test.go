package v1

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	testworkflowsv1 "github.com/kubeshop/testkube-operator/api/testworkflows/v1"
	configRepo "github.com/kubeshop/testkube/pkg/repository/config"
)

func Test_apiTCL_getClusterID(t *testing.T) {
	t.Parallel()

	t.Run("Get Cluster ID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		configRepo := configRepo.NewMockRepository(ctrl)
		clusterID := "cluster-id"
		var api apiTCL
		api.configMap = configRepo
		configRepo.EXPECT().GetUniqueClusterId(gomock.Any()).Return(clusterID, nil)
		if got := api.getClusterID(context.Background()); got != clusterID {
			t.Errorf("apiTCL.getClusterID() = %v, want %v", got, clusterID)
		}
	})
}

func Test_getImage(t *testing.T) {
	t.Parallel()

	t.Run("Get Image from empty container", func(t *testing.T) {
		if got := getImage(nil); got != "" {
			t.Errorf("getImage() = %v, wanted empty", got)
		}
	})
	t.Run("Get Image from container", func(t *testing.T) {
		image := "container-image"
		container := &testworkflowsv1.ContainerConfig{
			Image: image,
		}

		if got := getImage(container); got != image {
			t.Errorf("getImage() = %v, want %v", got, image)
		}
	})
}

func Test_hasArtifacts(t *testing.T) {
	type args struct {
		steps []testworkflowsv1.Step
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "No artifacts",
			args: args{
				steps: []testworkflowsv1.Step{
					{
						StepBase: testworkflowsv1.StepBase{},
						Use:      []testworkflowsv1.TemplateRef{},
						Template: &testworkflowsv1.TemplateRef{},
						Setup:    []testworkflowsv1.Step{},
						Steps:    []testworkflowsv1.Step{},
					},
				},
			},
			want: false,
		},
		{
			name: "Has artifacts on first level only",
			args: args{
				steps: []testworkflowsv1.Step{
					{
						StepBase: testworkflowsv1.StepBase{
							Artifacts: &testworkflowsv1.StepArtifacts{
								Paths: []string{"path"},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Has artifacts on third level",
			args: args{
				steps: []testworkflowsv1.Step{
					{
						Setup: []testworkflowsv1.Step{
							{
								Setup: []testworkflowsv1.Step{
									{
										StepBase: testworkflowsv1.StepBase{
											Artifacts: &testworkflowsv1.StepArtifacts{
												Paths: []string{"path"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Has artifacts on multiple levels",
			args: args{
				steps: []testworkflowsv1.Step{
					{
						StepBase: testworkflowsv1.StepBase{
							Artifacts: &testworkflowsv1.StepArtifacts{
								Paths: []string{"path"},
							},
						},
						Setup: []testworkflowsv1.Step{
							{
								StepBase: testworkflowsv1.StepBase{
									Artifacts: &testworkflowsv1.StepArtifacts{
										Paths: []string{"path"},
									},
								},
								Setup: []testworkflowsv1.Step{
									{
										StepBase: testworkflowsv1.StepBase{
											Artifacts: &testworkflowsv1.StepArtifacts{
												Paths: []string{"path"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasArtifacts(tt.args.steps); got != tt.want {
				t.Errorf("hasArtifacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasTemplateArtifacts(t *testing.T) {
	type args struct {
		steps []testworkflowsv1.IndependentStep
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "No artifacts",
			args: args{
				steps: []testworkflowsv1.IndependentStep{
					{
						StepBase: testworkflowsv1.StepBase{},
						Setup:    []testworkflowsv1.IndependentStep{},
						Steps:    []testworkflowsv1.IndependentStep{},
					},
				},
			},
			want: false,
		},
		{
			name: "Has artifacts on first level only",
			args: args{
				steps: []testworkflowsv1.IndependentStep{
					{
						StepBase: testworkflowsv1.StepBase{
							Artifacts: &testworkflowsv1.StepArtifacts{
								Paths: []string{"path"},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Has artifacts on third level",
			args: args{
				steps: []testworkflowsv1.IndependentStep{
					{
						Setup: []testworkflowsv1.IndependentStep{
							{
								Setup: []testworkflowsv1.IndependentStep{
									{
										StepBase: testworkflowsv1.StepBase{
											Artifacts: &testworkflowsv1.StepArtifacts{
												Paths: []string{"path"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Has artifacts on multiple levels",
			args: args{
				steps: []testworkflowsv1.IndependentStep{
					{
						StepBase: testworkflowsv1.StepBase{
							Artifacts: &testworkflowsv1.StepArtifacts{
								Paths: []string{"path"},
							},
						},
						Setup: []testworkflowsv1.IndependentStep{
							{
								StepBase: testworkflowsv1.StepBase{
									Artifacts: &testworkflowsv1.StepArtifacts{
										Paths: []string{"path"},
									},
								},
								Setup: []testworkflowsv1.IndependentStep{
									{
										StepBase: testworkflowsv1.StepBase{
											Artifacts: &testworkflowsv1.StepArtifacts{
												Paths: []string{"path"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasTemplateArtifacts(tt.args.steps); got != tt.want {
				t.Errorf("hasArtifacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasKubeshopGitURI(t *testing.T) {
	type args struct {
		spec testworkflowsv1.TestWorkflowSpec
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "No Kubeshop Git URI",
			args: args{
				spec: testworkflowsv1.TestWorkflowSpec{
					TestWorkflowSpecBase: testworkflowsv1.TestWorkflowSpecBase{
						Content: &testworkflowsv1.Content{
							Git: &testworkflowsv1.ContentGit{
								Uri: "test-uri",
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "Has Kubeshop URI on first level only",
			args: args{
				spec: testworkflowsv1.TestWorkflowSpec{
					TestWorkflowSpecBase: testworkflowsv1.TestWorkflowSpecBase{
						Content: &testworkflowsv1.Content{
							Git: &testworkflowsv1.ContentGit{
								Uri: "github.com/kubeshop/testkube-tests-uri",
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Has Kubeshop URI on third level only",
			args: args{
				spec: testworkflowsv1.TestWorkflowSpec{
					Steps: []testworkflowsv1.Step{
						{
							Setup: []testworkflowsv1.Step{
								{
									StepBase: testworkflowsv1.StepBase{
										Content: &testworkflowsv1.Content{
											Git: &testworkflowsv1.ContentGit{
												Uri: "github.com/kubeshop/testkube-tests-uri",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasKubeshopGitURI(tt.args.spec); got != tt.want {
				t.Errorf("hasKubeshopGitURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDataSource(t *testing.T) {
	type args struct {
		content *testworkflowsv1.Content
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty content",
			args: args{
				content: nil,
			},
			want: "",
		},
		{
			name: "Git data source",
			args: args{
				content: &testworkflowsv1.Content{
					Git: &testworkflowsv1.ContentGit{
						Uri: "test-uri",
					},
				},
			},
			want: "git",
		},
		{
			name: "Files data source",
			args: args{
				content: &testworkflowsv1.Content{
					Files: []testworkflowsv1.ContentFile{
						{
							Path: "test-path",
						},
					},
				},
			},
			want: "files",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDataSource(tt.args.content); got != tt.want {
				t.Errorf("getDataSource() = %v, want %v", got, tt.want)
			}
		})
	}
}
