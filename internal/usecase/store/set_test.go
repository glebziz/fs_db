package store

import (
	"bytes"
	"context"
	"io"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func TestUseCase_Set(t *testing.T) {
	t.Parallel()

	const (
		freeSpace uint64 = iota + 1
		key              = "key"
		content          = "content"
		dir              = "dir"
		root             = "root"
	)

	for _, tc := range []struct {
		name    string
		key     string
		prepare prepareFunc
		err     error
	}{
		{
			name: "success",
			key:  key,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Get(gomock.Any()).
					Return(model.Dirs{{
						Name: dir,
						Root: root,
						Free: freeSpace,
					}}, nil)

				td.cWriter.EXPECT().
					Write(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				td.cfRepo.EXPECT().
					Store(gomock.Any(), model.ContentFile{
						Id:     testContentId,
						Parent: path.Join(root, dir),
					}).
					Return(nil)

				td.fRepo.EXPECT().
					Store(gomock.Any(), model.File{
						Key:       key,
						TxId:      model.MainTxId,
						ContentId: testContentId,
					}).
					Return(nil)

				return nil
			},
		},
		{
			name: "ErrEmptyKey error",
			key:  "",
			prepare: func(td *testDeps) error {
				return nil
			},
			err: fs_db.ErrEmptyKey,
		},
		{
			name: "dir.Get error",
			key:  key,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Get(gomock.Any()).
					Return(nil, assert.AnError)

				return nil
			},
			err: assert.AnError,
		},
		{
			name: "writeContent error",
			key:  key,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Get(gomock.Any()).
					Return(model.Dirs{{
						Free: freeSpace,
					}}, nil)

				td.cWriter.EXPECT().
					Write(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return nil
			},
			err: assert.AnError,
		},
		{
			name: "cfRepo.Store error",
			key:  key,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Get(gomock.Any()).
					Return(model.Dirs{{
						Free: freeSpace,
					}}, nil)

				td.cWriter.EXPECT().
					Write(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				td.cfRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return nil
			},
			err: assert.AnError,
		},
		{
			name: "fRepo.Store error",
			key:  key,
			prepare: func(td *testDeps) error {
				td.dir.EXPECT().
					Get(gomock.Any()).
					Return(model.Dirs{{
						Free: freeSpace,
					}}, nil)

				td.cWriter.EXPECT().
					Write(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				td.cfRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(nil)

				td.fRepo.EXPECT().
					Store(gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return nil
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			u := td.newUseCase()
			err := u.Set(context.Background(), tc.key, func() model.Contents {
				ch := make(chan model.Content, 1)
				ch <- model.Content{
					Reader: strings.NewReader(content),
				}
				close(ch)

				return ch
			}())

			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestUseCase_writeContent(t *testing.T) {
	t.Parallel()

	const (
		freeSpace1 uint64 = iota + 1
		freeSpace2
		freeSpace3

		dir1      = "dir1"
		dir2      = "dir2"
		dir3      = "dir3"
		root      = "root"
		contentId = "contentId"
		content   = "content"
	)

	checkContents := func(t *testing.T, contents model.Contents) {
		var data bytes.Buffer
		for c := range contents {
			buf, err := io.ReadAll(c.Reader)
			require.NoError(t, err)

			data.Write(buf)
		}

		require.EqualValues(t, content, data.Bytes())
	}

	for _, tc := range []struct {
		name    string
		dirs    model.Dirs
		prepare prepareFunc
		err     error
	}{
		{
			name: "success",
			dirs: model.Dirs{{
				Name: dir1,
				Root: root,
				Free: freeSpace1,
			}},
			prepare: func(td *testDeps) error {
				td.cWriter.EXPECT().
					Write(gomock.Any(), path.Join(root, dir1, contentId), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, contents model.Contents) error {
						checkContents(t, contents)

						return nil
					})

				return nil
			},
		},
		{
			name: "success with not enough space in first dir",
			dirs: model.Dirs{{
				Name: dir1,
				Root: root,
				Free: freeSpace3,
			}, {
				Name: dir2,
				Root: root,
				Free: freeSpace2,
			}, {
				Name: dir3,
				Root: root,
				Free: freeSpace1,
			}},
			prepare: func(td *testDeps) error {
				td.cWriter.EXPECT().
					Write(gomock.Any(), path.Join(root, dir2, contentId), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, contents model.Contents) error {
						checkContents(t, contents)

						return model.NotEnoughSpaceError{
							Err:    model.ErrNotEnoughSpace,
							Start:  io.NopCloser(strings.NewReader(content)),
							Middle: strings.NewReader(""),
							End:    strings.NewReader(""),
						}
					})

				td.cWriter.EXPECT().
					Write(gomock.Any(), path.Join(root, dir1, contentId), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, contents model.Contents) error {
						checkContents(t, contents)
						return nil
					})

				return nil
			},
		},
		{
			name: "success with not enough space in first second dirs",
			dirs: model.Dirs{{
				Name: dir1,
				Root: root,
				Free: freeSpace3,
			}, {
				Name: dir2,
				Root: root,
				Free: freeSpace1,
			}, {
				Name: dir3,
				Root: root,
				Free: freeSpace2,
			}},
			prepare: func(td *testDeps) error {
				td.cWriter.EXPECT().
					Write(gomock.Any(), path.Join(root, dir2, contentId), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, contents model.Contents) error {
						checkContents(t, contents)

						return model.NotEnoughSpaceError{
							Err:    model.ErrNotEnoughSpace,
							Start:  io.NopCloser(strings.NewReader(content)),
							Middle: strings.NewReader(""),
							End:    strings.NewReader(""),
						}
					})

				td.cWriter.EXPECT().
					Write(gomock.Any(), path.Join(root, dir3, contentId), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, contents model.Contents) error {
						checkContents(t, contents)

						return model.NotEnoughSpaceError{
							Err:    model.ErrNotEnoughSpace,
							Start:  io.NopCloser(strings.NewReader(content)),
							Middle: strings.NewReader(""),
							End:    strings.NewReader(""),
						}
					})

				td.cWriter.EXPECT().
					Write(gomock.Any(), path.Join(root, dir1, contentId), gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, contents model.Contents) error {
						checkContents(t, contents)
						return nil
					})

				return nil
			},
		},
		{
			name: "ErrNoFreeSpace error",
			dirs: model.Dirs{{
				Name: dir1,
				Root: root,
				Free: freeSpace1,
			}},
			prepare: func(td *testDeps) error {
				td.cWriter.EXPECT().
					Write(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(model.NotEnoughSpaceError{
						Err:    model.ErrNotEnoughSpace,
						Start:  io.NopCloser(strings.NewReader(content)),
						Middle: strings.NewReader(""),
						End:    strings.NewReader(""),
					})

				return nil
			},
			err: fs_db.ErrNoFreeSpace,
		},
		{
			name: "Write error",
			dirs: model.Dirs{{
				Name: dir1,
				Root: root,
				Free: freeSpace1,
			}},
			prepare: func(td *testDeps) error {
				td.cWriter.EXPECT().
					Write(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				return nil
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			u := td.newUseCase()
			err := u.writeContent(context.Background(), tc.dirs, &model.ContentFile{
				Id: contentId,
			}, func() model.Contents {
				ch := make(chan model.Content, 1)
				ch <- model.Content{
					Seek:   model.Seek{},
					Reader: strings.NewReader(content),
				}
				close(ch)

				return ch
			}())

			require.ErrorIs(t, err, tc.err)
		})
	}
}
