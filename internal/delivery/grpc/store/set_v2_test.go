package store

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

func TestService_SetFileV2(t *testing.T) {
	t.Parallel()

	const (
		offset int64 = 10

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
				td.s2Stream.EXPECT().
					Recv().
					Return(&store.SetFileV2Request{
						Data: &store.SetFileV2Request_Header{
							Header: &store.FileHeader{
								Key: key,
							},
						},
					}, nil)

				td.s2Stream.EXPECT().
					Context().
					Return(context.Background())

				td.suc.EXPECT().
					Set(gomock.Any(), key, gomock.Any()).
					DoAndReturn(func(_ context.Context, _ string, contents model.Contents) error {
						var (
							counter int
							data    bytes.Buffer
						)
						for c := range contents {
							counter++

							_, err := io.Copy(&data, c.Reader)
							require.NoError(td.t, err)
						}

						require.Equal(td.t, content, data.String())
						require.Equal(td.t, 1, counter)

						return nil
					})

				td.s2Stream.EXPECT().
					Recv().
					Return(&store.SetFileV2Request{
						Data: &store.SetFileV2Request_Chunk{
							Chunk: &store.DataChunk{
								Chunk: []byte(content),
								Pos: &store.Seek{
									Offset: offset,
									Dir:    store.Seek_SeekCurrent,
								},
							},
						},
					}, nil)

				td.s2Stream.EXPECT().
					Recv().
					Return(nil, io.EOF)

				td.s2Stream.EXPECT().
					SendAndClose(&store.SetFileV2Response{}).
					Return(nil)
			},
		},
		{
			name: "Recv error",
			prepare: func(td *testDeps) {
				td.s2Stream.EXPECT().
					Recv().
					Return(nil, assert.AnError)
			},
			errCode: codes.Internal,
		},
		{
			name: "ErrHeaderNotFound error",
			prepare: func(td *testDeps) {
				td.s2Stream.EXPECT().
					Recv().
					Return(&store.SetFileV2Request{}, nil)
			},
			errCode: codes.Internal,
		},
		{
			name: "Set error",
			prepare: func(td *testDeps) {
				td.s2Stream.EXPECT().
					Recv().
					Return(&store.SetFileV2Request{
						Data: &store.SetFileV2Request_Header{
							Header: &store.FileHeader{},
						},
					}, nil)

				td.s2Stream.EXPECT().
					Context().
					Return(context.Background())

				td.suc.EXPECT().
					Set(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(assert.AnError)

				td.s2Stream.EXPECT().
					Recv().
					Return(nil, io.EOF)
			},
			errCode: codes.Internal,
		},
		{
			name: "getChunk error",
			prepare: func(td *testDeps) {
				td.s2Stream.EXPECT().
					Recv().
					Return(&store.SetFileV2Request{
						Data: &store.SetFileV2Request_Header{
							Header: &store.FileHeader{},
						},
					}, nil)

				td.s2Stream.EXPECT().
					Context().
					Return(context.Background())

				td.suc.EXPECT().
					Set(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				td.s2Stream.EXPECT().
					Recv().
					Return(nil, assert.AnError)
			},
			errCode: codes.Internal,
		},
		{
			name: "SendAndClose error",
			prepare: func(td *testDeps) {
				td.s2Stream.EXPECT().
					Recv().
					Return(&store.SetFileV2Request{
						Data: &store.SetFileV2Request_Header{
							Header: &store.FileHeader{},
						},
					}, nil)

				td.s2Stream.EXPECT().
					Context().
					Return(context.Background())

				td.suc.EXPECT().
					Set(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				td.s2Stream.EXPECT().
					Recv().
					Return(nil, io.EOF)

				td.s2Stream.EXPECT().
					SendAndClose(gomock.Any()).
					Return(assert.AnError)
			},
			errCode: codes.Internal,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			s := td.newService()
			err := s.SetFileV2(td.s2Stream)

			require.Equal(t, tc.errCode, status.Code(err))
		})
	}
}

func Test_getChunk(t *testing.T) {
	t.Parallel()

	const (
		offset int64 = 10

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
				td.s2Stream.EXPECT().
					Recv().
					Return(&store.SetFileV2Request{
						Data: &store.SetFileV2Request_Chunk{
							Chunk: &store.DataChunk{
								Chunk: []byte(content),
								Pos: &store.Seek{
									Offset: offset,
									Dir:    store.Seek_SeekCurrent,
								},
							},
						},
					}, nil)

				td.w.EXPECT().
					Seek(offset, io.SeekCurrent).
					Return(offset, nil)

				td.w.EXPECT().
					Write([]byte(content)).
					Return(len(content), nil)

				td.s2Stream.EXPECT().
					Recv().
					Return(nil, io.EOF)
			},
		},
		{
			name: "Recv error",
			prepare: func(td *testDeps) {
				td.s2Stream.EXPECT().
					Recv().
					Return(nil, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "Seek error",
			prepare: func(td *testDeps) {
				td.s2Stream.EXPECT().
					Recv().
					Return(&store.SetFileV2Request{
						Data: &store.SetFileV2Request_Chunk{
							Chunk: &store.DataChunk{},
						},
					}, nil)

				td.w.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(offset, assert.AnError)
			},
			err: assert.AnError,
		},
		{
			name: "Write error",
			prepare: func(td *testDeps) {
				td.s2Stream.EXPECT().
					Recv().
					Return(&store.SetFileV2Request{
						Data: &store.SetFileV2Request_Chunk{
							Chunk: &store.DataChunk{},
						},
					}, nil)

				td.w.EXPECT().
					Seek(gomock.Any(), gomock.Any()).
					Return(offset, nil)

				td.w.EXPECT().
					Write(gomock.Any()).
					Return(0, assert.AnError)
			},
			err: assert.AnError,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			td := newTestDeps(t)
			tc.prepare(td)

			err := getChunk(td.s2Stream, td.w)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
