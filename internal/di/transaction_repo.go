package di

import (
	transactionRepo "github.com/glebziz/fs_db/internal/repository/transaction"
)

func (c *Container) TransactionRepo() *transactionRepo.Repo {
	if c.transactionRepo == nil {
		c.transactionRepo = transactionRepo.New()
	}

	return c.transactionRepo
}
