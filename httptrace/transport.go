package httptrace

import (
	"io"
	"net/http"
	"net/http/httputil"
)

// TraceTransport :
type TraceTransport struct {
	Transport http.RoundTripper
	GetWriter func(*http.Request) (io.WriteCloser, error)
}

// RoundTrip :
func (t *TraceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	w, err := t.GetWriter(req)
	if err != nil {
		return nil, err
	}
	defer w.Close()

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
}

// WrapTransport :
func WrapTransport(base http.RoundTripper, getwriter func(*http.Request) (io.WriteCloser, error)) http.RoundTripper {
	return &TraceTransport{
		Transport: base,
		GetWriter: getwriter,
	}
}
