package grpc

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glebziz/fs_db/pkg/model"
)

func Error(err error) error {
	var code codes.Code
	switch {
	case errors.Is(err, model.SizeErr):
		code = codes.ResourceExhausted
	case errors.Is(err, model.NotFoundErr):
		code = codes.NotFound
	case errors.Is(err, model.EmptyKeyErr):
		code = codes.InvalidArgument
	default:
		code = codes.Internal
	}

	return status.Error(code, err.Error())
}

func ClientError(err error) error {
	st := status.Convert(err)
	switch st.Code() {
	case codes.InvalidArgument:
		return fmt.Errorf("%s: %w", st.Message(), model.EmptyKeyErr)
	case codes.NotFound:
		return fmt.Errorf("%s: %w", st.Message(), model.NotFoundErr)
	case codes.ResourceExhausted:
		return fmt.Errorf("%s: %w", st.Message(), model.SizeErr)
	case codes.Internal:
		return fmt.Errorf("%s", st.Message())
	default:
		return err
	}
}
