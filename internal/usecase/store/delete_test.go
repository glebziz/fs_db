package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUseCase_Delete_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.fRepo.EXPECT().
		Delete(gomock.Any(), testTxId, testKey).
		Return(nil)

	uc := td.newUseCase()

	err := uc.Delete(testCtx, testKey)

	require.NoError(t, err)
}

func TestUseCase_Delete_Error(t *testing.T) {

	t.Parallel()

	td := newTestDeps(t)

	td.fRepo.EXPECT().
		Delete(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(assert.AnError)

	uc := td.newUseCase()

	err := uc.Delete(testCtx, testKey)

	require.ErrorIs(t, err, assert.AnError)
}
