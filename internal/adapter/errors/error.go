package errors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glebziz/fs_db"
	store "github.com/glebziz/fs_db/internal/proto"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

func Error(err error) error {
	var code codes.Code
	switch {
	case errors.Is(err, fs_db.ErrNoFreeSpace):
		code = codes.ResourceExhausted
	case errors.Is(err, fs_db.ErrNotFound):
		code = codes.NotFound
	case errors.Is(err, fs_db.ErrEmptyKey):
		code = codes.InvalidArgument
	case errors.Is(err, fs_db.ErrTxNotFound):
		code = codes.Aborted
	case errors.Is(err, fs_db.ErrTxAlreadyExists):
		code = codes.AlreadyExists
	case errors.Is(err, fs_db.ErrTxSerialization):
		code = codes.FailedPrecondition
	default:
		code = codes.Internal
	}

	st, _ := status.New(code, err.Error()).WithDetails(errorToPbError(err))
	return st.Err()
}

func ClientError(err error) error {
	st := status.Convert(err)

	err = detailsToError(st.Details())
	if err != nil {
		return err
	}

	switch st.Code() { //nolint:exhaustive
	case codes.InvalidArgument:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.ErrEmptyKey)
	case codes.NotFound:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.ErrNotFound)
	case codes.AlreadyExists:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.ErrTxAlreadyExists)
	case codes.ResourceExhausted:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.ErrNoFreeSpace)
	case codes.FailedPrecondition:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.ErrTxSerialization)
	case codes.Aborted:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.ErrTxNotFound)
	case codes.Internal:
		return fmt.Errorf("%s: %w", st.Message(), fs_db.ErrUnknown)
	default:
		return errors.Join(err, fs_db.ErrUnknown)
	}
}

func errorToPbError(err error) *store.Error {
	var errCode store.ErrorCode
	switch {
	case errors.Is(err, fs_db.ErrNoFreeSpace):
		errCode = store.ErrorCode_ErrNoFreeSpace
	case errors.Is(err, fs_db.ErrNotFound):
		errCode = store.ErrorCode_ErrNotFound
	case errors.Is(err, fs_db.ErrEmptyKey):
		errCode = store.ErrorCode_ErrEmptyKey
	case errors.Is(err, fs_db.ErrHeaderNotFound):
		errCode = store.ErrorCode_ErrHeaderNotFound
	case errors.Is(err, fs_db.ErrTxNotFound):
		errCode = store.ErrorCode_ErrTxNotFound
	case errors.Is(err, fs_db.ErrTxAlreadyExists):
		errCode = store.ErrorCode_ErrTxAlreadyExists
	case errors.Is(err, fs_db.ErrTxSerialization):
		errCode = store.ErrorCode_ErrTxSerialization
	}

	return &store.Error{
		Code:    errCode,
		Message: ptr.Ptr(err.Error()),
	}
}

func detailsToError(d []any) error { //nolint:cyclop
	for _, e := range d {
		err, ok := e.(*store.Error)
		if !ok {
			continue
		}

		switch err.GetCode() {
		case store.ErrorCode_ErrUnknown:
			return fmt.Errorf("%s: %w", err.GetMessage(), fs_db.ErrUnknown)
		case store.ErrorCode_ErrNoFreeSpace:
			return fmt.Errorf("%s: %w", err.GetMessage(), fs_db.ErrNoFreeSpace)
		case store.ErrorCode_ErrNotFound:
			return fmt.Errorf("%s: %w", err.GetMessage(), fs_db.ErrNotFound)
		case store.ErrorCode_ErrEmptyKey:
			return fmt.Errorf("%s: %w", err.GetMessage(), fs_db.ErrEmptyKey)
		case store.ErrorCode_ErrHeaderNotFound:
			return fmt.Errorf("%s: %w", err.GetMessage(), fs_db.ErrHeaderNotFound)
		case store.ErrorCode_ErrTxNotFound:
			return fmt.Errorf("%s: %w", err.GetMessage(), fs_db.ErrTxNotFound)
		case store.ErrorCode_ErrTxAlreadyExists:
			return fmt.Errorf("%s: %w", err.GetMessage(), fs_db.ErrTxAlreadyExists)
		case store.ErrorCode_ErrTxSerialization:
			return fmt.Errorf("%s: %w", err.GetMessage(), fs_db.ErrTxSerialization)
		}
	}

	return nil
}
