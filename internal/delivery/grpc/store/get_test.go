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

func TestService_GetFile(t *testing.T) {
	t.Parallel()

	const (
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
				td.gStream.EXPECT().
					Context().
					Return(context.Background())

				td.suc.EXPECT().
					Get(gomock.Any(), key).
					Return(td.r, nil)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.gStream.EXPECT().
					Send(&store.GetFileResponse{
						Data: &store.GetFileResponse_Header{
							Header: &store.FileHeader{
								Key: key,
							},
						},
					}).
					Return(nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					DoAndReturn(func(p []byte) (int, error) {
						copy(p, content)

						return len(content), nil
					})

				td.gStream.EXPECT().
					Send(&store.GetFileResponse{
						Data: &store.GetFileResponse_Chunk{
							Chunk: []byte(content),
						},
					}).
					Return(nil)

				td.r.EXPECT().
					Read(gomock.Any()).
					Return(0, io.EOF)
			},
		},
		{
			name: "Get error",
			prepare: func(td *testDeps) {
				td.gStream.EXPECT().
					Context().
					Return(context.Background())

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(nil, assert.AnError)
			},
			errCode: codes.Internal,
		},
		{
			name: "Send header error",
			prepare: func(td *testDeps) {
				td.gStream.EXPECT().
					Context().
					Return(context.Background())

				td.r.EXPECT().
					Close().
					Return(nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)

				td.gStream.EXPECT().
					Send(gomock.Any()).
					Return(assert.AnError)
			},
			errCode: codes.Internal,
		},
		{
			name: "Read error",
			prepare: func(td *testDeps) {
				td.gStream.EXPECT().
					Context().
					Return(context.Background())

				td.r.EXPECT().
					Read(gomock.Any()).
					Return(0, assert.AnError)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)

				td.gStream.EXPECT().
					Send(gomock.Any()).
					Return(nil)
			},
			errCode: codes.Internal,
		},
		{
			name: "Send chunk error",
			prepare: func(td *testDeps) {
				td.gStream.EXPECT().
					Context().
					Return(context.Background())

				td.r.EXPECT().
					Read(gomock.Any()).
					Return(1, nil)

				td.r.EXPECT().
					Close().
					Return(nil)

				td.suc.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Return(td.r, nil)

				td.gStream.EXPECT().
					Send(gomock.Any()).
					Return(nil)

				td.gStream.EXPECT().
					Send(gomock.Any()).
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
			err := s.GetFile(&store.GetFileRequest{
				Key: key,
			}, td.gStream)

			require.Equal(t, tc.errCode, status.Code(err))
		})
	}
}
