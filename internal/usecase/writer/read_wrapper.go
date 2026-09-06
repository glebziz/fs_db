package writer

import (
	"io"
)

type readWrapper struct {
	r io.Reader
}

func (r readWrapper) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}
