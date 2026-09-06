package wpool

import (
	"context"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	t.Parallel()

	t.Run("double run and double stop", func(t *testing.T) {
		t.Parallel()

		p := New()

		p.Run(t.Context())
		p.Run(t.Context())

		p.Stop()
		p.Stop()
	})

	t.Run("main scenario", func(t *testing.T) {
		t.Parallel()

		synctest.Test(t, func(t *testing.T) {
			p := New()

			p.Run(t.Context())
			p.Send(t.Context(), Event{
				Caller: "Caller",
				Fn: func(ctx context.Context) error {
					<-ctx.Done()
					return nil
				},
			})

			p.Stop()
		})
	})

	t.Run("schedule work", func(t *testing.T) {
		t.Parallel()

		synctest.Test(t, func(t *testing.T) {
			p := New()

			p.Run(t.Context())

			var wg sync.WaitGroup
			wg.Add(2)
			p.Sched(t.Context(), Event{
				Caller: "Caller",
				Fn: func(ctx context.Context) error {
					wg.Done()
					return assert.AnError
				},
			}, time.Second)

			wg.Wait()
			p.Stop()
		})
	})

	t.Run("send after stop", func(t *testing.T) {
		t.Parallel()

		p := New()

		p.Run(t.Context())
		p.Stop()

		p.Send(t.Context(), Event{
			Caller: "Caller",
			Fn: func(ctx context.Context) error {
				return nil
			},
		})
	})

	t.Run("stop when send", func(t *testing.T) {
		t.Parallel()

		synctest.Test(t, func(t *testing.T) {
			p := New(
				WithSendDuration(time.Second),
			)

			p.Run(t.Context())
			p.Send(t.Context(), Event{
				Caller: "Caller",
				Fn: func(ctx context.Context) error {
					<-ctx.Done()
					return nil
				},
			})

			go func() {
				time.Sleep(250 * time.Millisecond)
				p.Stop()
			}()

			for range 3 {
				p.Send(t.Context(), Event{
					Caller: "Caller",
					Fn: func(ctx context.Context) error {
						return nil
					},
				})
			}
		})
	})

	t.Run("use lazy send", func(t *testing.T) {
		t.Parallel()

		synctest.Test(t, func(t *testing.T) {
			p := New(
				WithSendDuration(time.Second),
			)
			p.Run(t.Context())

			p.Send(t.Context(), Event{
				Caller: "Caller",
				Fn: func(ctx context.Context) error {
					time.Sleep(10 * time.Second)
					return nil
				},
			})

			var wg sync.WaitGroup
			for range 5 {
				wg.Add(1)
				p.Send(t.Context(), Event{
					Caller: "Caller",
					Fn: func(ctx context.Context) error {
						defer wg.Done()
						return nil
					},
				})
			}

			wg.Wait()
			p.Stop()
		})
	})

	t.Run("cancel lazy send", func(t *testing.T) {
		t.Parallel()

		synctest.Test(t, func(t *testing.T) {
			p := New()

			p.Run(t.Context())
			p.Send(t.Context(), Event{
				Caller: "Caller",
				Fn: func(ctx context.Context) error {
					<-ctx.Done()
					return nil
				},
			})

			for range 4 {
				p.Send(t.Context(), Event{
					Caller: "Caller",
					Fn: func(ctx context.Context) error {
						return nil
					},
				})
			}

			p.Stop()
		})
	})
}
