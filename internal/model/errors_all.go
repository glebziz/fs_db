//go:build !windows

package model

import (
	"syscall"
)

const (
	ErrNotEnoughSpace = syscall.ENOSPC
	ErrPathNotFound   = syscall.ENOENT
)
