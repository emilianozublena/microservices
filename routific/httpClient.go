package routific

import (
	"io"
	"net/http"
)

// Client holds the http client
type Client http.Client

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	NewRequest(method, url string, body io.Reader) (*http.Request, error)
}

// Do is a wrapper for Do in http package, useful for decoupling and mocking outbound calls
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.Do(req)
}

// NewRequest creates a new request by using the http package
func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	httpRequest, err := http.NewRequest(method, url, body)
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", "Bearer "+accessToken)
	return httpRequest, err
}
