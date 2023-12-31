package store

import (
	"bytes"
	"context"
	"io"
	"net"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/glebziz/fs_db/internal/delivery/grpc/store/mocks"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/log"
)

type errReader struct{}

func (r errReader) Read([]byte) (int, error) {
	return 0, assert.AnError
}

var (
	testKey        = gofakeit.UUID()
	testContent    = []byte("some content")
	testReader     = io.NopCloser(bytes.NewReader(testContent))
	testErrReader  = io.NopCloser(errReader{})
	testContentLen = uint64(len(testContent))
)

type testDeps struct {
	uc *mock_store.MockUsecase

	client store.StoreV1Client
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	uc := mock_store.NewMockUsecase(ctrl)

	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	server := grpc.NewServer(grpc.MaxSendMsgSize(500))
	store.RegisterStoreV1Server(server, New(uc))
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalln("error serving server", err)
		}
	}()

	t.Cleanup(func() {
		err := lis.Close()
		if err != nil {
			log.Fatalln("error closing listener", err)
		}
		server.Stop()
	})

	conn, err := grpc.Dial(
		"buf dial",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("error connecting to server", err)
	}

	return &testDeps{
		uc:     uc,
		client: store.NewStoreV1Client(conn),
	}
}
