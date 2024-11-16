package model

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidFileFormat = errors.New("invalid file format")
)

type NotEnoughSpaceError struct {
	Err error

	Start  io.ReadCloser
	Middle io.Reader
	End    io.Reader
}

func (e NotEnoughSpaceError) Error() string {
	return fmt.Sprintf("not enough space: %v", e.Err)
}

func (e NotEnoughSpaceError) Unwrap() error {
	return e.Err
}

func (e NotEnoughSpaceError) Reader() io.Reader {
	return io.MultiReader(e.Start, e.Middle, e.End)
}

func (e NotEnoughSpaceError) Close() error {
	return e.Start.Close()
}
