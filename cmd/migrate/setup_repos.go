package main

import (
	"context"
	"fmt"
	"io"

	"github.com/glebziz/fs_db/internal/db/badger"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/transactor"
	contentFile "github.com/glebziz/fs_db/internal/repository/content_file"
	"github.com/glebziz/fs_db/internal/repository/file"
)

type fileRepo interface {
	transactor.Transactor
	Set(ctx context.Context, f model.File) error
}

type contentFileRepo interface {
	Store(ctx context.Context, f model.ContentFile) error
}

type repos struct {
	p io.Closer

	fRepo  fileRepo
	cfRepo contentFileRepo
}

func setupRepos(badgerPath string) (repos, error) {
	p, err := badger.New(badgerPath)
	if err != nil {
		return repos{}, fmt.Errorf("badger new: %w", err)
	}

	return repos{
		p:      p,
		fRepo:  file.New(p),
		cfRepo: contentFile.New(p),
	}, nil
}
