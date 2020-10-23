package image

import (
	"fmt"
	"net/http"

	"github.com/hasty-ai/hasty-go"
	"github.com/hasty-ai/hasty-go/client"
)

// New instantiates images client
func New(backend *client.Backend) *Client {
	return &Client{
		backend: backend,
	}
}

// Client to access images API
type Client struct {
	backend *client.Backend
}

// Upload one single image
func (c *Client) Upload(request *hasty.ImageUploadParams) (hasty.Image, error) {
	method := http.MethodPost
	path := fmt.Sprintf("/v1/projects/%s/images", request.Project)
	response:= ??
	status, err := c.backend.Request(method, path, request, &response)
	return response, err
}
