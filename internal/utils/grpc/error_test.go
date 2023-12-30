package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glebziz/fs_db/pkg/model"
)

func TestError(t *testing.T) {
	for _, tc := range []struct {
		name string
		err  error
		code codes.Code
	}{
		{
			name: "size",
			err:  model.SizeErr,
			code: codes.ResourceExhausted,
		},
		{
			name: "not found",
			err:  model.NotFoundErr,
			code: codes.NotFound,
		},
		{
			name: "header not found",
			err:  model.EmptyKeyErr,
			code: codes.InvalidArgument,
		},
		{
			name: "header not found",
			err:  model.HeaderNotFoundErr,
			code: codes.Internal,
		},
		{
			name: "other error",
			err:  assert.AnError,
			code: codes.Internal,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := Error(tc.err)

			st := status.Convert(err)
			require.Equal(t, tc.code, st.Code())
			require.Equal(t, tc.err.Error(), st.Message())
		})
	}
}

func TestClientError(t *testing.T) {
	const (
		errMessage = "some string"
	)

	for _, tc := range []struct {
		name    string
		code    codes.Code
		wantErr error
	}{
		{
			name:    "invalid arguments",
			code:    codes.InvalidArgument,
			wantErr: model.EmptyKeyErr,
		},
		{
			name:    "not found",
			code:    codes.NotFound,
			wantErr: model.NotFoundErr,
		},
		{
			name:    "resource exhausted",
			code:    codes.ResourceExhausted,
			wantErr: model.SizeErr,
		},
		{
			name:    "other error",
			code:    codes.OK,
			wantErr: status.Error(codes.OK, errMessage),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := ClientError(status.Error(tc.code, errMessage))

			require.ErrorIs(t, err, tc.wantErr)
		})
	}

	t.Run("internal", func(t *testing.T) {
		t.Parallel()

		err := ClientError(status.Error(codes.Internal, errMessage))

		require.Error(t, err)
		require.Equal(t, "some string", err.Error())
	})
}
