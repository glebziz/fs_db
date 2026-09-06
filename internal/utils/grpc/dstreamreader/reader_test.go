package dstreamreader_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/grpc/dstreamreader"
	mock_dstreamreader "github.com/glebziz/fs_db/internal/utils/grpc/dstreamreader/mocks"
)

func TestReader_Read(t *testing.T) {
	t.Parallel()

	const (
		offset  int64 = 10
		content       = "hello world"
	)

	for _, tc := range []struct {
		name    string
		prepare func(t *testing.T, ctrl *gomock.Controller, stream tStream, r io.ReadSeeker)
		n       int
		err     error
	}{
		{
			name: "success",
			prepare: func(t *testing.T, ctrl *gomock.Controller, stream tStream, r io.ReadSeeker) {
				resp := mock_dstreamreader.NewMockResponse(ctrl)
				resp.EXPECT().
					GetChunk().
					Return([]byte(content)).
					Times(2)

				stream.EXPECT().
					Send(&TestRequest{}).
					Return(nil)

				stream.EXPECT().
					Recv().
					Return(resp, nil)
			},
			n: len(content),
		},
		{
			name: "success with two chunks",
			prepare: func(t *testing.T, ctrl *gomock.Controller, stream tStream, r io.ReadSeeker) {
				var (
					mid  = len(content) / 2
					resp = mock_dstreamreader.NewMockResponse(ctrl)
				)

				resp.EXPECT().
					GetChunk().
					Return([]byte(content[:mid])).
					Times(2)

				resp.EXPECT().
					GetChunk().
					Return([]byte(content[mid:])).
					Times(2)

				stream.EXPECT().
					Send(&TestRequest{}).
					Return(nil)

				stream.EXPECT().
					Send(&TestRequest{
						Seek: model.Seek{
							Dir: model.SeekCurrent,
						},
					}).
					Return(nil)

				stream.EXPECT().
					Recv().
					Return(resp, nil).
					Times(2)
			},
			n: len(content),
		},
		{
			name: "success after seek",
			prepare: func(t *testing.T, ctrl *gomock.Controller, stream tStream, r io.ReadSeeker) {
				var (
					mid  = len(content) / 2
					resp = mock_dstreamreader.NewMockResponse(ctrl)
				)

				resp.EXPECT().
					GetChunk().
					Return([]byte(content[:mid])).
					Times(2)

				resp.EXPECT().
					GetChunk().
					Return([]byte(content[mid:])).
					Times(2)

				stream.EXPECT().
					Send(&TestRequest{
						Seek: model.Seek{
							Pos: offset,
							Dir: model.SeekStart,
						},
					}).
					Return(nil)

				stream.EXPECT().
					Send(&TestRequest{
						Seek: model.Seek{
							Dir: model.SeekCurrent,
						},
					}).
					Return(nil)

				stream.EXPECT().
					Recv().
					Return(resp, nil).
					Times(2)

				_, err := r.Seek(offset, io.SeekStart)
				require.NoError(t, err)
			},
			n: len(content),
		},
		{
			name: "success with eof",
			prepare: func(t *testing.T, ctrl *gomock.Controller, stream tStream, r io.ReadSeeker) {
				var (
					mid  = len(content) / 2
					resp = mock_dstreamreader.NewMockResponse(ctrl)
				)

				resp.EXPECT().
					GetChunk().
					Return([]byte(content[:mid])).
					Times(2)

				stream.EXPECT().
					Send(&TestRequest{}).
					Return(nil)

				stream.EXPECT().
					Send(&TestRequest{
						Seek: model.Seek{
							Dir: model.SeekCurrent,
						},
					}).
					Return(nil)

				stream.EXPECT().
					Recv().
					Return(resp, nil)

				stream.EXPECT().
					Recv().
					Return(nil, io.EOF)
			},
			n: len(content) / 2,
		},
		{
			name: "Send error",
			prepare: func(t *testing.T, ctrl *gomock.Controller, stream tStream, r io.ReadSeeker) {
				stream.EXPECT().
					Send(gomock.Any()).
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "Recv error",
			prepare: func(t *testing.T, ctrl *gomock.Controller, stream tStream, r io.ReadSeeker) {
				stream.EXPECT().
					Send(gomock.Any()).
					Return(nil)

				stream.EXPECT().
					Recv().
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			stream := mock_dstreamreader.NewMockStream[*TestRequest, *mock_dstreamreader.MockResponse](ctrl)
			r := dstreamreader.New[*TestRequest, *mock_dstreamreader.MockResponse](
				stream, func(seek model.Seek) *TestRequest {
					return &TestRequest{
						Seek: seek,
					}
				}, int64(len(content)),
			)

			tc.prepare(t, ctrl, stream, r)

			p := make([]byte, len(content))
			n, err := r.Read(p)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.n, n)
			require.EqualValues(t, content[:tc.n], p[:n])
		})
	}
}

