package wpool

import (
	"context"
	"log/slog"
	"sync"
)

func (p *Pool) Run(ctx context.Context) {
	if !p.runA.CompareAndSwap(false, true) {
		slog.Warn("worker pool already running")
		return
	}

	p.listCv = sync.NewCond(new(sync.Mutex))
	p.ctx, p.cancel = context.WithCancel(ctx)
	p.ch = make(chan Event, p.opts.numWorkers*2) //nolint:mnd
	for range p.opts.numWorkers {
		p.runWg.Go(p.run)
	}
}

func (p *Pool) run() {
	for {
		select {
		case <-p.ctx.Done():
			return
		case e := <-p.ch:
			p.exec(e)
		}
	}
}

func (p *Pool) exec(e Event) {
	ctx, cancel := context.WithCancel(e.ctx)
	stop := context.AfterFunc(p.ctx, func() {
		cancel()
	})
	defer func() {
		stop()
		cancel()
	}()

	err := e.Fn(ctx)
	if err != nil {
		slog.Error("the run function failed with an error",
			slog.String("caller", e.Caller),
			slog.String("err", err.Error()),
		)
	}
}
