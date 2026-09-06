package store

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (u *UseCase) Set(ctx context.Context, key string, contents model.Contents) error {
	if key == "" {
		return fs_db.ErrEmptyKey
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
			TxId:      model.GetTxId(ctx),
			ContentId: cFile.Id,
		}
	)
	err = u.writeContent(ctx, dirs, &cFile, contents)
	if err != nil {
		return fmt.Errorf("write content: %w", err)
	}

	err = u.cfRepo.Store(ctx, cFile)
	if err != nil {
		return fmt.Errorf("content file repository store: %w", err)
	}

	err = u.fRepo.Store(ctx, file)
	if err != nil {
		return fmt.Errorf("file repository store: %w", err)
	}

	return nil
}

func (u *UseCase) writeContent(ctx context.Context, dirs model.Dirs, cFile *model.ContentFile, contents model.Contents) error {
	var (
		err     error
		minSize uint64
		closer  io.Closer

		errNotEnoughSpace model.NotEnoughSpaceError
	)
	defer func() {
		if closer == nil {
			return
		}

		closer.Close()
	}()
	for dir, ok := range dirs.Iterate(u.randGen) {
		if !ok {
			return fs_db.ErrNoFreeSpace
		}

		if dir.Free <= minSize {
			continue
		}

		cFile.Parent = dir.Path()
		err = u.cWriter.Write(ctx, cFile.Path(), contents)
		if err != nil {
			if errors.As(err, &errNotEnoughSpace) {
				if closer != nil {
					closer.Close()
				}

				closer = errNotEnoughSpace
				contents = contents.InsertAtStart(ctx, model.Content{
					Reader: errNotEnoughSpace.Reader(),
				})
				minSize = dir.Free
				continue
			}

			return fmt.Errorf("content writer write: %w", err)
		}

		break
	}

	return nil
}
