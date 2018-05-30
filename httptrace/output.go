package httptrace

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
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

type fileOutput struct {
	i   int64
	Dir string
}

func (o *fileOutput) GetWriter(req *http.Request) (io.WriteCloser, error) {
	filename := o.FileName(req)
	log.Println("trace to", filename)
	return os.Create(filepath.Join(o.Dir, filename))
}

func (o *fileOutput) FileName(req *http.Request) string {
	i := atomic.AddInt64(&o.i, 1)
	return fmt.Sprintf("%04d%s", i, strings.Replace(req.URL.String(), "/", "@", -1))
}

// FileOutput :
func FileOutput(basedir string) (func(*http.Request) (io.WriteCloser, error), error) {
	if err := os.MkdirAll(basedir, 0744); err != nil {
		return nil, err
	}
	o := &fileOutput{
		Dir: basedir,
	}
	return o.GetWriter, nil
}

// MustFileOutput :
func MustFileOutput(basedir string) func(*http.Request) (io.WriteCloser, error) {
	getwriter, err := FileOutput(basedir)
	if err != nil {
		panic(err)
	}
	return getwriter
}
