package cleaner

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/usecase/cleaner/mocks"
)

func TestCleaner(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		var (
			contentIds   = []string{gofakeit.UUID(), gofakeit.UUID()}
			contentFiles = []model.ContentFile{{
				Id:         contentIds[0],
				ParentPath: gofakeit.UUID(),
			}, {
				Id:         contentIds[1],
				ParentPath: gofakeit.UUID(),
			}}
			n = 100
		)

		ctrl := gomock.NewController(t)
		cRepo := mock_cleaner.NewMockcontentRepository(ctrl)
		cfRepo := mock_cleaner.NewMockcontentFileRepository(ctrl)

		cfRepo.EXPECT().
			Delete(gomock.Any(), contentIds).
			Times(n).
			Return(contentFiles, nil)

		cRepo.EXPECT().
			Delete(gomock.Any(), contentFiles[0].GetPath()).
			Times(n).
			Return(nil)

		cRepo.EXPECT().
			Delete(gomock.Any(), contentFiles[1].GetPath()).
			Times(n).
			Return(nil)

		cc := New(cRepo, cfRepo)

		for i := 0; i < n/10; i++ {
			cc.Run()
		}

		for i := 0; i < n; i++ {
			cc.Send(contentIds)
		}

		cc.Stop()
		cc.Send(nil)
	})

	t.Run("errors", func(t *testing.T) {
		t.Parallel()

		var (
			contentIds1  = []string{gofakeit.UUID()}
			contentIds2  = []string{gofakeit.UUID()}
			contentFiles = []model.ContentFile{{
				Id:         contentIds2[0],
				ParentPath: gofakeit.UUID(),
			}}
		)

		ctrl := gomock.NewController(t)
		cRepo := mock_cleaner.NewMockcontentRepository(ctrl)
		cfRepo := mock_cleaner.NewMockcontentFileRepository(ctrl)

		cfRepo.EXPECT().
			Delete(gomock.Any(), contentIds1).
			Return(nil, assert.AnError)

		cfRepo.EXPECT().
			Delete(gomock.Any(), contentIds2).
			Return(contentFiles, nil)

		cRepo.EXPECT().
			Delete(gomock.Any(), contentFiles[0].GetPath()).
			Return(assert.AnError)

		cc := New(cRepo, cfRepo)

		cc.Run()
		cc.Send(contentIds1)
		cc.Send(contentIds2)

		cc.Stop()
	})

	t.Run("send timeout", func(t *testing.T) {
		t.Parallel()

		cc := New(nil, nil)
		cc.ch = make(chan []string)

		cc.Send(nil)
		cc.Send([]string{gofakeit.UUID()})

		cc.Stop()
	})
}
