package hasty

import (
	"context"
	"fmt"
	"net/http"
)

// ImageStatus for passing as status of the image
type ImageStatus string

// All possible image statuses
const (
	ImageStatusNew          ImageStatus = "NEW"
	ImageStatusInProgress   ImageStatus = "IN PROGRESS"
	ImageStatusDone         ImageStatus = "DONE"
	ImageStatusSkipped      ImageStatus = "SKIPPED"
	ImageStatusToReview     ImageStatus = "TO REVIEW"
	ImageStatusAutoLabelled ImageStatus = "AUTO-LABELLED"
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
	ID           *string      `json:"id"`
	Height       *int         `json:"height"`
	Width        *int         `json:"width"`
	Format       *string      `json:"format"`
	Mode         *string      `json:"mode"`
	Name         *string      `json:"name"`
	Status       *ImageStatus `json:"status"`
	OriginalURL  *string      `json:"public_url"`
	ThumbnailURL *string      `json:"thumbnail_url"`
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
	if err != nil {
		return nil, fmt.Errorf("unable to execute upload request: %w", err)
	}
	switch status {
	case http.StatusOK:
	case http.StatusUnauthorized:
		return nil, ErrAuth
	case http.StatusForbidden:
		return nil, ErrPerm
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusTooManyRequests:
		return nil, ErrRate
	default:
		return nil, fmt.Errorf("unexpected API response status: %d", status)
	}
	return &response, nil
}
