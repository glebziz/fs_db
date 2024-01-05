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

func TestImplementation_BeginTx_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.txuc.EXPECT().
		Begin(gomock.Any(), testLocalTxIsoLevel).
		Return(testTxId, nil)

	resp, err := td.client.BeginTx(context.Background(), &store.BeginTxRequest{
		IsoLevel: testTxIsoLevel,
	})

	require.NoError(t, err)
	require.Equal(t, testTxId, resp.Id)
}

func TestImplementation_BeginTx_Error(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.txuc.EXPECT().
		Begin(gomock.Any(), gomock.Any()).
		Return("", assert.AnError)

	resp, err := td.client.BeginTx(context.Background(), &store.BeginTxRequest{
		IsoLevel: testTxIsoLevel,
	})

	st := status.Convert(err)

	require.Error(t, err)
	require.Equal(t, codes.Internal, st.Code())
	require.Nil(t, resp)
}
