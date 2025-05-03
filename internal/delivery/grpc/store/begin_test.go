package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glebziz/fs_db"
	store "github.com/glebziz/fs_db/internal/proto"
)

func TestService_BeginTx(t *testing.T) {
	t.Parallel()

	const (
		txId = "txId"
	)

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		resp    *store.BeginTxResponse
		errCode codes.Code
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.txuc.EXPECT().
					Begin(gomock.Any(), fs_db.IsoLevelReadCommitted).
					Return(txId, nil)
			},
			resp: &store.BeginTxResponse{
				Id: txId,
			},
		},
		{
			name: "Begin error",
			prepare: func(td *testDeps) {
				td.txuc.EXPECT().
					Begin(gomock.Any(), gomock.Any()).
					Return("", assert.AnError)
			},
			errCode: codes.Internal,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			s := td.newService()
			resp, err := s.BeginTx(context.Background(), &store.BeginTxRequest{
				IsoLevel: store.TxIsoLevel_ISO_LEVEL_READ_COMMITTED,
			})

			require.Equal(t, tc.errCode, status.Code(err))
			require.Equal(t, tc.resp, resp)
		})
	}
}

// func TestImplementation_BeginTx_Success(t *testing.T) {
// 	t.Parallel()
//
// 	td := newTestDeps(t)
//
// 	td.txuc.EXPECT().
// 		Begin(gomock.Any(), testLocalTxIsoLevel).
// 		Return(testTxId, nil)
//
// 	resp, err := td.client.BeginTx(context.Background(), &store.BeginTxRequest{
// 		IsoLevel: testTxIsoLevel,
// 	})
//
// 	require.NoError(t, err)
// 	require.Equal(t, testTxId, resp.Id)
// }
//
// func TestImplementation_BeginTx_Error(t *testing.T) {
// 	t.Parallel()
//
// 	td := newTestDeps(t)
//
// 	td.txuc.EXPECT().
// 		Begin(gomock.Any(), gomock.Any()).
// 		Return("", assert.AnError)
//
// 	resp, err := td.client.BeginTx(context.Background(), &store.BeginTxRequest{
// 		IsoLevel: testTxIsoLevel,
// 	})
//
// 	st := status.Convert(err)
//
// 	require.Error(t, err)
// 	require.Equal(t, codes.Internal, st.Code())
// 	require.Nil(t, resp)
// }
