package hasty

import (
	"context"
)

// NewClient is a suggested way to instantiate a new client with default backend. It accepts an API key that has to be
// obtained using Hasty application
func NewClient(key string) *Client {
	b := NewAPIKeyBackend(key)
	return NewClientWithBackend(b)
}

// NewClientWithBackend instantiates a client with provided backend
func NewClientWithBackend(b Backend) *Client {
	return &Client{
		backend: b,
		Image:   NewImageClient(b),
		Dataset: NewDatasetClient(b),
	}
}

// Backend can perform authenticated HTTP requests for client. It can be substituted with mock
// or with some custom authentication method
type Backend interface {
	Request(ctx context.Context, method, path string, payload, response interface{}) error
}

// Client to access whole Hasty API
type Client struct {
	backend Backend
	Image   *ImageClient
	Dataset *DatasetClient
}
