package server

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/glebziz/fs_db/internal/model"
)

func TestContextStreamInterceptor(t *testing.T) {
	t.Parallel()

	stream := newWrappedStream(nil, testCtx)
	err := ContextStreamInterceptor(testRequest, stream, nil, func(srv any, stream grpc.ServerStream) error {
		require.Equal(t, testRequest, srv)
		require.Equal(t, testTxId, model.GetTxId(stream.Context()))

		return nil
	})

	require.NoError(t, err)
}
