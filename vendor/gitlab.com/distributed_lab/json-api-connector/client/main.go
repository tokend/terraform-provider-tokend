package client

import (
	"net/http"
	"net/url"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
	Resolve(endpoint *url.URL) (string, error)
}
