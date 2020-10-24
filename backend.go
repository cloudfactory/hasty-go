package hasty

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

// NewAPIKeyBackend returns backend configured to Hasty production. In most cases, there's no need to configure anything
// else, but it may be needed to change endpoint if client should work against custom installation or mocked Hasty API
func NewAPIKeyBackend(key string) *APIKeyBackend {
	return &APIKeyBackend{
		Endpoint: defaultBackendAPI,
		client:   http.DefaultClient,
		key:      key,
	}
}

// APIKeyBackend allows client to execute authorised HTTP calls against specific endpoint using API key authorisation
type APIKeyBackend struct {
	Endpoint string // API endpoint
	key      string
	client   *http.Client
}

// Request performs an arbitrary HTTP Request, sending payload marshaled to JSON, and unmarshaling
// the response into respective variable (if provided).
// TODO: evaluate if stastus should be returned together with error or not in case HTTP call actually has happened
func (b *APIKeyBackend) Request(ctx context.Context, method, path string, payload, response interface{}) (int, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf("%s/%s", b.Endpoint, path)
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
		return resp.StatusCode, err
	}
	if response == nil {
		return resp.StatusCode, nil
	}
	err = json.Unmarshal(body, response)
	return resp.StatusCode, err
}
