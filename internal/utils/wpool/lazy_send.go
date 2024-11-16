package wpool

func (p *pool) lazySend(e Event) {
	p.listM.Lock()
	defer p.listM.Unlock()

	p.el.PushBack(p.pool.Acquire().SetV(e))
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
				p.pool.Release(n)
			}
		}
	}()
}
