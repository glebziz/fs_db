package cleaner

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model/sequence"
	"github.com/glebziz/fs_db/internal/usecase/cleaner/mocks"
)

const (
	testContentId  = "testContentId"
	testContentId2 = "testContentId2"
	testContentId3 = "testContentId3"
	testParent     = "testParent"
	testParent2    = "testParent2"
	testSeq        = sequence.Seq(1)
)

type prepareFunc func(td *testDeps)

type testDeps struct {
	core   *mock_cleaner.Mockcore
	cRepo  *mock_cleaner.MockcontentRepository
	cfRepo *mock_cleaner.MockcontentFileRepository
	db     *mock_cleaner.MockdbProvider
	dRepo  *mock_cleaner.MockdirRepository
	fRepo  *mock_cleaner.MockfileRepository
	sender *mock_cleaner.Mocksender
	txRepo *mock_cleaner.MocktransactionRepository
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	return &testDeps{
		core:   mock_cleaner.NewMockcore(ctrl),
		cRepo:  mock_cleaner.NewMockcontentRepository(ctrl),
		cfRepo: mock_cleaner.NewMockcontentFileRepository(ctrl),
		db:     mock_cleaner.NewMockdbProvider(ctrl),
		dRepo:  mock_cleaner.NewMockdirRepository(ctrl),
		fRepo:  mock_cleaner.NewMockfileRepository(ctrl),
		sender: mock_cleaner.NewMocksender(ctrl),
		txRepo: mock_cleaner.NewMocktransactionRepository(ctrl),
	}
}

func (td *testDeps) newUseCase() *useCase {
	return New(
		td.core, td.cRepo,
		td.cfRepo, td.db,
		td.dRepo, td.fRepo,
		td.sender, td.txRepo,
	)
}
