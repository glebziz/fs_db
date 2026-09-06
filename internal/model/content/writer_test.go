package content

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/async"
)

func TestContent_Close(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare func(w *content)
		err     error
	}{
		{
			name: "success with rw",
			prepare: func(w *content) {
				w.rw = async.NewReadWriter()
			},
		},
		{
			name:    "success without rw",
			prepare: func(w *content) {},
		},
		{
			name: "error",
			prepare: func(w *content) {
				w.SetError(assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w, _ := New()

			tc.prepare(w)

			err := w.Close()
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestContent_Seek(t *testing.T) {
	t.Parallel()

	const (
		offset int64 = 10

		contentData = "content"
	)

	for _, tc := range []struct {
		name           string
		offset         int64
		whence         int
		prepare        func(t *testing.T, w *content)
		contentCounter int
		pos            int64
		err            error
	}{
		{
			name:    "success",
			offset:  offset,
			whence:  io.SeekStart,
			prepare: func(t *testing.T, w *content) {},
			pos:     offset,
		},
		{
			name:   "seek after write",
			offset: offset,
			whence: io.SeekCurrent,
			prepare: func(t *testing.T, w *content) {
				n, err := w.Write([]byte(contentData))
				require.NoError(t, err)
				require.Equal(t, len(contentData), n)
			},
			contentCounter: 1,
			pos:            offset + int64(len(contentData)),
		},
		{
			name:   "seek zero current",
			offset: 0,
			whence: io.SeekCurrent,
			prepare: func(t *testing.T, w *content) {
				pos, err := w.Seek(offset, io.SeekStart)
				require.NoError(t, err)
				require.Equal(t, offset, pos)
			},
			pos: offset,
		},
		{
			name:    "err seek",
			offset:  -offset,
			whence:  io.SeekEnd,
			prepare: func(t *testing.T, w *content) {},
			err:     model.ErrInvalidPosition,
		},
		{
			name: "checkError error",
			prepare: func(t *testing.T, w *content) {
				w.SetError(assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w, contents := New()
			w.Add(1)

			var counter int
			go func() {
				defer w.Done()

				for c := range contents {
					counter++

					_, _ = io.ReadAll(c.Reader)
				}
			}()

			tc.prepare(t, w)

			pos, err := w.Seek(tc.offset, tc.whence)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.pos, pos)

			_ = w.Close()

			require.Equal(t, tc.contentCounter, counter)
		})
	}
}

func TestContent_Write(t *testing.T) {
	t.Parallel()

	const (
		offset int64 = 10

		contentData = "content"
	)

	for _, tc := range []struct {
		name          string
		content       string
		prepare       func(t *testing.T, w *content)
		checkContents func(t *testing.T, contents model.Contents)
		pos           int64
		err           error
	}{
		{
			name:    "success",
			content: contentData,
			prepare: func(t *testing.T, w *content) {},
			checkContents: func(t *testing.T, contents model.Contents) {
				var counter int
				for c := range contents {
					counter++

					require.Zero(t, c.Seek)

					data, err := io.ReadAll(c.Reader)
					require.NoError(t, err)
					require.EqualValues(t, contentData, data)
				}

				require.Equal(t, 1, counter)
			},
			pos: int64(len(contentData)),
		},
		{
			name:    "empty content",
			content: "",
			prepare: func(t *testing.T, w *content) {},
			checkContents: func(t *testing.T, contents model.Contents) {
				for range contents {
					require.False(t, true)
				}
			},
		},
		{
			name:    "write after write",
			content: contentData,
			prepare: func(t *testing.T, w *content) {
				n, err := w.Write([]byte(contentData))
				require.NoError(t, err)
				require.Equal(t, len(contentData), n)
			},
			checkContents: func(t *testing.T, contents model.Contents) {
				var counter int
				for c := range contents {
					counter++

					require.Zero(t, c.Seek)

					data, err := io.ReadAll(c.Reader)
					require.NoError(t, err)
					require.EqualValues(t, contentData+contentData, data)
				}

				require.Equal(t, 1, counter)
			},
			pos: int64(len(contentData) * 2),
		},
		{
			name:    "write after seek",
			content: contentData,
			prepare: func(t *testing.T, w *content) {
				pos, err := w.Seek(offset, io.SeekStart)
				require.NoError(t, err)
				require.Equal(t, offset, pos)

				pos, err = w.Seek(offset, io.SeekCurrent)
				require.NoError(t, err)
				require.Equal(t, offset*2, pos)
			},
			checkContents: func(t *testing.T, contents model.Contents) {
				var counter int
				for c := range contents {
					counter++

					require.Equal(t, model.Seek{
						Pos: offset * 2,
						Dir: model.SeekStart,
					}, c.Seek)

					data, err := io.ReadAll(c.Reader)
					require.NoError(t, err)
					require.EqualValues(t, contentData, data)
				}

				require.Equal(t, 1, counter)
			},
			pos: offset*2 + int64(len(contentData)),
		},
		{
			name:    "write after seek and write",
			content: contentData,
			prepare: func(t *testing.T, w *content) {
				pos, err := w.Seek(offset, io.SeekStart)
				require.NoError(t, err)
				require.Equal(t, offset, pos)

				n, err := w.Write([]byte(contentData))
				require.NoError(t, err)
				require.Equal(t, len(contentData), n)

				pos, err = w.Seek(offset, io.SeekCurrent)
				require.NoError(t, err)
				require.Equal(t, offset*2+int64(len(contentData)), pos)
			},
			checkContents: func(t *testing.T, contents model.Contents) {
				var (
					counter int

					seek = []model.Seek{{
						Pos: offset,
						Dir: model.SeekStart,
					}, {
						Pos: offset*2 + int64(len(contentData)),
						Dir: model.SeekStart,
					}}
				)
				for c := range contents {
					require.Equal(t, seek[counter], c.Seek)

					data, err := io.ReadAll(c.Reader)
					require.NoError(t, err)
					require.EqualValues(t, contentData, data)

					counter++
				}

				require.Equal(t, 2, counter)
			},
			pos: offset*2 + int64(len(contentData)*2),
		},
		{
			name:    "checkError error",
			content: contentData,
			prepare: func(_ *testing.T, w *content) {
				w.SetError(assert.AnError)
			},
			checkContents: func(t *testing.T, contents model.Contents) {
				for range contents {
					require.False(t, true)
				}
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w, contents := New()

			w.Go(func() {

				tc.checkContents(t, contents)
			})

			tc.prepare(t, w)

			n, err := w.Write([]byte(tc.content))
			if err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, len(tc.content), n)
			}

			_ = w.Close()

			require.Equal(t, tc.pos, w.pos.Pos())
		})
	}
}

func TestContent_WriteAt(t *testing.T) {
	t.Parallel()

	const (
		offset int64 = 10

		contentData = "content"
	)

	for _, tc := range []struct {
		name          string
		offset        int64
		prepare       func(w *content)
		checkContents func(t *testing.T, contents model.Contents)
		pos           int64
		err           error
	}{
		{
			name:    "success",
			offset:  offset,
			prepare: func(*content) {},
			checkContents: func(t *testing.T, contents model.Contents) {
				for c := range contents {
					data, err := io.ReadAll(c.Reader)
					require.NoError(t, err)
					require.EqualValues(t, contentData, data)
				}
			},
			pos: offset + int64(len(contentData)),
		},
		{
			name:    "Seek error",
			offset:  -1,
			prepare: func(*content) {},
			checkContents: func(t *testing.T, contents model.Contents) {
				for range contents {
					require.False(t, true)
				}
			},
			err: model.ErrInvalidPosition,
		},
		{
			name: "checkErr error",
			prepare: func(w *content) {
				w.SetError(assert.AnError)
			},
			checkContents: func(t *testing.T, contents model.Contents) {
				for range contents {
					require.False(t, true)
				}
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w, contents := New()

			w.Go(func() {

				tc.checkContents(t, contents)
			})

			tc.prepare(w)

			n, err := w.WriteAt([]byte(contentData), tc.offset)
			if err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, len(contentData), n)
			}

			_ = w.Close()

			require.Equal(t, tc.pos, w.pos.Pos())
		})
	}
}
