package hasty

import (
	"context"
	"fmt"
	"net/http"
)

// DatasetParams is parameters set for creating or updating a dataset
type DatasetParams struct {
	Project *string `json:"-"`
	Name    *string `json:"name"`
	Order   *int    `json:"norder"`
}

// Dataset describes an dataser information that API may return
type Dataset struct {
	ID    *string `json:"id"`
	Name  *string `json:"name"`
	Order *int    `json:"norder"`
}

// NewDatasetClient instantiates dataset client
func NewDatasetClient(backend Backend) *DatasetClient {
	return &DatasetClient{
		backend: backend,
	}
}

// DatasetClient is a client to access datasets API
type DatasetClient struct {
	backend Backend
}

// New creates a new dataset
func (c *DatasetClient) New(ctx context.Context, params *DatasetParams) (*Dataset, error) {
	if params.Project == nil || *params.Project == "" {
		return nil, fmt.Errorf("project must be specified")
	}
	path := fmt.Sprintf("/v1/projects/%s/datasets", *params.Project)
	method := http.MethodPost
	var response Dataset
	status, err := c.backend.Request(ctx, method, path, params, &response)
	switch status {
	case http.StatusOK:
	case http.StatusTooManyRequests:
		return nil, RateError
	default:
		return nil, fmt.Errorf("unexpected API response status: %d", status)
	}
	return &response, err
}
