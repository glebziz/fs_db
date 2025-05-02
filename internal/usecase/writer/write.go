package writer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/glebziz/fs_db/internal/model"
)

func (u *UseCase) Write(ctx context.Context, path string, contents model.Contents) (err error) {
	cw, err := u.cRepo.Create(ctx, path)
	if err != nil {
		return fmt.Errorf("content repository open: %w", err)
	}

	var (
		cr readWrapper

		w = bufWriter{w: cw}
	)
	defer func() {
		if !errors.Is(err, model.ErrNotEnoughSpace) {
			cw.Close()
			return
		}

		_, seekErr := cw.Seek(0, io.SeekEnd)
		if seekErr != nil {
			err = fmt.Errorf("seek: %w", seekErr)
			return
		}

		err = model.NotEnoughSpaceError{
			Err:    err,
			Start:  cw,
			Middle: bytes.NewReader(w.buf),
			End:    cr,
		}
	}()

	for content := range contents {
		_, err = content.Seek.Apply(cw)
		if err != nil {
			return fmt.Errorf("seek apply: %w", err)
		}

		// Trim the implementation of io.Reader to a single Read method.
		cr = readWrapper{content.Reader}
		_, err = io.Copy(&w, cr)
		if err != nil {
			return fmt.Errorf("copy: %w", err)
		}
	}

	return nil
}
