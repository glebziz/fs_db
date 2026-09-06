package writer

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/usecase/writer/mocks"
)

func TestUseCase_Write(t *testing.T) {
	t.Parallel()

	const (
		testOffset int64 = iota + 1

		testPath     = "testPath"
		testContent1 = "testContent1"
		testContent2 = "testContent2"
	)

	var (
		contents = []string{testContent1, testContent2}
	)

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				var buf bytes.Buffer
				td.t.Cleanup(func() {
					require.Equal(td.t, []byte(strings.Join(contents, "")), buf.Bytes())
				})

				cw := mocks.NewMockReadWriteSeekCloser(td.ctrl)
				cw.EXPECT().
					Write(gomock.Any()).
					AnyTimes().
					DoAndReturn(func(p []byte) (int, error) {
						return buf.Write(p)
					})

				cw.EXPECT().
					Seek(testOffset, io.SeekStart).
					Times(len(contents)).
					Return(testOffset, nil)

				cw.EXPECT().
					Close().
					Return(nil)

				td.cRepo.EXPECT().
					Create(gomock.Any(), testPath).
					Return(cw, nil)
			},
		},
		{
			name: "Create error",
			prepare: func(td *testDeps) {
				td.cRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "Seek.Apply error",
			prepare: func(td *testDeps) {
				cw := mocks.NewMockReadWriteSeekCloser(td.ctrl)
				cw.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(0, assert.AnError)

				cw.EXPECT().
					Close().
					Return(nil)

				td.cRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(cw, nil)
			},
			err: assert.AnError,
		},
		{
			name: "Copy error",
			prepare: func(td *testDeps) {
				cw := mocks.NewMockReadWriteSeekCloser(td.ctrl)
				cw.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(testOffset, nil)

				cw.EXPECT().
					Write(gomock.Any()).
					Return(0, assert.AnError)

				cw.EXPECT().
					Close().
					Return(nil)

				td.cRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(cw, nil)
			},
			err: assert.AnError,
		},
		{
			name: "Seek error",
			prepare: func(td *testDeps) {
				cw := mocks.NewMockReadWriteSeekCloser(td.ctrl)
				cw.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(testOffset, nil)

				cw.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(0, assert.AnError)

				cw.EXPECT().
					Write(gomock.Any()).
					Return(0, model.ErrNotEnoughSpace)

				td.cRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(cw, nil)
			},
			err: assert.AnError,
		},
		{
			name: "ErrNotEnoughSpace",
			prepare: func(td *testDeps) {
				cw := mocks.NewMockReadWriteSeekCloser(td.ctrl)
				cw.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(testOffset, nil)

				cw.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(0, nil)

				cw.EXPECT().
					Write(gomock.Any()).
					Return(0, model.ErrNotEnoughSpace)

				td.cRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(cw, nil)
			},
			err: model.ErrNotEnoughSpace,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			ch := make(chan model.Content, len(contents))
			for _, content := range contents {
				ch <- model.Content{
					Seek: model.Seek{
						Pos: testOffset,
						Dir: model.SeekStart,
					},
					Reader: strings.NewReader(content),
				}
			}
			close(ch)

			u := td.newUseCase()
			err := u.Write(context.Background(), testPath, ch)

			require.ErrorIs(t, err, tc.err)
		})
	}
}
