package di

import (
	"github.com/glebziz/fs_db/internal/usecase/transaction"
)

func (c *Container) Transaction() *transaction.UseCase {
	if c.transactionUseCase == nil {
		c.transactionUseCase = transaction.New(
			c.Cleaner(),
			c.Core(),
			c.TransactionRepo(),
			c.Gen(),
		)
	}

	return c.transactionUseCase
}
