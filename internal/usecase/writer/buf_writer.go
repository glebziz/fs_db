package writer

import (
	"io"
)

type bufWriter struct {
	buf []byte
	w   io.Writer
}

func (w *bufWriter) Write(p []byte) (n int, err error) {
	if len(w.buf) < len(p) {
		w.buf = make([]byte, len(p))
	}

	w.buf = w.buf[:len(p)]
	copy(w.buf, p)

	return w.w.Write(w.buf)
}
