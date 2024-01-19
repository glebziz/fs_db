package cleaner

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/usecase/cleaner/mocks"
)

var (
	testCtx        = context.Background()
	testContentId  = gofakeit.UUID()
	testContentId2 = gofakeit.UUID()
	testParent     = gofakeit.UUID()
	testTxCreateTs = time.Now().Add(-1 * time.Minute)
)

type prepareFunc func(td *testDeps) error

type testDeps struct {
	sched  *mock_cleaner.Mockscheduler
	cRepo  *mock_cleaner.MockcontentRepository
	cfRepo *mock_cleaner.MockcontentFileRepository
	fRepo  *mock_cleaner.MockfileRepository
	txRepo *mock_cleaner.MocktransactionRepository
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	return &testDeps{
		sched:  mock_cleaner.NewMockscheduler(ctrl),
		cRepo:  mock_cleaner.NewMockcontentRepository(ctrl),
		cfRepo: mock_cleaner.NewMockcontentFileRepository(ctrl),
		fRepo:  mock_cleaner.NewMockfileRepository(ctrl),
		txRepo: mock_cleaner.NewMocktransactionRepository(ctrl),
	}
}

func (d *testDeps) newCleaner() *Cleaner {
	return New(
		d.sched, d.cRepo,
		d.cfRepo, d.fRepo,
		d.txRepo,
	)
}
