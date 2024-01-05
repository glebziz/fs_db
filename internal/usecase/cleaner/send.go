package cleaner

import (
	"log/slog"
	"time"
)

func (c *Cleaner) Send(contentIds []string) {
	if c.closed.Load() {
		slog.Warn("cleaner is stopped",
			"contentIds", contentIds,
		)
		return
	}

	if len(contentIds) == 0 {
		return
	}

	c.sWg.Add(1)
	go func() {
		select {
		case c.ch <- contentIds:
		case <-time.After(sendTimeout):
			slog.Warn(
				"failed to send files",
				"contentIds", contentIds,
			)
		}

		c.sWg.Done()
	}()
}
