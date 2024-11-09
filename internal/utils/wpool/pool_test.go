package wpool

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	t.Run("double run and double stop", func(t *testing.T) {
		p := New(Options{})
		p.Run(context.Background())
		p.Run(context.Background())
		p.Stop()
		p.Stop()
	})
	t.Run("main scenario", func(t *testing.T) {
		p := New(Options{})
		p.Run(context.Background())
		p.Send(context.Background(), Event{
			Caller: "Caller",
			Fn: func(ctx context.Context) error {
				<-ctx.Done()
				return nil
			},
		})
		p.Stop()
	})
	t.Run("schedule work", func(t *testing.T) {
		p := New(Options{})
		p.Run(context.Background())
		wg := sync.WaitGroup{}
		wg.Add(2)
		p.Sched(context.Background(), Event{
			Caller: "Caller",
			Fn: func(ctx context.Context) error {
				wg.Done()
				return assert.AnError
			},
		}, time.Millisecond)
		wg.Wait()
		p.Stop()
	})
	t.Run("send after stop", func(t *testing.T) {
		p := New(Options{})
		p.Run(context.Background())
		p.Stop()
		p.Send(context.Background(), Event{
			Caller: "Caller",
			Fn: func(ctx context.Context) error {
				return nil
			},
		})
	})
	t.Run("stop when send", func(t *testing.T) {
		p := New(Options{
			SendDuration: time.Second,
		})
		p.Run(context.Background())
		p.Send(context.Background(), Event{
			Caller: "Caller",
			Fn: func(ctx context.Context) error {
				<-ctx.Done()
				return nil
			},
		})

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(250 * time.Millisecond)
			p.Stop()
		}()
		for range 3 {
			p.Send(context.Background(), Event{
				Caller: "Caller",
				Fn: func(ctx context.Context) error {
					return nil
				},
			})
		}
		wg.Wait()
	})
	t.Run("use lazy send", func(t *testing.T) {
		p := New(Options{})
		p.Run(context.Background())
		firstWg := sync.WaitGroup{}
		firstWg.Add(1)
		p.Send(context.Background(), Event{
			Caller: "Caller",
			Fn: func(ctx context.Context) error {
				firstWg.Wait()
				return nil
			},
		})
		wg := sync.WaitGroup{}
		for range 5 {
			wg.Add(1)
			p.Send(context.Background(), Event{
				Caller: "Caller",
				Fn: func(ctx context.Context) error {
					defer wg.Done()
					return nil
				},
			})
		}

		firstWg.Done()
		wg.Wait()
		p.Stop()
	})
	t.Run("cancel lazy send", func(t *testing.T) {
		p := New(Options{})
		p.Run(context.Background())
		p.Send(context.Background(), Event{
			Caller: "Caller",
			Fn: func(ctx context.Context) error {
				<-ctx.Done()
				return nil
			},
		})
		for range 4 {
			p.Send(context.Background(), Event{
				Caller: "Caller",
				Fn: func(ctx context.Context) error {
					return nil
				},
			})
		}
		p.Stop()
	})
}
