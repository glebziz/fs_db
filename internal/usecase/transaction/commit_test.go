package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestUseCase_Commit(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		err     error
	}{
		{
			name: "read committed iso level",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), testId).
					Times(1).
					Return(model.Transaction{
						Id:       testId,
						IsoLevel: fs_db.IsoLevelReadCommitted,
						Seq:      sequence.Next(),
					}, nil)

				td.fRepo.EXPECT().
					UpdateTx(gomock.Any(), testId, model.MainTxId, model.FileFilter{}).
					Times(1).
					Return([]model.File{{}}, nil)

				td.cleaner.EXPECT().
					DeleteFilesAsync(gomock.Any(), []model.File{{}}).
					Times(1)
			},
		},
		{
			name: "serializable iso level",
			prepare: func(td *testDeps) {
				seq := sequence.Next()

				td.txRepo.EXPECT().
					Delete(gomock.Any(), testId).
					Times(1).
					Return(model.Transaction{
						Id:       testId,
						IsoLevel: fs_db.IsoLevelSerializable,
						Seq:      seq,
					}, nil)

				td.fRepo.EXPECT().
					UpdateTx(gomock.Any(), testId, model.MainTxId, model.FileFilter{
						BeforeSeq: ptr.Ptr(seq),
					}).
					Times(1).
					Return([]model.File{}, nil)
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
		{
			name: "file repo update tx error",
			prepare: func(td *testDeps) {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Times(1).
					Return(model.Transaction{}, nil)

				td.fRepo.EXPECT().
					UpdateTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return([]model.File{{}}, assert.AnError)

				td.cleaner.EXPECT().
					DeleteFilesAsync(gomock.Any(), []model.File{{}}).
					Times(1)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			uc := td.newUseCase()

			tc.prepare(td)

			err := uc.Commit(testCtx)

			require.ErrorIs(t, err, tc.err)
		})
	}
}
