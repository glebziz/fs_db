package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Delete_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.fRepo.EXPECT().
		Store(gomock.Any(), model.File{
			Key:       testKey,
			TxId:      testTxId,
			ContentId: testContentId,
		}).
		Return(nil)

	uc := td.newUseCase()

	err := uc.Delete(testCtx, testKey)

	require.NoError(t, err)
}

func TestUseCase_Delete_Error(t *testing.T) {

	t.Parallel()

	td := newTestDeps(t)

	td.fRepo.EXPECT().
		Store(gomock.Any(), gomock.Any()).
		Return(assert.AnError)

	uc := td.newUseCase()

	err := uc.Delete(testCtx, testKey)

	require.ErrorIs(t, err, assert.AnError)
}
