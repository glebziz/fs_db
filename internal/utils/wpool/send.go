package wpool

import (
	"context"
	"time"
)

func (p *Pool) Send(ctx context.Context, e Event) {
	e.ctx = ctx

	p.sendWg.Add(1)
	defer p.sendWg.Done()

	if p.ctx.Err() != nil {
		return
	}

	select {
	case <-p.ctx.Done():
		return
	case p.ch <- e:
	case <-time.After(p.opts.SendDuration):
		p.lazySend(e)
	}
}
