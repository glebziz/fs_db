package content

import (
	"context"
	"os"
)

//go:generate mockgen -source repository.go -package mocks -destination mocks/mocks.go -typed true

type osAdapter interface {
	Create(ctx context.Context, path string) (*os.File, error)
	Open(ctx context.Context, path string) (*os.File, error)
	Remove(ctx context.Context, path string) error
}

type Repo struct {
	os osAdapter
}

func New(os osAdapter) *Repo {
	return &Repo{
		os: os,
	}
}
