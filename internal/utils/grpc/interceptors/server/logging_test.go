package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestLoggingInterceptor(t *testing.T) {
	for _, tc := range []struct {
		name    string
		handler grpc.UnaryHandler
		expErr  error
		expResp any
	}{
		{
			name: "success",
			handler: func(ctx context.Context, req any) (any, error) {
				require.Equal(t, testRequest, req)

				return testResponse, nil
			},
			expResp: testResponse,
		},
		{
			name: "handler error",
			handler: func(ctx context.Context, req any) (any, error) {
				require.Equal(t, testRequest, req)

				return nil, assert.AnError
			},
			expErr: assert.AnError,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			resp, err := LoggingInterceptor(context.Background(), testRequest, &grpc.UnaryServerInfo{
				Server:     nil,
				FullMethod: "testService/test",
			}, tc.handler)

			if tc.expErr != nil {
				require.ErrorIs(t, err, tc.expErr)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.expResp, resp)
		})
	}
}

func TestSplitFullMethodName(t *testing.T) {
	for _, tc := range []struct {
		fullMethod string
		service    string
		method     string
	}{
		{
			fullMethod: "/service/method",
			service:    "service",
			method:     "method",
		},
		{
			fullMethod: "service/method",
			service:    "service",
			method:     "method",
		},
		{
			fullMethod: "//method",
			method:     "method",
		},
		{
			fullMethod: "some_text",
			service:    unknownValue,
			method:     unknownValue,
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("fullMethod: %s", tc.fullMethod), func(t *testing.T) {
			t.Parallel()

			service, method := splitFullMethodName(tc.fullMethod)

			require.Equal(t, tc.service, service)
			require.Equal(t, tc.method, method)
		})
	}
}
