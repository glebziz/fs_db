//go:build !windows

package os

import (
	"syscall"
)

const (
	ErrNotEnoughSpace = syscall.ENOSPC
	ErrPathNotFound   = syscall.ENOENT
)
