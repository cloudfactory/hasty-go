package hasty

import (
	"context"
	"fmt"
	"net/http"
)

// ImageUploadExternalParams is used for uploading images from external sources
type ImageUploadExternalParams struct {
	Project  *string `json:"-"`
	Dataset  *string `json:"dataset_id"`
	URL      *string `json:"url"`
	Copy     *bool   `json:"copy_original,omitempty"`
	Filename *string `json:"filename,omitempty"`
}

// Image describes an image information that API may return
type Image struct {
	ID *string `json:"image_id"`
}

// NewImageClient instantiates images client
func NewImageClient(backend Backend) *ImageClient {
	return &ImageClient{
		backend: backend,
	}
}

// ImageClient is a client to access images API
type ImageClient struct {
	backend Backend
}

// UploadExternal one single image from an external source
func (c *ImageClient) UploadExternal(ctx context.Context, params *ImageUploadExternalParams) (*Image, error) {
	if params.Project == nil || *params.Project == "" {
		return nil, fmt.Errorf("project must be specified")
	}
	path := fmt.Sprintf("/v1/projects/%s/images", *params.Project)
	method := http.MethodPost
	var response Image
	status, err := c.backend.Request(ctx, method, path, params, &response)
	_ = status
	return nil, err
}
