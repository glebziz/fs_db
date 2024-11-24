package async

import (
	"bytes"
	"errors"
	"sync"
	"sync/atomic"
)

type readWriter struct {
	sync.WaitGroup

	m      sync.Mutex
	cv     *sync.Cond
	closed atomic.Bool

	errM sync.Mutex
	err  error

	buf bytes.Buffer
}

func NewReadWriter() *readWriter {
	var rw readWriter
	rw.cv = sync.NewCond(&rw.m)

	return &rw
}

func (rw *readWriter) Read(p []byte) (n int, err error) {
	rw.m.Lock()
	defer rw.m.Unlock()

	if !rw.closed.Load() && rw.buf.Len() == 0 {
		rw.cv.Wait()
	}

	return rw.buf.Read(p)
}

func (rw *readWriter) Write(p []byte) (n int, err error) {
	rw.m.Lock()
	defer func() {
		rw.m.Unlock()
		rw.cv.Signal()
	}()

	err = rw.checkErr()
	if err != nil {
		return 0, err
	}

	return rw.buf.Write(p)
}

func (rw *readWriter) Close() error {
	rw.closed.Store(true)
	rw.cv.Broadcast()
	rw.Wait()

	return rw.err
}

func (rw *readWriter) SetError(err error) {
	rw.errM.Lock()
	defer rw.errM.Unlock()

	rw.err = errors.Join(rw.err, err)
}

func (rw *readWriter) checkErr() error {
	rw.errM.Lock()
	defer rw.errM.Unlock()

	return rw.err
}
