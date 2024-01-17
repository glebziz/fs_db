package transaction

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Rollback_Success(t *testing.T) {
	t.Parallel()

	var (
		contentIds = []string{gofakeit.UUID(), gofakeit.UUID()}
	)

	td := newTestDeps(t)

	td.txRepo.EXPECT().
		Delete(gomock.Any(), testId).
		Return(&model.Transaction{
			Id:       testId,
			IsoLevel: testIsoLevel,
		}, nil)

	td.fRepo.EXPECT().
		HardDelete(gomock.Any(), testId, nil).
		Return(contentIds, nil)

	td.cleaner.EXPECT().
		Clean(contentIds).
		Return(nil)

	uc := td.newUseCase()

	err := uc.Rollback(testCtx)

	require.NoError(t, err)
}

func TestUseCase_Rollback_Error(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "tx repo delete",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "file repo hard delete",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(&model.Transaction{}, nil)

				td.fRepo.EXPECT().
					HardDelete(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "clean",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(&model.Transaction{}, nil)

				td.fRepo.EXPECT().
					HardDelete(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, nil)

				td.cleaner.EXPECT().
					Clean(gomock.Any()).
					Return(assert.AnError)

				return assert.AnError
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			wantErr := tc.prepare(td)

			uc := td.newUseCase()

			err := uc.Rollback(testCtx)

			require.ErrorIs(t, err, wantErr)
		})
	}
}
