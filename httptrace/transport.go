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

// WrapTransport :
func WrapTransport(base http.RoundTripper, w io.Writer) http.RoundTripper {
	return &TraceTransport{
		Transport: http.DefaultTransport,
		Output: func(call func(io.Writer) (*http.Response, error)) (*http.Response, error) {
			return call(w)
		},
	}
}

// WrapTransportWithGetWriter :
func WrapTransportWithGetWriter(base http.RoundTripper, getwriter func() (io.Writer, error)) http.RoundTripper {
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
