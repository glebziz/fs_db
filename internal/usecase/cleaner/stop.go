package cleaner

import (
	"context"
	"fmt"
)

func (c *Cleaner) Stop(ctx context.Context) error {
	c.sched.Stop()
	c.sched.Wait(ctx)

	err := c.sched.Clear()
	if err != nil {
		return fmt.Errorf("sched clear: %w", err)
	}

	return nil
}
