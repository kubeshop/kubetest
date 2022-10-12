package secret

import (
	"bytes"
	"fmt"
	"os"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	corev1 "k8s.io/api/core/v1"
)

const (
	// SecretEnvVarPrefix is a prefix for secret env vars
	SecretEnvVarPrefix = "RUNNER_SECRET_ENV"
	// SecretVarPrefix is a prefix for secret vars
	SecretVarPrefix = "RUNNER_SECRET_VAR_"
)

// Manager is responsible for exchanging secrets with executor pod
type Manager interface {
	// Prepare prepares secret env vars based on secret envs and variables
	Prepare(secretEnvs map[string]string, variables map[string]testkube.Variable) (secretEnvVars []corev1.EnvVar)
	// GetEnvs get secret envs
	GetEnvs() (secretEnvs []string)
	// GetVars gets secret vars
	GetVars(variables map[string]testkube.Variable)
	// Obfuscate obfuscates secret values
	Obfuscate(p []byte) []byte
}

// NewEnvManager returns an implementation of the Manager
func NewEnvManager() *EnvManager {
	return &EnvManager{}
}

func NewEnvManagerWithVars(variables map[string]testkube.Variable) *EnvManager {
	return &EnvManager{
		Variables: variables,
	}
}

// EnvManager manages secret exchange from job pods using env
type EnvManager struct {
	Variables map[string]testkube.Variable
}

// Prepare prepares secret env vars based on secret envs and variables
func (m EnvManager) Prepare(secretEnvs map[string]string, variables map[string]testkube.Variable) (secretEnvVars []corev1.EnvVar) {
	// preparet secret envs
	i := 1
	for secretName, secretVar := range secretEnvs {
		// TODO: these are duplicated because Postman executor is expecting it as json string
		// and gets unmarshalled and the name and the value are taken from there, for other executors it will be like a normal env var.
		secretEnvVars = append(secretEnvVars, corev1.EnvVar{
			Name: secretVar,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secretName,
					},
					Key: secretVar,
				},
			},
		})
		secretEnvVars = append(secretEnvVars, corev1.EnvVar{
			Name: fmt.Sprintf("%s%d", SecretEnvVarPrefix, i),
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secretName,
					},
					Key: secretVar,
				},
			},
		})
		i++
	}

	// prepare secret vars
	for name, variable := range variables {
		if !variable.IsSecret() || variable.SecretRef == nil {
			continue
		}

		secretEnvVars = append(secretEnvVars, corev1.EnvVar{
			Name: fmt.Sprintf("%s%s", SecretVarPrefix, name),
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: variable.SecretRef.Name,
					},
					Key: variable.SecretRef.Key,
				},
			},
		})
	}

	return secretEnvVars
}

// GetEnvs gets secret envs
func (m EnvManager) GetEnvs() (secretEnvs []string) {
	i := 1
	for {
		secretEnv, ok := os.LookupEnv(fmt.Sprintf("%s%d", SecretEnvVarPrefix, i))
		if !ok {
			break
		}

		secretEnvs = append(secretEnvs, secretEnv)
		i++
	}

	return secretEnvs
}

// GetVars gets secret vars
func (m EnvManager) GetVars(variables map[string]testkube.Variable) {
	for name, variable := range variables {
		if !variable.IsSecret() {
			continue
		}

		value, ok := os.LookupEnv(fmt.Sprintf("%s%s", SecretVarPrefix, name))
		if !ok {
			continue
		}

		variable.Value = value
		variables[name] = variable
	}

	return
}

func (m EnvManager) Obfuscate(p []byte) []byte {
	if m.Variables == nil {
		return p
	}
	m.GetVars(m.Variables)
	for _, variable := range m.Variables {
		if !variable.IsSecret() {
			continue
		}

		p = bytes.ReplaceAll(p, []byte(variable.Value), []byte("*****"))
	}

	return p
}
