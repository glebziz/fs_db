package wpool

import (
	"log/slog"
)

func (p *Pool) Stop() {
	defer p.runA.Store(false)
	if !p.runA.CompareAndSwap(true, false) {
		slog.Warn("worker pool already stopped")
		return
	}

	p.cancel()
	p.listCv.Broadcast()
	p.sendWg.Wait()
	p.runWg.Wait()

	close(p.ch)
	p.el.Clear()
}
