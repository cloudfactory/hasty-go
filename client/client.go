package client

// New instantiates a client. It accepts an API key that has to be obtained using Hasty application
// and optionally backed configuration (if client runs against custom Hasty installation or mock).
// If no backend provided, default backend will be used
func New(key string, backend *Backend) *Client {
	c := &Client{
		backend: defaultBackend(),
	}
	if backend != nil {
		c.backend = backend
	}
	return c
}

// Client is a
type Client struct {
	backend *Backend
}
