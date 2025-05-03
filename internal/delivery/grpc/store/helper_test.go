package store

import (
	"testing"

	"go.uber.org/mock/gomock"

	mock_store "github.com/glebziz/fs_db/internal/delivery/grpc/store/mocks"
	_ "github.com/glebziz/fs_db/internal/utils/log"
)

type prepareFunc func(td *testDeps)

type testDeps struct {
	t        *testing.T
	suc      *mock_store.MockstoreUseCase
	txuc     *mock_store.MocktxUseCase
	r        *mock_store.MockReadSeekCloser
	w        *mock_store.MockReadWriteSeekCloser
	sStream  *mock_store.MockStoreV1_SetFileServer
	s2Stream *mock_store.MockStoreV1_SetFileV2Server
	gStream  *mock_store.MockStoreV1_GetFileServer
	g2Stream *mock_store.MockStoreV1_GetFileV2Server
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	return &testDeps{
		t:        t,
		suc:      mock_store.NewMockstoreUseCase(ctrl),
		txuc:     mock_store.NewMocktxUseCase(ctrl),
		r:        mock_store.NewMockReadSeekCloser(ctrl),
		w:        mock_store.NewMockReadWriteSeekCloser(ctrl),
		sStream:  mock_store.NewMockStoreV1_SetFileServer(ctrl),
		s2Stream: mock_store.NewMockStoreV1_SetFileV2Server(ctrl),
		gStream:  mock_store.NewMockStoreV1_GetFileServer(ctrl),
		g2Stream: mock_store.NewMockStoreV1_GetFileV2Server(ctrl),
	}
}

func (td *testDeps) newService() *Service {
	return New(td.suc, td.txuc)
}
