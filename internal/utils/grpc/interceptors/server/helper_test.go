package server

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	testTxId     = gofakeit.UUID()
	testRequest  = gofakeit.UUID()
	testResponse = gofakeit.UUID()

	testCtx = metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		TxIdKey: testTxId,
	}))
)

type testStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *testStream) Context() context.Context {
	return s.ctx
}

func newTestStream() grpc.ServerStream {
	return &testStream{nil, context.Background()}
}
