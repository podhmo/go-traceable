package httptrace

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

// DefaultClient :
var DefaultClient *http.Client

// DefaultTransport :
var DefaultTransport http.RoundTripper

func init() {
	path := os.Getenv("TRACE")
	if path == "" {
		DefaultTransport = http.DefaultTransport
		DefaultClient = http.DefaultClient
		return
	}

	if path != "" {
		transport := WrapTransport(http.DefaultTransport, StderrOutput)
		DefaultTransport = transport
		DefaultClient = &http.Client{Transport: transport}
	}
}

// Head is a wrapper around DefaultClient.Head
func Head(url string) (resp *http.Response, err error) {
	return DefaultClient.Head(url)
}

// Get is a wrapper around DefaultClient.Get
func Get(url string) (resp *http.Response, err error) {
	return DefaultClient.Get(url)
}

// Post is a wrapper around DefaultClient.Post
func Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return DefaultClient.Post(url, contentType, body)
}

// PostForm is a wrapper around DefaultClient.PostForm
func PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return DefaultClient.PostForm(url, data)
}
