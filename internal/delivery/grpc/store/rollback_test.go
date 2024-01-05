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

func TestImplementation_RollbackTx_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.txuc.EXPECT().
		Rollback(gomock.Any()).
		Return(nil)

	_, err := td.client.RollbackTx(context.Background(), &store.RollbackTxRequest{})

	require.NoError(t, err)
}

func TestImplementation_RollbackTx_Error(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.txuc.EXPECT().
		Rollback(gomock.Any()).
		Return(assert.AnError)

	_, err := td.client.RollbackTx(context.Background(), &store.RollbackTxRequest{})

	st := status.Convert(err)

	require.Error(t, err)
	require.Equal(t, codes.Internal, st.Code())
}
