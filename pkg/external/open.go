package external

import (
	"context"
	"fmt"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/pkg/external/db"
)

// Open returns the client for the external fs db.
func Open(ctx context.Context, url string, opt ...db.OptionFunc) (fs_db.DB, error) {
	b, err := db.New(ctx, url, opt...)
	if err != nil {
		return nil, fmt.Errorf("new db: %w", err)
	}

	return b, nil
}

// WithCertPath applies the certPath option.
func WithCertPath(certPath string) db.OptionFunc {
	return db.WithCertPath(certPath)
}

// WithServerNameOverride applies the serverNameOverride option.
func WithServerNameOverride(serverNameOverride string) db.OptionFunc {
	return db.WithServerNameOverride(serverNameOverride)
}
