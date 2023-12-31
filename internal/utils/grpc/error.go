package grpc

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glebziz/fs_db"
)

func Error(err error) error {
	var code codes.Code
	switch {
	case errors.Is(err, fs_db.SizeErr):
		code = codes.ResourceExhausted
	case errors.Is(err, fs_db.NotFoundErr):
		code = codes.NotFound
	case errors.Is(err, fs_db.EmptyKeyErr):
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
		return fmt.Errorf("%s: %w", st.Message(), fs_db.EmptyKeyErr)
	case codes.NotFound:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.NotFoundErr)
	case codes.ResourceExhausted:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.SizeErr)
	case codes.Internal:
		return fmt.Errorf("%s", st.Message())
	default:
		return err
	}
}
