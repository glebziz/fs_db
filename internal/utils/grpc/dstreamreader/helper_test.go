package dstreamreader_test

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/glebziz/fs_db/internal/model"
	mock_dstreamreader "github.com/glebziz/fs_db/internal/utils/grpc/dstreamreader/mocks"
)

type tStream = *mock_dstreamreader.MockStream[*TestRequest, *mock_dstreamreader.MockResponse]

type TestRequest struct {
	Seek model.Seek
}

func (tr *TestRequest) ProtoReflect() protoreflect.Message {
	return nil
}
