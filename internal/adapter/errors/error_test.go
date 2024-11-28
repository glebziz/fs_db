package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glebziz/fs_db"
	store "github.com/glebziz/fs_db/internal/proto"
)

type testMessage struct{}

func (tm testMessage) Reset() {}

func (tm testMessage) String() string {
	return ""
}

func (tm testMessage) ProtoMessage() {}

func (tm testMessage) Marshal() ([]byte, error) {
	return nil, nil
}

func TestError(t *testing.T) {
	for _, tc := range []struct {
		name    string
		err     error
		errCode store.ErrorCode
		code    codes.Code
	}{
		{
			name:    "size",
			err:     fs_db.ErrNoFreeSpace,
			errCode: store.ErrorCode_ErrNoFreeSpace,
			code:    codes.ResourceExhausted,
		},
		{
			name:    "not found",
			err:     fs_db.ErrNotFound,
			errCode: store.ErrorCode_ErrNotFound,
			code:    codes.NotFound,
		},
		{
			name:    "empty key error",
			err:     fs_db.ErrEmptyKey,
			errCode: store.ErrorCode_ErrEmptyKey,
			code:    codes.InvalidArgument,
		},
		{
			name:    "tx not found",
			err:     fs_db.ErrTxNotFound,
			errCode: store.ErrorCode_ErrTxNotFound,
			code:    codes.Aborted,
		},
		{
			name:    "tx already exists",
			err:     fs_db.ErrTxAlreadyExists,
			errCode: store.ErrorCode_ErrTxAlreadyExists,
			code:    codes.AlreadyExists,
		},
		{
			name:    "serialization",
			err:     fs_db.ErrTxSerialization,
			errCode: store.ErrorCode_ErrTxSerialization,
			code:    codes.FailedPrecondition,
		},
		{
			name:    "header not found",
			err:     fs_db.ErrHeaderNotFound,
			errCode: store.ErrorCode_ErrHeaderNotFound,
			code:    codes.Internal,
		},
		{
			name:    "other error",
			err:     assert.AnError,
			errCode: store.ErrorCode_ErrUnknown,
			code:    codes.Internal,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := Error(tc.err)

			st := status.Convert(err)
			require.Equal(t, tc.code, st.Code())
			require.Equal(t, tc.err.Error(), st.Message())
			require.Len(t, st.Details(), 1)

			e, ok := st.Details()[0].(*store.Error)
			require.True(t, ok)
			require.Equal(t, tc.errCode, e.GetCode())
		})
	}
}

func TestClientError(t *testing.T) {
	const (
		errMessage = "some string"
	)

	for _, tc := range []struct {
		name    string
		err     func(t *testing.T) error
		wantErr error
	}{
		{
			name: "invalid arguments",
			err: func(*testing.T) error {
				return status.Error(codes.InvalidArgument, errMessage)
			},
			wantErr: fs_db.ErrEmptyKey,
		},
		{
			name: "not found",
			err: func(*testing.T) error {
				return status.Error(codes.NotFound, errMessage)
			},
			wantErr: fs_db.ErrNotFound,
		},
		{
			name: "already exists",
			err: func(*testing.T) error {
				return status.Error(codes.AlreadyExists, errMessage)
			},
			wantErr: fs_db.ErrTxAlreadyExists,
		},
		{
			name: "resource exhausted",
			err: func(*testing.T) error {
				return status.Error(codes.ResourceExhausted, errMessage)
			},
			wantErr: fs_db.ErrNoFreeSpace,
		},
		{
			name: "failed precondition",
			err: func(*testing.T) error {
				return status.Error(codes.FailedPrecondition, errMessage)
			},
			wantErr: fs_db.ErrTxSerialization,
		},
		{
			name: "aborted",
			err: func(*testing.T) error {
				return status.Error(codes.Aborted, errMessage)
			},
			wantErr: fs_db.ErrTxNotFound,
		},
		{
			name: "internal",
			err: func(*testing.T) error {
				return status.Error(codes.Internal, errMessage)
			},
			wantErr: fs_db.ErrUnknown,
		},
		{
			name: "other error",
			err: func(*testing.T) error {
				return status.Error(codes.Unknown, errMessage)
			},
			wantErr: fs_db.ErrUnknown,
		},
		{
			name: "ErrUnknown",
			err: func(t *testing.T) error {
				st, err := status.New(codes.Unknown, errMessage).
					WithDetails(testMessage{}, &store.Error{
						Code: store.ErrorCode_ErrUnknown,
					})
				require.NoError(t, err)

				return st.Err()
			},
			wantErr: fs_db.ErrUnknown,
		},
		{
			name: "ErrNoFreeSpace",
			err: func(t *testing.T) error {
				st, err := status.New(codes.Unknown, errMessage).
					WithDetails(&store.Error{
						Code: store.ErrorCode_ErrNoFreeSpace,
					})
				require.NoError(t, err)

				return st.Err()
			},
			wantErr: fs_db.ErrNoFreeSpace,
		},
		{
			name: "ErrNotFound",
			err: func(t *testing.T) error {
				st, err := status.New(codes.Unknown, errMessage).
					WithDetails(&store.Error{
						Code: store.ErrorCode_ErrNotFound,
					})
				require.NoError(t, err)

				return st.Err()
			},
			wantErr: fs_db.ErrNotFound,
		},
		{
			name: "ErrEmptyKey",
			err: func(t *testing.T) error {
				st, err := status.New(codes.Unknown, errMessage).
					WithDetails(&store.Error{
						Code: store.ErrorCode_ErrEmptyKey,
					})
				require.NoError(t, err)

				return st.Err()
			},
			wantErr: fs_db.ErrEmptyKey,
		},
		{
			name: "ErrHeaderNotFound",
			err: func(t *testing.T) error {
				st, err := status.New(codes.Unknown, errMessage).
					WithDetails(&store.Error{
						Code: store.ErrorCode_ErrHeaderNotFound,
					})
				require.NoError(t, err)

				return st.Err()
			},
			wantErr: fs_db.ErrHeaderNotFound,
		},
		{
			name: "ErrTxNotFound",
			err: func(t *testing.T) error {
				st, err := status.New(codes.Unknown, errMessage).
					WithDetails(&store.Error{
						Code: store.ErrorCode_ErrTxNotFound,
					})
				require.NoError(t, err)

				return st.Err()
			},
			wantErr: fs_db.ErrTxNotFound,
		},
		{
			name: "ErrTxAlreadyExists",
			err: func(t *testing.T) error {
				st, err := status.New(codes.Unknown, errMessage).
					WithDetails(&store.Error{
						Code: store.ErrorCode_ErrTxAlreadyExists,
					})
				require.NoError(t, err)

				return st.Err()
			},
			wantErr: fs_db.ErrTxAlreadyExists,
		},
		{
			name: "ErrTxSerialization",
			err: func(t *testing.T) error {
				st, err := status.New(codes.Unknown, errMessage).
					WithDetails(&store.Error{
						Code: store.ErrorCode_ErrTxSerialization,
					})
				require.NoError(t, err)

				return st.Err()
			},
			wantErr: fs_db.ErrTxSerialization,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := ClientError(tc.err(t))

			require.ErrorIs(t, err, tc.wantErr)
		})
	}
}
