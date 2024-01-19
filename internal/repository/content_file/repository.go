package file

import (
	"fmt"
	"strings"

	"github.com/glebziz/fs_db/internal/db"
)

type rep struct {
	p db.Provider
}

func New(p db.Provider) *rep {
	return &rep{p}
}

type arrayType interface {
	~int | ~string | ~float64
}

func arrayArg[T arrayType](arr []T) (in string, args []interface{}) {
	args = make([]interface{}, 0, len(arr))
	str := strings.Builder{}

	for i, s := range arr {
		str.WriteString(fmt.Sprintf("$%d,", i))
		args = append(args, s)
	}

	return strings.Trim(str.String(), ","), args
}
