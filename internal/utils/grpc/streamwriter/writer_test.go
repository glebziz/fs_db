package streamwriter_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/grpc/streamwriter"
	mock_streamwriter "github.com/glebziz/fs_db/internal/utils/grpc/streamwriter/mocks"
)

func TestWriter(t *testing.T) {
	t.Parallel()

	for bufLen := 1; bufLen < 50; bufLen += 5 {
		t.Run(fmt.Sprintf("bufLen %d", bufLen), func(t *testing.T) {
			t.Parallel()

			var (
				err error
				n   int

				data   []byte
				chunks = [5][]byte{
					[]byte("hello"),
					[]byte("world"),
					[]byte("i am"),
					[]byte("gleb zhizhchenko"),
					[]byte("how are you?"),
				}

				allData     = bytes.Join(chunks[:], []byte(""))
				sendCounter = len(allData) / bufLen
			)

			ctrl := gomock.NewController(t)
			stream := mock_streamwriter.NewMockStreamTest(ctrl)

			stream.EXPECT().
				Send(gomock.Any()).
				DoAndReturn(func(r *streamwriter.TestRequest) error {
					r.ProtoReflect()

					require.Equal(t, model.Seek{}, r.Seek)

					data = append(data, r.P...)
					if sendCounter > 0 {
						require.Len(t, r.P, bufLen)
					} else {
						require.Len(t, r.P, len(allData)%bufLen)
					}
					sendCounter--
					return nil
				})

			stream.EXPECT().
				Send(gomock.Any()).
				DoAndReturn(func(r *streamwriter.TestRequest) error {
					require.Equal(t, model.Seek{
						Dir: model.SeekCurrent,
					}, r.Seek)

					data = append(data, r.P...)
					if sendCounter > 0 {
						require.Len(t, r.P, bufLen)
					} else {
						require.Len(t, r.P, len(allData)%bufLen)
					}
					sendCounter--
					return nil
				}).
				AnyTimes()

			stream.EXPECT().
				CloseAndRecv().
				Return(nil, nil)

			w := streamwriter.New(bufLen, stream, func(p []byte, seek model.Seek) *streamwriter.TestRequest {
				return &streamwriter.TestRequest{
					P:    p,
					Seek: seek,
				}
			})

			for _, chunk := range chunks {
				n, err = w.Write(chunk)

				require.NoError(t, err)
				require.Len(t, chunk, n)
			}

			pos, err := w.Seek(0, io.SeekCurrent)
			require.NoError(t, err)
			require.EqualValues(t, len(allData), pos)

			err = w.Close()

			require.NoError(t, err)
			require.Equal(t, allData, data)
		})
	}
}

func TestWriter_Write(t *testing.T) {
	t.Parallel()

	t.Run("send error", func(t *testing.T) {
		t.Parallel()

		const (
			chunkSize = 10
		)

		ctrl := gomock.NewController(t)
		stream := mock_streamwriter.NewMockStreamTest(ctrl)

		stream.EXPECT().
			Send(gomock.Any()).
			Return(assert.AnError)

		w := streamwriter.New(chunkSize, stream, func(p []byte, seek model.Seek) *streamwriter.TestRequest {
			return &streamwriter.TestRequest{
				P:    p,
				Seek: seek,
			}
		})

		n, err := w.Write(bytes.Repeat([]byte("12"), chunkSize))

		require.ErrorIs(t, err, assert.AnError)
		require.Zero(t, n)
	})
}

