package http

import (
	"io"
	"net/http"
	"os"
)

// Patch :
func Patch(w io.Writer) (teardown func()) {
	previousTransport := http.DefaultTransport
	previousClient := http.DefaultClient

	transport := WrapTransport(previousTransport, os.Stderr)
	http.DefaultTransport = transport
	http.DefaultClient = &http.Client{Transport: transport}
	return func() {
		http.DefaultTransport = previousTransport
		http.DefaultClient = previousClient
	}
}
