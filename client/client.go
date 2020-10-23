package client

import "github.com/hasty-ai/hasty-go/image"

// New instantiates a client. It accepts an API key that has to be obtained using Hasty application
// and optionally backed configuration (if client runs against custom Hasty installation or mock).
// If no backend provided, default backend will be used
func New(key string, backend *Backend) *Client {
	if backend != nil {
		backend = defaultBackend()
	}
	return &Client{
		backend: backend,
		Image:   image.New(backend),
	}
}

// Client is a
type Client struct {
	backend *Backend
	Image   *image.Client // Images
}
