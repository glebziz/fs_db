package transactor

import "context"

type TransactionFn func(ctx context.Context) error

type Transactor interface {
	RunTransaction(ctx context.Context, fn TransactionFn) error
}
