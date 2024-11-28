package di

import (
	"math/rand/v2"

	"github.com/glebziz/fs_db/config"
	"github.com/glebziz/fs_db/internal/db/badger"
	storeService "github.com/glebziz/fs_db/internal/delivery/grpc/store"
	contentRepo "github.com/glebziz/fs_db/internal/repository/content"
	contentFileRepo "github.com/glebziz/fs_db/internal/repository/content_file"
	dirRepo "github.com/glebziz/fs_db/internal/repository/dir"
	fileRepo "github.com/glebziz/fs_db/internal/repository/file"
	transactionRepo "github.com/glebziz/fs_db/internal/repository/transaction"
	"github.com/glebziz/fs_db/internal/usecase/cleaner"
	"github.com/glebziz/fs_db/internal/usecase/core"
	"github.com/glebziz/fs_db/internal/usecase/dir"
	"github.com/glebziz/fs_db/internal/usecase/store"
	"github.com/glebziz/fs_db/internal/usecase/transaction"
	"github.com/glebziz/fs_db/internal/utils/generator"
	"github.com/glebziz/fs_db/internal/utils/wpool"
)

type Container struct {
	cfg config.Config

	badger *badger.Manager
	pool   *wpool.Pool
	gen    *generator.Gen
	rand   *rand.Rand

	storeService *storeService.Service

	cleanerUseCase     *cleaner.UseCase
	coreUseCase        *core.UseCase
	dirUseCase         *dir.UseCase
	storeUseCase       *store.UseCase
	transactionUseCase *transaction.UseCase

	contentRepo     *contentRepo.Repo
	contentFileRepo *contentFileRepo.Repo
	dirRepo         *dirRepo.Repo
	fileRepo        *fileRepo.Repo
	transactionRepo *transactionRepo.Repo
}

func New(cfg config.Config) *Container {
	return &Container{
		cfg: cfg,
	}
}
