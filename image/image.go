package image

import (
	"fmt"

	"github.com/hasty-ai/hasty-go"
	"github.com/hasty-ai/hasty-go/client"
)

func New(caller) *Client {

}

// Client to access images API
type Client struct {
	caller *client.Backend
}

// Upload one single image
func (c *Client) Upload(params *hasty.ImageUploadParams) error {

	endpoint := fmt.Sprintf("https://api.hasty.ai", projectID)
	req := imageUpload{
		DatasetID: datasetID,
		URL:       url,
		Copy:      copy,
	}

	return c.client.post(endpoint, session, req)
}
