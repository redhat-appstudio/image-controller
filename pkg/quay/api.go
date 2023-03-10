package quay

type Tag struct {
	ImageId        string `json:"image_id"`
	TrustEnabled   string `json:"trust_enabled"`
	Name           string `json:"name"`
	ManifestDigest string `json:"manifest_digest,omitempty"`
	Size           int    `json:"int"`
}

type Repository struct {
	TrustEnabled   bool           `json:"trust_enabled"`
	Description    string         `json:"description"`
	CanAdmin       bool           `json:"can_admin"`
	CanWrite       bool           `json:"can_write"`
	IsOrganization bool           `json:"is_organization"`
	IsStarred      bool           `json:"is_starred"`
	IsPublic       bool           `json:"is_public"`
	Name           string         `json:"name"`
	Namespace      string         `json:"namespace"`
	Image          string         `json:"image"`
	TagExpirationS int            `json:"tag_expiration_s"`
	Tags           map[string]Tag `json:"tags"`
	StatusToken    string         `json:"status_token"`
	ErrorMessage   string         `json:"error_message"`
}

type RepositoryRequest struct {
	Namespace   string `json:"namespace"`
	Visibility  string `json:"visibility"`
	Repository  string `json:"repository"`
	Description string `json:"description"`
	//Kind        string `json:"repo_kind"`
}
type RobotAccount struct {
	Description string `json:"description"`
	Created     string `json:"created"`
	// UnstructuredData []byte  `json:"unstructured_metadata"`
	LastAccessed string `json:"last_accessed"`
	Token        string `json:"token"`
	Name         string `json:"name"`
	Message      string `json:"message"`
}

type QuayError struct {
	ErrorMessage string `json:"error_type"`
}
