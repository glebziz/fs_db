package transaction

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Begin_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.txRepo.EXPECT().
		Store(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, tx model.Transaction) error {
			require.Equal(t, testId, tx.Id)
			require.Equal(t, testIsoLevel, tx.IsoLevel)
			return nil
		})

	uc := td.newUseCase()

	txId, err := uc.Begin(context.Background(), testIsoLevel)

	require.NoError(t, err)
	require.Equal(t, testId, txId)
}

func TestUseCase_Begin_Error(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.txRepo.EXPECT().
		Store(gomock.Any(), gomock.Any()).
		Return(assert.AnError)

	uc := td.newUseCase()

	txId, err := uc.Begin(context.Background(), testIsoLevel)

	require.ErrorIs(t, err, assert.AnError)
	require.Zero(t, txId)
}
