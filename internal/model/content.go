package model

import (
	"io"
)

type Content struct {
	Size   uint64
	Reader io.ReadCloser
}
