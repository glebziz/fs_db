package streamreader_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/glebziz/fs_db/internal/utils/grpc/streamreader"
	"github.com/glebziz/fs_db/internal/utils/grpc/streamreader/mocks"
)

func TestReader_Read(t *testing.T) {
	for bufLen := 1; bufLen < 50; bufLen += 5 {
		t.Run(fmt.Sprintf("bufLen %d", bufLen), func(t *testing.T) {
			var (
				err error
				n   int

				numStreamCall  = 0
				numRequestCall = 0

				buf    = make([]byte, bufLen)
				chunks = [5][]byte{
					[]byte("hello"),
					[]byte("world"),
					[]byte("i am"),
					[]byte("gleb zhizhchenko"),
					[]byte("how are you?"),
				}
			)

			ctrl := gomock.NewController(t)
			stream := mock_streamreader.NewMockStream(ctrl)
			request := mock_streamreader.NewMockRequest(ctrl)

			request.EXPECT().
				GetChunk().
				DoAndReturn(func() []byte {
					numRequestCall++

					if numRequestCall <= 5 {
						return chunks[numRequestCall-1]
					}

					return []byte{}
				}).
				Times(5)

			stream.EXPECT().
				Recv().
				DoAndReturn(func() (streamreader.Request, error) {
					numStreamCall++

					if numStreamCall <= 5 {
						return request, nil
					}

					return nil, io.EOF
				}).
				AnyTimes()

			r := streamreader.New(stream)

			for err == nil {
				n, err = r.Read(buf)

				if err == nil {
					require.NotZero(t, n)
					require.NotEmpty(t, buf[:n])
				}
			}

			require.ErrorIs(t, err, io.EOF)
			require.Zero(t, n)
		})
	}
}
