package cleaner

import (
	"context"
	"log/slog"
	"sync"
)

func (c *Cleaner) Run() {
	c.rWg.Add(1)
	go func() {
		ctx := context.Background()
		lWg := sync.WaitGroup{}
		for contentIds := range c.ch {
			lWg.Add(1)

			go func(contentIds []string) {
				defer lWg.Done()

				cfs, err := c.cfRepo.Delete(ctx, contentIds)
				if err != nil {
					slog.Error("content file repo delete",
						"err", err,
						"contentIds", contentIds,
					)
					return
				}

				for _, cf := range cfs {
					err = c.cRepo.Delete(context.Background(), cf.GetPath())
					if err != nil {
						slog.Error("content repo delete",
							"err", err,
							"content file", cf,
						)
					}
				}
			}(contentIds)
		}

		lWg.Wait()
		c.rWg.Done()
	}()
}
