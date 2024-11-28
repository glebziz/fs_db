package wpool

import (
	"log/slog"
)

func (p *Pool) Stop() {
	defer p.runM.Unlock()
	if p.runM.TryLock() {
		slog.Warn("worker pool already stopped")
		return
	}

	p.cancel()
	p.sendWg.Wait()
	p.runWg.Wait()

	close(p.ch)
	p.el.Clear()
}
