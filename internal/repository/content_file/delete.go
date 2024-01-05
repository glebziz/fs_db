package file

import (
	"context"
	"fmt"
	"strings"

	"github.com/glebziz/fs_db/internal/model"
)

func (r *rep) Delete(ctx context.Context, ids []string) ([]model.ContentFile, error) {
	in, args := r.stringArrayArg(ids)
	rows, err := r.p.DB(ctx).Query(ctx, fmt.Sprintf(`
		delete from content_file
		where id in (%s)
		returning id, parent_path`, in), args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var files []model.ContentFile
	for rows.Next() {
		var file model.ContentFile

		err = rows.Scan(&file.Id, &file.ParentPath)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		files = append(files, file)
	}

	return files, nil
}

func (_ *rep) stringArrayArg(arr []string) (in string, args []interface{}) {
	args = make([]interface{}, 0, len(arr))
	str := strings.Builder{}

	for i, s := range arr {
		str.WriteString(fmt.Sprintf("$%d,", i))
		args = append(args, s)
	}

	return strings.Trim(str.String(), ","), args
}
