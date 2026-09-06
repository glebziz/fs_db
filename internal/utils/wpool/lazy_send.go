package wpool

func (p *Pool) lazySend(e Event) {
	p.listCv.L.Lock()
	defer p.listCv.L.Unlock()

	p.el.PushBack(p.pool.Acquire().SetV(e))
	p.listCv.Signal()
	p.lazyResend()
}

func (p *Pool) lazyResend() {
	if !p.lazySendA.CompareAndSwap(false, true) {
		return
	}

	p.sendWg.Go(func() {
		for {
			p.listCv.L.Lock()
			n := p.el.PopBack()
			if n == nil {
				p.listCv.Wait()
				n = p.el.PopBack()
			}
			p.listCv.L.Unlock()

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
	})
}
