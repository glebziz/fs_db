package store

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
)

func (u *useCase) Set(ctx context.Context, key string, content *model.Content) error {
	if key == "" {
		return fs_db.EmptyKeyErr
	}

	dir, err := u.dir.Select(ctx, content.Size)
	if err != nil {
		return fmt.Errorf("get parent dir: %w", err)
	}

	var (
		cFile = model.ContentFile{
			Id:         u.idGen.Generate(),
			ParentPath: dir.GetPath(),
		}
		file = model.File{
			Key:       key,
			ContentId: cFile.Id,
		}
	)

	err = u.cRepo.Store(ctx, cFile.GetPath(), content)
	if err != nil {
		return fmt.Errorf("content repository store: %w", err)
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
