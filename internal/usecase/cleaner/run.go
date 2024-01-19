package cleaner

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/reugn/go-quartz/job"
	"github.com/reugn/go-quartz/quartz"
)

func (c *Cleaner) Run(ctx context.Context) error {
	err := c.sched.ScheduleJob(
		quartz.NewJobDetail(
			job.NewFunctionJob(func(ctx context.Context) (any, error) {
				err := c.deleteLines(ctx)
				if err != nil {
					slog.Warn("delete lines",
						slog.Any("err", err),
					)
				}

				return nil, nil
			}),
			quartz.NewJobKey("cleanAll"),
		),
		quartz.NewSimpleTrigger(cleanPeriod),
	)
	if err != nil {
		return fmt.Errorf("schedule cleanAll job: %w", err)
	}

	c.sched.Start(ctx)
	return nil
}
