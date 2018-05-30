package httptrace

import (
	"io"
	"net/http"
	"os"
)

type nopCloser struct {
	io.Writer
}

func (*nopCloser) Close() error { return nil }

// Constant :
func Constant(w io.Writer) func(*http.Request) (io.WriteCloser, error) {
	return func(_ *http.Request) (io.WriteCloser, error) {
		return &nopCloser{w}, nil
	}
}

// Output
var (
	StderrOutput = Constant(os.Stderr)
	StdoutOutput = Constant(os.Stdout)
)
