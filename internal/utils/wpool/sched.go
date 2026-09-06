package wpool

import (
	"context"
	"time"
)

func (p *Pool) Sched(ctx context.Context, e Event, period time.Duration) {
	go func() {
		timer := time.NewTimer(period)
		defer timer.Stop()

		for {
			select {
			case <-p.ctx.Done():
				return
			case <-timer.C:
				timer.Reset(period)
				p.Send(ctx, e)
			}
		}
	}()
}
