package httptrace

import (
	"io"
	"os"
)

// Constant :
func Constant(w io.Writer) func() (io.Writer, error) {
	return func() (io.Writer, error) {
		return w, nil
	}
}

// Output
var (
	StderrOutput = Constant(os.Stderr)
	StdoutOutput = Constant(os.Stdout)
)

// FileOutput :
func FileOutput() {
}
