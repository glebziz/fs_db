package wpool

import (
	"context"
	"time"
)

func (p *pool) Sched(ctx context.Context, e Event, period time.Duration) {
	go func() {
		for {
			select {
			case <-p.ctx.Done():
				return
			case <-time.After(period):
				p.Send(ctx, e)
			}
		}
	}()
}
