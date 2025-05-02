package store

import (
	"context"
	"net"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/glebziz/fs_db"
	mock_store "github.com/glebziz/fs_db/internal/delivery/grpc/store/mocks"
	store "github.com/glebziz/fs_db/internal/proto"
	_ "github.com/glebziz/fs_db/internal/utils/log"
)

var (
	testKey     = gofakeit.UUID()
	testContent = []byte("some content")

	testTxId            = gofakeit.UUID()
	testTxIsoLevel      = store.TxIsoLevel_ISO_LEVEL_READ_COMMITTED
	testLocalTxIsoLevel = fs_db.IsoLevelDefault
)

type prepareFunc func(td *testDeps)

type testDeps struct {
	t    *testing.T
	suc  *mock_store.MockstoreUseCase
	txuc *mock_store.MocktxUseCase
	r    *mock_store.MockReadSeekCloser

	client store.StoreV1Client
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	suc := mock_store.NewMockstoreUseCase(ctrl)
	txuc := mock_store.NewMocktxUseCase(ctrl)

	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	server := grpc.NewServer(grpc.MaxSendMsgSize(500))
	store.RegisterStoreV1Server(server, New(suc, txuc))
	go func() {
		err := server.Serve(lis)
		require.NoError(t, err)
	}()

	t.Cleanup(func() {
		err := lis.Close()
		require.NoError(t, err)
		server.Stop()
	})

	conn, err := grpc.NewClient(
		"passthrough:buf_dial",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)

	return &testDeps{
		t:      t,
		suc:    suc,
		txuc:   txuc,
		r:      mock_store.NewMockReadSeekCloser(ctrl),
		client: store.NewStoreV1Client(conn),
	}
}
