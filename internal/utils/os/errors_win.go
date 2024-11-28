//go:build windows

package os

import (
	"syscall"
)

const (
	ErrNotEnoughSpace = syscall.Errno(112)
	ErrPathNotFound   = syscall.ERROR_PATH_NOT_FOUND
)
