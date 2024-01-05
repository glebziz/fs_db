package transaction

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func TestUseCase_Commit_Success(t *testing.T) {
	for _, tc := range []struct {
		name    string
		prepare prepareFunc
	}{
		{
			name: "read committed iso level",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), testId).
					Return(&model.Transaction{
						Id:       testId,
						IsoLevel: fs_db.IsoLevelReadCommitted,
					}, nil)

				td.fRepo.EXPECT().
					UpdateTx(gomock.Any(), testId, model.MainTxId, nil).
					Return(nil)

				return nil
			},
		},
		{
			name: "serializable iso level",
			prepare: func(td *testDeps) error {
				var (
					now = time.Now().UTC()
				)

				td.txRepo.EXPECT().
					Delete(gomock.Any(), testId).
					Return(&model.Transaction{
						Id:       testId,
						IsoLevel: fs_db.IsoLevelSerializable,
						CreateTs: now,
					}, nil)

				td.fRepo.EXPECT().
					UpdateTx(gomock.Any(), testId, model.MainTxId, &model.FileFilter{
						BeforeTs: ptr.Ptr(now),
					}).
					Return(nil)

				td.fRepo.EXPECT().
					HardDelete(gomock.Any(), testId, nil).
					Return(nil, fs_db.NotFoundErr)

				return nil
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			_ = tc.prepare(td)

			uc := td.newUseCase()

			err := uc.Commit(testCtx)

			require.NoError(t, err)
		})
	}
}

func TestUseCase_Commit_Error(t *testing.T) {
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
			name: "file repo update tx",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(&model.Transaction{
						IsoLevel: fs_db.IsoLevelSerializable,
					}, nil)

				td.fRepo.EXPECT().
					UpdateTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "file repo hard delete",
			prepare: func(td *testDeps) error {
				td.txRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(&model.Transaction{
						IsoLevel: fs_db.IsoLevelSerializable,
					}, nil)

				td.fRepo.EXPECT().
					UpdateTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				td.fRepo.EXPECT().
					HardDelete(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)

				return assert.AnError
			},
		},
		{
			name: "serialization",
			prepare: func(td *testDeps) error {
				var (
					contentIds = []string{gofakeit.UUID()}
				)
				td.txRepo.EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					Return(&model.Transaction{
						IsoLevel: fs_db.IsoLevelSerializable,
					}, nil)

				td.fRepo.EXPECT().
					UpdateTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				td.fRepo.EXPECT().
					HardDelete(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(contentIds, nil)

				td.cleaner.EXPECT().
					Send(contentIds)

				return fs_db.TxSerializationErr
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)

			wantErr := tc.prepare(td)

			uc := td.newUseCase()

			err := uc.Commit(testCtx)

			require.ErrorIs(t, err, wantErr)
		})
	}
}
