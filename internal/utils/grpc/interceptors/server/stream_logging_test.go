package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestStreamLoggingInterceptor(t *testing.T) {
	for _, tc := range []struct {
		name    string
		handler grpc.StreamHandler
		wantErr error
	}{
		{
			name: "success",
			handler: func(srv any, stream grpc.ServerStream) error {
				require.Equal(t, testRequest, srv)

				return nil
			},
		},
		{
			name: "handler error",
			handler: func(srv any, stream grpc.ServerStream) error {
				require.Equal(t, testRequest, srv)

				return assert.AnError
			},
			wantErr: assert.AnError,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := StreamLoggingInterceptor(testRequest, newTestStream(), &grpc.StreamServerInfo{
				FullMethod:     "testService/test",
				IsServerStream: true,
			}, tc.handler)

			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
