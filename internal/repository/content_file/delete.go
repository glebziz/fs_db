package file

import (
	"context"
	"fmt"
)

func (r *rep) Delete(ctx context.Context, ids []string) error {
	in, args := arrayArg(ids)
	_, err := r.p.DB(ctx).Exec(ctx, fmt.Sprintf(`
		delete from content_file
		where id in (%s)`, in), args...)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}
