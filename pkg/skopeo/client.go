package skopeo

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"

	"github.com/kubeshop/testkube/pkg/process"
)

// DockerAuths contains an embedded DockerAuthConfigs
type DockerAuths struct {
	Auths DockerAuthConfigs `json:"auths"`
}

// DockerAuthConfigs is a map of registries and their credentials
type DockerAuthConfigs map[string]DockerAuthConfig

// DockerAuthConfig contains authorization information for connecting to a registry
// It mirrors "github.com/docker/docker/api/types.AuthConfig"
type DockerAuthConfig struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Auth     string `json:"auth,omitempty"`

	// Email is an optional value associated with the username.
	// This field is deprecated and will be removed in a later
	// version of docker.
	Email string `json:"email,omitempty"`
}

// DockerImage contains definition of docker image
type DockerImage struct {
	Config struct {
		User       string   `json:"User"`
		Entrypoint []string `json:"Entrypoint"`
		Cmd        []string `json:"Cmd"`
		WorkingDir string   `json:"WorkingDir"`
	} `json:"config"`
	History []struct {
		Created   time.Time `json:"created"`
		CreatedBy string    `json:"created_by"`
	} `json:"history"`
	Shell string `json:"-"`
}

// Inspector is image inspector interface
type Inspector interface {
	Inspect(registry, image string) (*DockerImage, error)
}

type client struct {
	dockerAuthConfigs []DockerAuthConfig
}

var _ Inspector = (*client)(nil)

// NewClient creates new empty client
func NewClient() *client {
	return &client{}
}

// NewClientFromSecrets creats new client from secrets
func NewClientFromSecrets(imageSecrets []corev1.Secret, registry string) (*client, error) {
	auths, err := ParseSecretData(imageSecrets, registry)
	if err != nil {
		return nil, err
	}

	return &client{dockerAuthConfigs: auths}, nil
}

// Inspect inspect a docker image
func (c *client) Inspect(registry, image string) (*DockerImage, error) {
	args := []string{
		"--override-os",
		"linux",
		"inspect",
	}

	if len(c.dockerAuthConfigs) != 0 {
		i := rand.Intn(len(c.dockerAuthConfigs))
		args = append(args, "--creds", c.dockerAuthConfigs[i].Username+":"+c.dockerAuthConfigs[i].Password)
	}

	config := "docker://" + image
	if registry != "" {
		config = registry + "/" + image
	}

	args = append(args, "--config", config)
	result, err := process.Execute("skopeo", args...)
	if err != nil {
		return nil, err
	}

	var dockerImage DockerImage
	if err = json.Unmarshal(result, &dockerImage); err != nil {
		return nil, err
	}

	var shell string
	for i := len(dockerImage.History); i > 0; i-- {
		command := dockerImage.History[i-1].CreatedBy
		re, err := regexp.Compile(`/bin/([a-z]*)sh`)
		if err != nil {
			return nil, err
		}

		shell = re.FindString(command)
		if shell != "" {
			break
		}
	}

	dockerImage.Shell = shell
	return &dockerImage, nil
}

// ParseSecretData parses secret data for docker auth config
func ParseSecretData(imageSecrets []corev1.Secret, registry string) ([]DockerAuthConfig, error) {
	var results []DockerAuthConfig
	for _, imageSecret := range imageSecrets {
		auths := DockerAuths{}
		if jsonData, ok := imageSecret.Data[".dockerconfigjson"]; ok {
			if err := json.Unmarshal(jsonData, &auths); err != nil {
				return nil, err
			}
		} else if configData, ok := imageSecret.Data[".dockercfg"]; ok {
			if err := json.Unmarshal(configData, &auths.Auths); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("imagePullSecret %s contains neither .dockercfg nor .dockerconfigjson", imageSecret.Name)
		}

		// Determine if there is a secret for the specified registry
		if creds, ok := auths.Auths[registry]; ok {
			username, password, err := extractRegistryCredentials(creds)
			if err != nil {
				return nil, err
			}

			results = append(results, DockerAuthConfig{Username: username, Password: password})
		} else {
			return nil, fmt.Errorf("secret %s is not defined for registry: %s", imageSecret.Name, registry)
		}
	}

	return results, nil
}

func extractRegistryCredentials(creds DockerAuthConfig) (username, password string, err error) {
	if creds.Auth == "" {
		return creds.Username, creds.Password, nil
	}

	decoder := base64.StdEncoding
	if !strings.HasSuffix(strings.TrimSpace(creds.Auth), "=") {
		// Modify the decoder to be raw if no padding is present
		decoder = decoder.WithPadding(base64.NoPadding)
	}

	base64Decoded, err := decoder.DecodeString(creds.Auth)
	if err != nil {
		return "", "", err
	}

	splitted := strings.SplitN(string(base64Decoded), ":", 2)
	if len(splitted) != 2 {
		return creds.Username, creds.Password, nil
	}

	return splitted[0], splitted[1], nil
}
