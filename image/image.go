package image

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hasty-ai/hasty-go"
	"github.com/hasty-ai/hasty-go/backend"
)

// New instantiates images client
func New(backend *backend.Backend) *Client {
	return &Client{
		backend: backend,
	}
}

// Client to access images API
type Client struct {
	backend *backend.Backend
}

// UploadExternal one single image from an external source
func (c *Client) UploadExternal(ctx context.Context, params *hasty.ImageUploadExternalParams) (*hasty.Image, error) {
	if params.Project == nil {
		return nil, fmt.Errorf("project must be specified")
	}
	path := fmt.Sprintf("/v1/projects/%s/images", *params.Project)
	method := http.MethodPost
	var response hasty.Image
	status, err := c.backend.Request(ctx, method, path, params, &response)
	_ = status
	return nil, err
}
