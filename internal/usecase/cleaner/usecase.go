package cleaner

import (
	"context"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	"github.com/glebziz/fs_db/internal/utils/wpool"
)

//go:generate mockgen -source usecase.go -destination mocks/mocks.go -typed true

type core interface {
	DeleteOld(ctx context.Context, txId string, beforeSeq sequence.Seq) []model.File
}

type contentRepository interface {
	Delete(ctx context.Context, path string) error
}

type contentFileRepository interface {
	Get(ctx context.Context, id string) (model.ContentFile, error)
	Delete(ctx context.Context, id string) error
}

type dbProvider interface {
	GC()
}

type dirRepository interface {
	Add(ctx context.Context, dir model.Dir) error
}

type fileRepository interface {
	Delete(ctx context.Context, file model.File) error
}

type sender interface {
	Send(ctx context.Context, event wpool.Event)
}

type transactionRepository interface {
	Oldest(ctx context.Context) (model.Transaction, error)
}

type UseCase struct {
	core   core
	cRepo  contentRepository
	cfRepo contentFileRepository
	db     dbProvider
	dRepo  dirRepository
	fRepo  fileRepository
	sender sender
	txRepo transactionRepository
}

func New(
	core core, cRepo contentRepository,
	cfRepo contentFileRepository, db dbProvider,
	dirRepo dirRepository, fRepo fileRepository,
	sender sender, txRepo transactionRepository,
) *UseCase {
	return &UseCase{
		core: core, cRepo: cRepo,
		cfRepo: cfRepo, db: db,
		dRepo: dirRepo, fRepo: fRepo,
		sender: sender, txRepo: txRepo,
	}
}
