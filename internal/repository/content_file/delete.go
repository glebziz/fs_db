package file

import (
	"context"
	"fmt"
)

func (r *rep) Delete(ctx context.Context, id string) error {
	err := r.p.DB(ctx).Delete(r.key(id))
	if err != nil {
		return fmt.Errorf("db delete: %w", err)
	}

	return nil
}
