package hasty

import (
	"context"
	"fmt"
	"net/http"
)

// DatasetParams is parameters set for creating or updating a dataset
type DatasetParams struct {
	Project *string  `json:"-"`
	Name    *string  `json:"name"`
	Order   *float64 `json:"norder"`
}

// Dataset describes an dataser information that API may return
type Dataset struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Order float64 `json:"norder"`
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
	if err := c.backend.Request(ctx, method, path, params, &response); err != nil {
		return nil, fmt.Errorf("unable to create dataset: %w", err)
	}
	return &response, nil
}

// Update updates an existing dataset
func (c *DatasetClient) Update(ctx context.Context, id string, params *DatasetParams) (*Dataset, error) {
	if params.Project == nil || *params.Project == "" {
		return nil, fmt.Errorf("project must be specified")
	}
	path := fmt.Sprintf("/v1/projects/%s/datasets/%s", *params.Project, id)
	method := http.MethodPut
	var response Dataset
	if err := c.backend.Request(ctx, method, path, params, &response); err != nil {
		return nil, fmt.Errorf("unable to update dataset: %w", err)
	}
	return &response, nil
}

// Delete deletes an existing dataset
func (c *DatasetClient) Delete(ctx context.Context, id string, params *DatasetParams) error {
	if params.Project == nil || *params.Project == "" {
		return fmt.Errorf("project must be specified")
	}
	path := fmt.Sprintf("/v1/projects/%s/datasets/%s", *params.Project, id)
	method := http.MethodDelete
	if err := c.backend.Request(ctx, method, path, params, nil); err != nil {
		return fmt.Errorf("unable to delete dataset: %w", err)
	}
	return nil
}
