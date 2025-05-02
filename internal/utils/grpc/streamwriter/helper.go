//go:build test

package streamwriter

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/glebziz/fs_db/internal/model"
)

//go:generate mockgen -source helper.go -destination mocks/mocks.go -typed true

type TestRequest struct {
	P    []byte
	Seek model.Seek
}

type StreamTest interface {
	Send(req *TestRequest) error
	CloseAndRecv() (*TestRequest, error)
}

func (tr *TestRequest) ProtoReflect() protoreflect.Message {
	return nil
}
