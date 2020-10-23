package client

import (
	"github.com/hasty-ai/hasty-go/backend"
	"github.com/hasty-ai/hasty-go/image"
)

// New instantiates a client. It accepts an API key that has to be obtained using Hasty application
// and optionally backed configuration (if client runs against custom Hasty installation or mock).
// If no backend provided, default backend will be used
func New(key string, b *backend.Backend) *Client {
	if b == nil {
		b = backend.DefaultBackend()
	}
	b.SetAPIKey(key)
	return &Client{
		backend: b,
		Image:   image.New(b),
	}
}

// Client to access whole Hasty API
type Client struct {
	backend *backend.Backend
	Image   *image.Client
}
