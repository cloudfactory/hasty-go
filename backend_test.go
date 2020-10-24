package hasty

import (
	"context"
	"fmt"
	"io/ioutil"
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
}

func TestAPIKeyBackend(t *testing.T) {
	reqs := make(chan testReq, 1)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		reqs <- testReq{
			method: r.Method,
			path:   strings.TrimPrefix(r.URL.Path, "/"),
			body:   string(body),
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
		{"regular", http.MethodPut, "bar/baz/42", "body", &testResp{}},
		{"empty body", http.MethodGet, "foo", "", &testResp{}},
		{"with no resp expected", http.MethodDelete, "baz/42", "another body", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, err := b.Request(ctx, tt.method, tt.path, tt.payload, tt.response)
			if err != nil {
				t.Errorf("APIKeyBackend.Request() error = %v", err)
				return
			}
			if status != http.StatusOK {
				t.Errorf("Expected status 200 OK, got: %d", status)
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
