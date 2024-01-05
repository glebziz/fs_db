package cleaner

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/glebziz/fs_db/internal/model"
)

const (
	sendTimeout = 15 * time.Second
)

//go:generate mockgen -source cleaner.go -destination mocks/mocks.go -typed true

type contentRepository interface {
	Delete(ctx context.Context, path string) error
}

type contentFileRepository interface {
	Delete(ctx context.Context, ids []string) ([]model.ContentFile, error)
}

type Cleaner struct {
	rWg    sync.WaitGroup
	sWg    sync.WaitGroup
	ch     chan []string
	closed atomic.Bool

	cRepo  contentRepository
	cfRepo contentFileRepository
}

func New(cRepo contentRepository, cfRepo contentFileRepository) *Cleaner {
	return &Cleaner{
		ch:     make(chan []string, 100),
		cRepo:  cRepo,
		cfRepo: cfRepo,
	}
}
