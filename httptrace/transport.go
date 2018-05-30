package httptrace

import (
	"io"
	"net/http"
	"net/http/httputil"
)

// TraceTransport :
type TraceTransport struct {
	Transport http.RoundTripper
	Output    func(func(io.Writer) (*http.Response, error)) (*http.Response, error)
}

// RoundTrip :
func (t *TraceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.Output(func(w io.Writer) (*http.Response, error) {
		{
			b, err := httputil.DumpRequestOut(req, true)
			if err != nil {
				return nil, err
			}
			w.Write(b)
		}

		res, err := t.Transport.RoundTrip(req)
		if err != nil {
			return nil, err
		}

		{
			b, err := httputil.DumpResponse(res, true)
			if err != nil {
				return nil, err
			}
			w.Write(b)
		}
		return res, nil
	})
}

// Constant :
func Constant(w io.Writer) func() (io.Writer, error) {
	return func() (io.Writer, error) {
		return w, nil
	}
}

// WrapTransport :
func WrapTransport(base http.RoundTripper, getwriter func() (io.Writer, error)) http.RoundTripper {
	return &TraceTransport{
		Transport: http.DefaultTransport,
		Output: func(call func(io.Writer) (*http.Response, error)) (*http.Response, error) {
			w, err := getwriter()
			if err != nil {
				return nil, err
			}
			return call(w)
		},
	}
}
