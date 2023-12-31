package db_pool

import "database/sql"

type Rows struct {
	*sql.Rows
}
