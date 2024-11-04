package cleaner

import (
	"context"
	"time"

	"github.com/reugn/go-quartz/quartz"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

const (
	cleanPeriod = time.Minute * 1
)

//go:generate mockgen -source cleaner.go -destination mocks/mocks.go -typed true

type scheduler interface {
	Start(context.Context)
	ScheduleJob(jobDetail *quartz.JobDetail, trigger quartz.Trigger) error
	Clear() error
	Wait(context.Context)
	Stop()
}

type contentRepository interface {
	Delete(ctx context.Context, path string) error
}

type contentFileRepository interface {
	Get(ctx context.Context, id string) (model.ContentFile, error)
	Delete(ctx context.Context, id string) error
}

type fileRepository interface {
	DeleteOld(ctx context.Context, txId string, beforeSeq sequence.Seq)
}

type transactionRepository interface {
	Oldest(ctx context.Context) (*model.Transaction, error)
}

type Cleaner struct {
	sched  scheduler
	cRepo  contentRepository
	cfRepo contentFileRepository
	fRepo  fileRepository
	txRepo transactionRepository
}

func New(
	sched scheduler, cRepo contentRepository,
	cfRepo contentFileRepository, fRepo fileRepository,
	txRepo transactionRepository,
) *Cleaner {
	return &Cleaner{
		sched:  sched,
		cRepo:  cRepo,
		cfRepo: cfRepo,
		fRepo:  fRepo,
		txRepo: txRepo,
	}
}
