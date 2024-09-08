package store

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Set(ctx context.Context, key string, content io.Reader) error {
	if key == "" {
		return fs_db.EmptyKeyErr
	}

	dirs, err := u.dir.Get(ctx)
	if err != nil {
		return fmt.Errorf("get parent dir: %w", err)
	}

	var (
		cFile = model.ContentFile{
			Id: u.idGen.Generate(),
		}
		file = model.File{
			Key:       key,
			ContentId: cFile.Id,
		}
	)

	var (
		minSize uint64
		closer  io.Closer

		nextFunc = dirs.Iterate(u.randGen)
	)
	for {
		dir, ok := nextFunc()
		if !ok {
			return fs_db.SizeErr
		}

		if dir.Free <= minSize {
			continue
		}

		cFile.Parent = dir.Path()
		err = u.cRepo.Store(ctx, cFile.Path(), content)
		if err != nil {
			var errNotEnoughSpace model.NotEnoughSpaceError
			if errors.As(err, &errNotEnoughSpace) {
				if closer != nil {
					closer.Close()
				}

				closer = errNotEnoughSpace
				content = errNotEnoughSpace.Reader()
				minSize = dir.Free
				continue
			}

			return fmt.Errorf("content repository store: %w", err)
		}

		break
	}

	if closer != nil {
		closer.Close()
	}

	err = u.cfRepo.Store(ctx, cFile)
	if err != nil {
		return fmt.Errorf("content file repository store: %w", err)
	}

	txId := model.GetTxId(ctx)
	err = u.fRepo.Store(ctx, txId, file)
	if err != nil {
		return fmt.Errorf("file repository store: %w", err)
	}

	return nil
}