func TestReader_Close(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name    string
		prepare func(stream tStream)
		err     error
	}{
		{
			name: "success",
			prepare: func(stream tStream) {
				stream.EXPECT().
					CloseSend().
					Return(nil)
			},
		},
		{
			name: "close send error",
			prepare: func(stream tStream) {
				stream.EXPECT().
					CloseSend().
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			stream := mock_dstreamreader.NewMockStream[*TestRequest, *mock_dstreamreader.MockResponse](ctrl)
			r := dstreamreader.New[*TestRequest, *mock_dstreamreader.MockResponse](stream, nil, 0)

			tc.prepare(stream)

			err := r.Close()
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestReader_Seek(t *testing.T) {
	t.Parallel()

	const (
		size int64 = 10
	)

	for _, tc := range []struct {
		name   string
		offset int64
		whence int
		pos    int64
		err    error
	}{
		{
			name:   "success",
			offset: size,
			whence: io.SeekStart,
			pos:    size,
		},
		{
			name:   "success seek end",
			offset: -size,
			whence: io.SeekEnd,
			pos:    0,
		},
		{
			name:   "seek error",
			offset: -1,
			whence: io.SeekStart,
			err:    model.ErrInvalidPosition,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := dstreamreader.New[*TestRequest, *mock_dstreamreader.MockResponse](nil, nil, size)

			pos, err := r.Seek(tc.offset, tc.whence)
			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.pos, pos)
		})
	}
}

func TestReader_ReadAt(t *testing.T) {
	t.Parallel()

	const (
		offset  int64 = 10
		content       = "hello world"
	)

	for _, tc := range []struct {
		name    string
		offset  int64
		prepare func(t *testing.T, ctrl *gomock.Controller, stream tStream)
		n       int
		err     error
	}{
		{
			name:   "success",
			offset: offset,
			prepare: func(t *testing.T, ctrl *gomock.Controller, stream tStream) {
				resp := mock_dstreamreader.NewMockResponse(ctrl)
				resp.EXPECT().
					GetChunk().
					Return([]byte(content)).
					Times(2)

				stream.EXPECT().
					Send(&TestRequest{
						Seek: model.Seek{
							Pos: offset,
							Dir: model.SeekStart,
						},
					}).
					Return(nil)

				stream.EXPECT().
					Recv().
					Return(resp, nil)
			},
			n: len(content),
		},
		{
			name:    "seek error",
			offset:  -1,
			prepare: func(t *testing.T, ctrl *gomock.Controller, stream tStream) {},
			err:     model.ErrInvalidPosition,
		},
		{
			name: "write error",
			prepare: func(t *testing.T, ctrl *gomock.Controller, stream tStream) {
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
			stream := mock_dstreamreader.NewMockStream[*TestRequest, *mock_dstreamreader.MockResponse](ctrl)
			r := dstreamreader.New[*TestRequest, *mock_dstreamreader.MockResponse](
				stream, func(seek model.Seek) *TestRequest {
					return &TestRequest{
						Seek: seek,
					}
				}, int64(len(content)),
			)

			tc.prepare(t, ctrl, stream)

			p := make([]byte, len(content))
			n, err := r.ReadAt(p, tc.offset)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.n, n)
			require.EqualValues(t, content[:tc.n], p[:n])
		})
	}
}
