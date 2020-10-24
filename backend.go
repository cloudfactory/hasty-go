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

var _ Backend = &APIKeyBackend{}

// NewAPIKeyBackend returns backend configured to Hasty production. In most cases, there's no need to configure anything
// else, but it may be needed to change endpoint if client should work against custom installation or mocked Hasty API
func NewAPIKeyBackend(key string) *APIKeyBackend {
	return &APIKeyBackend{
		Endpoint: defaultBackendAPI,
		client:   http.DefaultClient,
		key:      key,
	}
}

// APIKeyBackend allows client to execute authorized HTTP calls against specific endpoint using API key authorization
type APIKeyBackend struct {
	Endpoint string // API endpoint
	key      string
	client   *http.Client
}

// Request performs an arbitrary HTTP Request, sending payload marshaled to JSON, and unmarshalling
// the response into respective variable (if provided).
// TODO: evaluate if stastus should be returned together with error or not in case HTTP call actually has happened
func (b *APIKeyBackend) Request(ctx context.Context, method, path string, payload, response interface{}) (int, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("unable to marshal JSON: %w", err)
	}

	url := fmt.Sprintf("%s/%s", b.Endpoint, path)
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	if err != nil {
		return 0, fmt.Errorf("unable to create HTTP request: %w", err)
	}

	req.Header[headerContentType] = []string{contentTypeJSON}
	req.Header[headerAPIKey] = []string{contentTypeJSON}

	resp, err := b.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("unable to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if response == nil {
		return resp.StatusCode, nil
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("unable to read from HTTP response body: %w", err)
	}
	err = json.Unmarshal(body, response)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("unable to unmarshal HTTP response body as JSON: %w", err)
	}
	return resp.StatusCode, nil
}