func TestWriter_Close(t *testing.T) {
	t.Parallel()

	t.Run("send error", func(t *testing.T) {
		t.Parallel()

		const (
			chunkSize = 10
		)

		ctrl := gomock.NewController(t)
		stream := mock_streamwriter.NewMockStreamTest(ctrl)

		stream.EXPECT().
			Send(gomock.Any()).
			Return(assert.AnError)

		w := streamwriter.New(chunkSize, stream, func(p []byte, seek model.Seek) *streamwriter.TestRequest {
			return &streamwriter.TestRequest{
				P:    p,
				Seek: seek,
			}
		})

		n, err := w.Write([]byte("12"))
		require.NoError(t, err)
		require.Equal(t, 2, n)

		err = w.Close()

		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("send error", func(t *testing.T) {
		t.Parallel()

		const (
			chunkSize = 10
		)

		ctrl := gomock.NewController(t)
		stream := mock_streamwriter.NewMockStreamTest(ctrl)

		stream.EXPECT().
			CloseAndRecv().
			Return(nil, assert.AnError)

		w := streamwriter.New(chunkSize, stream, func(p []byte, seek model.Seek) *streamwriter.TestRequest {
			return &streamwriter.TestRequest{
				P:    p,
				Seek: seek,
			}
		})

		err := w.Close()
		require.ErrorIs(t, err, assert.AnError)
	})
}

func TestWriter_Seek(t *testing.T) {
	t.Parallel()

	const (
		chunkSize = 1 << (iota + 10)
		offset

		content = "testContent"
	)

	for _, tc := range []struct {
		name    string
		offset  int64
		whence  int
		prepare func(t *testing.T, w io.Writer, stream *mock_streamwriter.MockStreamTest)
		pos     int64
		err     error
	}{
		{
			name:    "success with buffered data",
			offset:  offset,
			whence:  io.SeekStart,
			prepare: func(_ *testing.T, _ io.Writer, _ *mock_streamwriter.MockStreamTest) {},
			pos:     offset,
		},
		{
			name:   "success with buffered data",
			offset: offset,
			whence: io.SeekCurrent,
			prepare: func(t *testing.T, w io.Writer, stream *mock_streamwriter.MockStreamTest) {
				n, err := w.Write([]byte(content))
				require.NoError(t, err)
				require.Equal(t, len(content), n)

				stream.EXPECT().
					Send(&streamwriter.TestRequest{
						P:    []byte(content),
						Seek: model.Seek{},
					}).
					Return(nil)
			},
			pos: offset + int64(len(content)),
		},
		{
			name:   "flush error",
			offset: offset,
			whence: io.SeekStart,
			prepare: func(t *testing.T, w io.Writer, stream *mock_streamwriter.MockStreamTest) {
				n, err := w.Write([]byte(content))
				require.NoError(t, err)
				require.Equal(t, len(content), n)

				stream.EXPECT().
					Send(gomock.Any()).
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name:   "seek error",
			offset: -int64(len(content) + 1),
			whence: io.SeekCurrent,
			prepare: func(t *testing.T, w io.Writer, stream *mock_streamwriter.MockStreamTest) {
				n, err := w.Write([]byte(content))
				require.NoError(t, err)
				require.Equal(t, len(content), n)

				stream.EXPECT().
					Send(gomock.Any()).
					Return(nil)
			},
			err: model.ErrInvalidPosition,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			stream := mock_streamwriter.NewMockStreamTest(ctrl)

			w := streamwriter.New(chunkSize, stream, func(p []byte, seek model.Seek) *streamwriter.TestRequest {
				return &streamwriter.TestRequest{
					P:    p,
					Seek: seek,
				}
			})

			tc.prepare(t, w, stream)

			pos, err := w.Seek(tc.offset, tc.whence)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.pos, pos)
		})
	}
}

func TestWriter_WriteAt(t *testing.T) {
	t.Parallel()

	const (
		pos          = 1 << 10
		content      = "testContent"
		shortContent = "content"
	)

	for _, tc := range []struct {
		name    string
		pos     int64
		prepare func(t *testing.T, w io.Writer, stream *mock_streamwriter.MockStreamTest)
		n       int
		err     error
	}{
		{
			name: "success",
			pos:  pos,
			prepare: func(_ *testing.T, _ io.Writer, stream *mock_streamwriter.MockStreamTest) {
				stream.EXPECT().
					Send(&streamwriter.TestRequest{
						P: []byte(content),
						Seek: model.Seek{
							Pos: pos,
							Dir: model.SeekStart,
						},
					}).
					Return(nil)
			},
			n: len(content),
		},
		{
			name: "success with buffered data",
			pos:  pos,
			prepare: func(t *testing.T, w io.Writer, stream *mock_streamwriter.MockStreamTest) {
				n, err := w.Write([]byte(shortContent))
				require.NoError(t, err)
				require.Equal(t, len(shortContent), n)

				stream.EXPECT().
					Send(&streamwriter.TestRequest{
						P:    []byte(shortContent),
						Seek: model.Seek{},
					}).
					Return(nil)

				stream.EXPECT().
					Send(&streamwriter.TestRequest{
						P: []byte(content),
						Seek: model.Seek{
							Pos: pos,
							Dir: model.SeekStart,
						},
					}).
					Return(nil)
			},
			n: len(content),
		},
		{
			name:    "seek error",
			pos:     -pos,
			prepare: func(*testing.T, io.Writer, *mock_streamwriter.MockStreamTest) {},
			err:     model.ErrInvalidPosition,
		},
		{
			name: "write error",
			pos:  pos,
			prepare: func(_ *testing.T, _ io.Writer, stream *mock_streamwriter.MockStreamTest) {
				stream.EXPECT().
					Send(gomock.Any()).
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			stream := mock_streamwriter.NewMockStreamTest(ctrl)

			w := streamwriter.New(len(content), stream, func(p []byte, seek model.Seek) *streamwriter.TestRequest {
				return &streamwriter.TestRequest{
					P:    p,
					Seek: seek,
				}
			})

			tc.prepare(t, w, stream)
			n, err := w.WriteAt([]byte(content), tc.pos)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.n, n)
		})
	}
}
