package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glebziz/fs_db"
)

func TestError(t *testing.T) {
	for _, tc := range []struct {
		name string
		err  error
		code codes.Code
	}{
		{
			name: "size",
			err:  fs_db.SizeErr,
			code: codes.ResourceExhausted,
		},
		{
			name: "not found",
			err:  fs_db.NotFoundErr,
			code: codes.NotFound,
		},
		{
			name: "empty key error",
			err:  fs_db.EmptyKeyErr,
			code: codes.InvalidArgument,
		},
		{
			name: "tx not found",
			err:  fs_db.TxNotFoundErr,
			code: codes.Aborted,
		},
		{
			name: "tx already exists",
			err:  fs_db.TxAlreadyExistsErr,
			code: codes.AlreadyExists,
		},
		{
			name: "serialization",
			err:  fs_db.TxSerializationErr,
			code: codes.FailedPrecondition,
		},
		{
			name: "header not found",
			err:  fs_db.HeaderNotFoundErr,
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
			wantErr: fs_db.EmptyKeyErr,
		},
		{
			name:    "not found",
			code:    codes.NotFound,
			wantErr: fs_db.NotFoundErr,
		},
		{
			name:    "already exists",
			code:    codes.AlreadyExists,
			wantErr: fs_db.TxAlreadyExistsErr,
		},
		{
			name:    "resource exhausted",
			code:    codes.ResourceExhausted,
			wantErr: fs_db.SizeErr,
		},
		{
			name:    "failed precondition",
			code:    codes.FailedPrecondition,
			wantErr: fs_db.TxSerializationErr,
		},
		{
			name:    "aborted",
			code:    codes.Aborted,
			wantErr: fs_db.TxNotFoundErr,
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
