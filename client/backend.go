package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bitly/go-simplejson"
)

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

// Request performs an arbitrary HTTP Request, sending payload marshaled to JSON, and unmarshaling
// the response into respective variable (if provided).
func (b *Backend) Request(method, path, key string, payload, response interface{}) (int, error) {
	req, err := http.NewRequest(method, endpoint+path, payloadReader)
	if err != nil {
		panic(err)
	}
	req.Header["Content-Type"] = []string{contentType}

	var resp *http.Response
	var reqs = 0
	for {
		start := time.Now()
		if resp, err = client.Do(req); err != nil {
			fmt.Printf("* ERROR %.1fs %s %s %s\n", time.Since(start).Seconds(), method, path, err.Error())
			return nil, 0, err
		}
		reqID := resp.Header.Get(headerRequestID)
		if reqID != "" && len(reqID) >= 8 {
			reqID = reqID[:8] // Cut the UUID
		}
		fmt.Printf("* %s %d %.1fs %s %s \n", reqID, resp.StatusCode, time.Since(start).Seconds(), method, path)
		if resp.StatusCode != http.StatusTooManyRequests || reqs >= httpMaxRetries {
			break
		}
		// Retry only on 429 and only httpMaxRetries times
		_, _ = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		reqs++
		time.Sleep(httpRetryInterval)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	j, err := simplejson.NewFromReader(bytes.NewReader(b))
	if err == io.EOF {
		return simplejson.New(), resp.StatusCode, nil
	}

	if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// This may help if we print some debug data
		fmt.Printf("[STATUS %d] %s %s\n%s\n", resp.StatusCode, method, path, string(b))
	}

	if err != nil {
		return nil, resp.StatusCode, err
	}

	return j, resp.StatusCode, nil
}
