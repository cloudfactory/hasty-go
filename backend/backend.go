package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const defaultBackendAPI = "https://api.hasty.ai/"
const headerAPIKey = "X-API-Key"
const headerContentType = "Content-Type"
const contentTypeJSON = "application/json"

// DefaultBackend returns defaut Hasty backend that can execute API calls
func DefaultBackend() *Backend {
	return &Backend{
		API:    defaultBackendAPI,
		client: http.DefaultClient,
	}
}

// Backend points client to specific endpoint. In most cases, there's no need to configure special backend, but it may
// be needed if client should work against custom installation or mock
type Backend struct {
	API    string // API endpoint
	key    string
	client *http.Client
}

// SetAPIKey is used to set/update API key for requests
func (b *Backend) SetAPIKey(key string) {
	b.key = key
}

// Request performs an arbitrary HTTP Request, sending payload marshaled to JSON, and unmarshaling
// the response into respective variable (if provided).
func (b *Backend) Request(ctx context.Context, method, path string, payload, response interface{}) (int, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf("%s/%s", b.API, path)
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	req.Header[headerContentType] = []string{contentTypeJSON}
	req.Header[headerAPIKey] = []string{contentTypeJSON}

	resp, err := b.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO: evaluate if stastus should be returned together with error or not
		return resp.StatusCode, err
	}
	if response == nil {
		return resp.StatusCode, nil
	}
	err = json.Unmarshal(body, response)
	// TODO: evaluate if stastus should be returned together with error or not
	return resp.StatusCode, err
}
