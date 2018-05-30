package httptrace

import (
	"net/http"
)

// Patch :
func Patch() (teardown func()) {
	previousTransport := http.DefaultTransport
	previousClient := http.DefaultClient

	http.DefaultTransport = DefaultTransport
	http.DefaultClient = DefaultClient
	return func() {
		http.DefaultTransport = previousTransport
		http.DefaultClient = previousClient
	}
}
