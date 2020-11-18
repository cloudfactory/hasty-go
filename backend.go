package hasty

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const defaultBackendAPI = "https://api.hasty.ai"
const headerAPIKey = "X-API-Key"
const headerRequestID = "X-Request-ID"
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
func (b *APIKeyBackend) Request(ctx context.Context, method, path string, payload, response interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("unable to marshal JSON: %w", err)
	}

	url := fmt.Sprintf("%s%s", b.Endpoint, path)
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("unable to create HTTP request: %w", err)
	}

	req.Header[headerContentType] = []string{contentTypeJSON}
	req.Header[headerAPIKey] = []string{b.key}

	resp, err := b.client.Do(req)
	if err != nil {
		return fmt.Errorf("unable to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read from HTTP response body: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusAccepted:
	case http.StatusNoContent:
	case http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusTooManyRequests,
		http.StatusInternalServerError:
		var apiErr Error
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("unexpected API response status: %d, body: %s", resp.StatusCode, string(body))
		}
		apiErr.Status = resp.StatusCode
		apiErr.RequestID = resp.Header.Get(headerRequestID)
		return apiErr
	default:
		return fmt.Errorf("unexpected API response status: %d, body: %s", resp.StatusCode, string(body))
	}

	if response == nil {
		return nil
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return fmt.Errorf("unable to unmarshal HTTP response body as JSON: %w", err)
	}
	return nil
}
