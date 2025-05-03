package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	store "github.com/glebziz/fs_db/internal/proto"
)

func TestService_CommitTx(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		errCode codes.Code
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.txuc.EXPECT().
					Commit(gomock.Any()).
					Return(nil)
			},
		},
		{
			name: "Commit error",
			prepare: func(td *testDeps) {
				td.txuc.EXPECT().
					Commit(gomock.Any()).
					Return(assert.AnError)
			},
			errCode: codes.Internal,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			s := td.newService()
			_, err := s.CommitTx(context.Background(), &store.CommitTxRequest{})
			require.Equal(t, tc.errCode, status.Code(err))
		})
	}
}

// func TestImplementation_CommitTx_Success(t *testing.T) {
// 	t.Parallel()
//
// 	td := newTestDeps(t)
//
// 	td.txuc.EXPECT().
// 		Commit(gomock.Any()).
// 		Return(nil)
//
// 	_, err := td.client.CommitTx(context.Background(), &store.CommitTxRequest{})
//
// 	require.NoError(t, err)
// }
//
// func TestImplementation_CommitTx_Error(t *testing.T) {
// 	t.Parallel()
//
// 	td := newTestDeps(t)
//
// 	td.txuc.EXPECT().
// 		Commit(gomock.Any()).
// 		Return(assert.AnError)
//
// 	_, err := td.client.CommitTx(context.Background(), &store.CommitTxRequest{})
//
// 	st := status.Convert(err)
//
// 	require.Error(t, err)
// 	require.Equal(t, codes.Internal, st.Code())
// }
