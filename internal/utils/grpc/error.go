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
	case errors.Is(err, fs_db.TxNotFoundErr):
		code = codes.Aborted
	case errors.Is(err, fs_db.TxAlreadyExistsErr):
		code = codes.AlreadyExists
	case errors.Is(err, fs_db.TxSerializationErr):
		code = codes.FailedPrecondition
	default:
		code = codes.Internal
	}

	return status.Error(code, err.Error())
}

func ClientError(err error) error {
	st := status.Convert(err)
	switch st.Code() { //nolint:exhaustive
	case codes.InvalidArgument:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.EmptyKeyErr)
	case codes.NotFound:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.NotFoundErr)
	case codes.AlreadyExists:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.TxAlreadyExistsErr)
	case codes.ResourceExhausted:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.SizeErr)
	case codes.FailedPrecondition:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.TxSerializationErr)
	case codes.Aborted:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.TxNotFoundErr)
	case codes.Internal:
		return fmt.Errorf("%s", st.Message()) //nolint:err113
	default:
		return err
	}
}
