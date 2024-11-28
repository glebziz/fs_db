package content

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/os"
)

type readWrapper struct {
	r io.Reader
}

func (r readWrapper) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}

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

func (r *Repo) Store(_ context.Context, path string, content io.Reader) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}

	w := bufWriter{w: f}
	content = readWrapper{content}

	defer func() {
		if errors.Is(err, os.ErrNotEnoughSpace) {
			err = model.NotEnoughSpaceError{
				Err:    err,
				Start:  f,
				Middle: bytes.NewReader(w.buf),
				End:    content,
			}
		} else {
			f.Close()
		}
	}()

	_, err = io.Copy(&w, content)
	if errors.Is(err, os.ErrNotEnoughSpace) {
		_, seekErr := f.Seek(0, io.SeekStart)
		if seekErr != nil {
			return fmt.Errorf("seek: %w", seekErr)
		}

		return err
	}
	if err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	return nil
}
