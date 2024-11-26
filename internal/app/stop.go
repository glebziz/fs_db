package app

import (
	"fmt"
)

func (a *app) Stop() error {
	a.container.Pool().Stop()

	err := a.container.Badger().Close()
	if err != nil {
		return fmt.Errorf("badger close: %w", err)
	}

	return nil
}
