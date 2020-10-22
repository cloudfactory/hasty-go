package client

const defaultBackendAPI = "https://api.hasty.ai/"

// Backend points client to specific endpoint. In most cases, there's no need to configure special backend, but it may
// be needed if client should work against custom installation or mock
type Backend struct {
	API string // API endpoint
}

func defaultBackend() *Backend {
	return &Backend{
		API: defaultBackendAPI,
	}
}
