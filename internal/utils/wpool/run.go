package wpool

import (
	"context"
	"log/slog"
)

func (p *pool) Run(ctx context.Context) {
	if !p.runM.TryLock() {
		slog.Warn("worker pool already running")
		return
	}

	p.ctx, p.cancel = context.WithCancel(ctx)
	p.ch = make(chan Event, p.opts.NumWorkers*2) //nolint:mnd
	for range p.opts.NumWorkers {
		p.runWg.Add(1)
		go p.run() //nolint:contextcheck
	}
}

func (p *pool) run() {
	defer p.runWg.Done()
	for {
		select {
		case <-p.ctx.Done():
			return
		case e := <-p.ch:
			p.exec(e)
		}
	}
}

func (p *pool) exec(e Event) {
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
