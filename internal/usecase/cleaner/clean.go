package cleaner

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/reugn/go-quartz/job"
	"github.com/reugn/go-quartz/quartz"
)

func (c *Cleaner) Clean(contentIds []string) error {
	err := c.sched.ScheduleJob(
		quartz.NewJobDetail(
			job.NewFunctionJob(func(ctx context.Context) (any, error) {
				err := c.deleteContent(ctx, contentIds)
				if err != nil {
					slog.Warn("delete lines",
						slog.Any("err", err),
						slog.Any("contentIds", contentIds),
					)
				}

				return nil, nil
			}),
			quartz.NewJobKey(fmt.Sprintf("clean-%s", uuid.NewString())),
		),
		quartz.NewRunOnceTrigger(0),
	)
	if err != nil {
		return fmt.Errorf("schedule job: %w", err)
	}

	return nil
}
