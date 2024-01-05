package cleaner

func (c *Cleaner) Stop() {
	c.closed.Store(true)
	c.sWg.Wait()
	close(c.ch)
	c.rWg.Wait()
}
