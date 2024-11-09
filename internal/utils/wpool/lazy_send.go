package wpool

import "github.com/glebziz/fs_db/internal/model/core"

func (p *pool) lazySend(e Event) {
	p.listM.Lock()
	defer p.listM.Unlock()

	p.el.PushBack((&core.Node[Event]{}).SetV(e)) // TODO use pool
	p.lazyResend()
}

func (p *pool) lazyResend() {
	if !p.lazySendM.TryLock() {
		return
	}

	p.sendWg.Add(1)
	go func() {
		defer func() {
			p.lazySendM.Unlock()
			p.sendWg.Done()
		}()

		for {
			p.listM.Lock()
			n := p.el.PopBack()
			p.listM.Unlock()
			if n == nil {
				return
			}

			select {
			case <-p.ctx.Done():
				return
			case p.ch <- n.V():
				// TODO free node
			}
		}
	}()
}
