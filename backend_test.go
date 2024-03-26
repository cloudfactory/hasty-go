package hasty

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testResp struct {
	Foo string `json:"foo"`
}

type testReq struct {
	method string
	path   string
	body   string
	key    string
}

func TestAPIKeyBackend(t *testing.T) {
	reqs := make(chan testReq, 1)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		reqs <- testReq{
			method: r.Method,
			path:   r.URL.Path,
			body:   string(body),
			key:    strings.Join(r.Header["X-Api-Key"], ","),
		}
		fmt.Fprintln(w, `{"foo":"bar"}`)
	}))
	defer ts.Close()

	b := NewAPIKeyBackend("secret-api-key")
	b.Endpoint = ts.URL

	ctx := context.TODO()
	tests := []struct {
		name     string
		method   string
		path     string
		payload  string
		response interface{}
	}{
		{"regular", http.MethodPut, "/bar/baz/42", "body", &testResp{}},
		{"empty body", http.MethodGet, "/foo", "", &testResp{}},
		{"with no resp expected", http.MethodDelete, "/baz/42", "another body", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := b.Request(ctx, tt.method, tt.path, tt.payload, tt.response); err != nil {
				t.Errorf("APIKeyBackend.Request() error = %v", err)
				return
			}
			req := <-reqs
			_ = req
			if req.method != tt.method {
				t.Errorf("Got method = %s, want %s", req.method, tt.method)
			}
			if req.path != tt.path {
				t.Errorf("Got path = %s, want %s", req.path, tt.path)
			}
			if req.key != "secret-api-key" {
				t.Errorf("Got key = %s, want secret-api-key", req.key)
			}
			// Mind, it will be json wrapped in quotes!
			if req.body != fmt.Sprintf(`"%s"`, tt.payload) {
				t.Errorf("Got body = %s, want %s", req.body, tt.payload)
			}
			if tt.response == nil {
				return
			}
			if tt.response.(*testResp).Foo != "bar" {
				t.Errorf("Expected response to be bar, got: %s", tt.response.(*testResp).Foo)
				return
			}
		})
	}
}
