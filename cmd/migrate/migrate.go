package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func migrateAll(ctx context.Context, db *sql.DB, r repos) error {
	fmt.Println("Start migrations.")

	return r.fRepo.RunTransaction(ctx, func(ctx context.Context) error {
		err := migrateFiles(ctx, db, r)
		if err != nil {
			return fmt.Errorf("migrate files: %w", err)
		}

		err = migrateContentFiles(ctx, db, r)
		if err != nil {
			return fmt.Errorf("migrate content files: %w", err)
		}

		return nil
	})
}

func migrateFiles(ctx context.Context, db *sql.DB, r repos) error {
	fmt.Println("Migrate files.")

	rows, err := db.QueryContext(ctx, `select key, content_id
		from (
			select key, content_id, rank() over (partition by key order by ts desc) r
      		from file
      		where tx_id = $1
      	)
		where r = 1`, model.MainTxId,
	)
	if err != nil {
		return fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			key       string
			contentId sql.Null[string]
		)

		err = rows.Scan(&key, &contentId)
		if err != nil {
			return fmt.Errorf("row scan: %w", err)
		}

		if !contentId.Valid {
			contentId.V = uuid.NewString()
		}

		err = r.fRepo.Set(ctx, model.File{
			Key:       key,
			TxId:      model.MainTxId,
			ContentId: contentId.V,
			Seq:       sequence.Next(),
		})
		if err != nil {
			return fmt.Errorf("set file: %w", err)
		}
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("rows err: %w", err)
	}

	return nil
}

func migrateContentFiles(ctx context.Context, db *sql.DB, r repos) error {
	fmt.Println("Migrate file contents.")

	rows, err := db.QueryContext(ctx, `select id, parent_path from content_file`)
	if err != nil {
		return fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cFile model.ContentFile
		err = rows.Scan(&cFile.Id, &cFile.Parent)
		if err != nil {
			return fmt.Errorf("row scan: %w", err)
		}

		err = r.cfRepo.Store(ctx, cFile)
		if err != nil {
			return fmt.Errorf("store content file: %w", err)
		}
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("rows err: %w", err)
	}

	return nil
}
