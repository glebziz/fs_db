package store

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	store "github.com/glebziz/fs_db/internal/proto"
)

func TestService_GetFileV2(t *testing.T) {
	t.Parallel()

	const (
		zero int64 = iota

		key     = "key"
		content = "content"
	)

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		errCode codes.Code
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Context().
					Return(context.Background())

				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Header{
							Header: &store.FileHeader{
								Key: key,
							},
						},
					}, nil)

				td.suc.EXPECT().
					Get(gomock.Any(), key).
					Return(td.r, nil)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.r.EXPECT().
					Seek(zero, io.SeekEnd).
					Return(int64(len(content)), nil)

				td.r.EXPECT().
					Seek(zero, io.SeekStart).
					Return(zero, nil)

				td.g2Stream.EXPECT().
					Send(&store.GetFileV2Response{
						Data: &store.GetFileV2Response_Size{
							Size: int64(len(content)),
						},
					}).
					Return(nil)

				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &store.Seek{
								Offset: zero,
								Dir:    store.Seek_SeekCurrent,
							},
						},
					}, nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					DoAndReturn(func(p []byte) (int, error) {
						copy(p, content)

						return len(content), nil
					})

				td.g2Stream.EXPECT().
					Send(&store.GetFileV2Response{
						Data: &store.GetFileV2Response_Chunk{
							Chunk: []byte(content),
						},
					}).
					Return(nil)

				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &store.Seek{
								Offset: zero,
								Dir:    store.Seek_SeekStart,
							},
						},
					}, nil)

				td.r.EXPECT().
					Seek(zero, io.SeekStart).
					Return(zero, nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					Return(0, io.EOF)
			},
		},
		{
			name: "Recv error",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Recv().
					Return(nil, assert.AnError)
			},
			errCode: codes.Internal,
		},
		{
			name: "Get error",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Context().
					Return(context.Background())

				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Header{
							Header: &store.FileHeader{},
						},
					}, nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)
			},
			errCode: codes.Internal,
		},
		{
			name: "getSize error",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Context().
					Return(context.Background())

				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Header{
							Header: &store.FileHeader{},
						},
					}, nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, assert.AnError)
			},
			errCode: codes.Internal,
		},
		{
			name: "Send size error",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Context().
					Return(context.Background())

				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Header{
							Header: &store.FileHeader{},
						},
					}, nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(int64(len(content)), nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, nil)

				td.g2Stream.EXPECT().
					Send(gomock.Any()).
					Return(assert.AnError)
			},
			errCode: codes.Internal,
		},
		{
			name: "sendContent error",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Context().
					Return(context.Background())

				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Header{
							Header: &store.FileHeader{},
						},
					}, nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(int64(len(content)), nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, nil)

				td.g2Stream.EXPECT().
					Send(gomock.Any()).
					Return(nil)

				td.g2Stream.EXPECT().
					Recv().
					Return(nil, assert.AnError)
			},
			errCode: codes.Internal,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			s := td.newService()
			err := s.GetFileV2(td.g2Stream)
			require.Equal(t, tc.errCode, status.Code(err))
		})
	}
}

func Test_sendContent(t *testing.T) {
	t.Parallel()

	const (
		zero int64 = iota
		offset

		content = "content"
	)

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &store.Seek{
								Offset: zero,
								Dir:    store.Seek_SeekCurrent,
							},
						},
					}, nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					DoAndReturn(func(p []byte) (int, error) {
						copy(p, content)

						return len(content), nil
					})

				td.g2Stream.EXPECT().
					Send(&store.GetFileV2Response{
						Data: &store.GetFileV2Response_Chunk{
							Chunk: []byte(content),
						},
					}).
					Return(nil)

				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &store.Seek{
								Offset: zero,
								Dir:    store.Seek_SeekStart,
							},
						},
					}, nil)

				td.r.EXPECT().
					Seek(zero, io.SeekStart).
					Return(zero, nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					Return(0, io.EOF)
			},
		},
		{
			name: "success with Recv EOF",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &store.Seek{
								Offset: zero,
								Dir:    store.Seek_SeekCurrent,
							},
						},
					}, nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					DoAndReturn(func(p []byte) (int, error) {
						copy(p, content)

						return len(content), nil
					})

				td.g2Stream.EXPECT().
					Send(&store.GetFileV2Response{
						Data: &store.GetFileV2Response_Chunk{
							Chunk: []byte(content),
						},
					}).
					Return(nil)

				td.g2Stream.EXPECT().
					Recv().
					Return(nil, io.EOF)
			},
		},
		{
			name: "Recv error",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Recv().
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "Seek error",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &store.Seek{
								Offset: offset,
								Dir:    store.Seek_SeekCurrent,
							},
						},
					}, nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "Read error",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &store.Seek{},
						},
					}, nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					Return(0, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "Send error",
			prepare: func(td *testDeps) {
				td.g2Stream.EXPECT().
					Recv().
					Return(&store.GetFileV2Request{
						Data: &store.GetFileV2Request_Pos{
							Pos: &store.Seek{},
						},
					}, nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					Return(1, nil)

				td.g2Stream.EXPECT().
					Send(gomock.Any()).
					Return(assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			err := sendContent(td.g2Stream, 0, td.r)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func Test_getSize(t *testing.T) {
	t.Parallel()

	const (
		zero int64 = iota
		size
	)

	for _, tc := range []struct {
		name    string
		prepare prepareFunc
		size    int64
		err     error
	}{
		{
			name: "success",
			prepare: func(td *testDeps) {
				td.r.EXPECT().
					Seek(zero, io.SeekEnd).
					Return(size, nil)

				td.r.EXPECT().
					Seek(zero, io.SeekStart).
					Return(zero, nil)
			},
			size: size,
		},
		{
			name: "Seek end error",
			prepare: func(td *testDeps) {
				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "Seek start error",
			prepare: func(td *testDeps) {
				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(size, nil)

				td.r.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(zero, assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			s, err := getSize(td.r)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.size, s)
		})
	}
}
