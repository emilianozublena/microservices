package routific

import (
	"net/http"
)

// Client holds the http client
var Client HTTPClient

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func init() {
	Client = &http.Client{}
}
