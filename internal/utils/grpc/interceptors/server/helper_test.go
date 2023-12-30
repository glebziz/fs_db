package server

import (
	"context"

	"google.golang.org/grpc"
)

var (
	testRequest  = "testRequest"
	testResponse = "testResponse"
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
