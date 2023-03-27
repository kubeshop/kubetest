/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

// repository representation for tests in git repositories
type Repository struct {
	// VCS repository type
	Type_ string `json:"type"`
	// uri of content file or git directory
	Uri string `json:"uri"`
	// branch/tag name for checkout
	Branch string `json:"branch,omitempty"`
	// commit id (sha) for checkout
	Commit string `json:"commit,omitempty"`
	// if needed we can checkout particular path (dir or file) in case of BIG/mono repositories
	Path string `json:"path,omitempty"`
	// git auth username for private repositories
	Username string `json:"username,omitempty"`
	// git auth token for private repositories
	Token          string     `json:"token,omitempty"`
	UsernameSecret *SecretRef `json:"usernameSecret,omitempty"`
	TokenSecret    *SecretRef `json:"tokenSecret,omitempty"`
	// secret with certificate for private repositories
	CertificateSecret string `json:"certificateSecret,omitempty"`
	// if provided we checkout the whole repository and run test from this directory
	WorkingDir string `json:"workingDir,omitempty"`
	// auth type for git requests
	AuthType string `json:"authType,omitempty"`
}
