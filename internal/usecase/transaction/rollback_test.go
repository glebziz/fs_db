package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func TestUseCase_Rollback(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), testId).
					Times(1).
					Return(model.Transaction{
						Id:       testId,
						IsoLevel: testIsoLevel,
						Seq:      sequence.Next(),
					}, nil)

				td.fRepo.EXPECT().
					DeleteTx(gomock.Any(), testId).
					Times(1).
					Return([]model.File{{}})

				td.cleaner.EXPECT().
					DeleteFilesAsync(gomock.Any(), []model.File{{}}).
					Times(1)
			},
		},
		{
			name: "success with empty deleteFiles",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), testId).
					Times(1).
					Return(model.Transaction{
						Id:       testId,
						IsoLevel: testIsoLevel,
						Seq:      sequence.Next(),
					}, nil)

				td.fRepo.EXPECT().
					DeleteTx(gomock.Any(), testId).
					Times(1).
					Return(nil)
			},
		},
		{
			name: "tx repo delete not found",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Times(1).
					Return(model.Transaction{}, fs_db.ErrTxNotFound)
			},
		},
		{
			name: "tx repo delete error",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Times(1).
					Return(model.Transaction{}, assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			uc := td.newUseCase()

			tc.prepare(td)

			err := uc.Rollback(testCtx)

			require.ErrorIs(t, err, tc.err)
		})
	}
}
