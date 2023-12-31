package db_pool

import "database/sql"

type Conn struct {
	*sql.DB
	pool chan<- *Conn
}

func (c *Conn) Release() {
	c.pool <- c
}
