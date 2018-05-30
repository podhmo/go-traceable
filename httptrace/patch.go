package httptrace

import (
	"io"
	"net/http"
)

// Patch :
func Patch(w io.Writer) (teardown func()) {
	previousTransport := http.DefaultTransport
	previousClient := http.DefaultClient

	transport := WrapTransport(previousTransport, StderrOutput)
	http.DefaultTransport = transport
	http.DefaultClient = &http.Client{Transport: transport}
	return func() {
		http.DefaultTransport = previousTransport
		http.DefaultClient = previousClient
	}
}
