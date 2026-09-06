package wpool

import (
	"context"
	"time"
)

const (
	minNumWorkers   = 1
	minSendDuration = time.Nanosecond
)

type options struct {
	numWorkers   int
	sendDuration time.Duration
}

type OptionFunc func(*options)

type RunFunc func(ctx context.Context) error

type Event struct {
	ctx    context.Context
	Caller string
	Fn     RunFunc
}

func WithNumWorkers(numWorkers int) OptionFunc {
	return func(o *options) {
		o.numWorkers = max(numWorkers, minNumWorkers)
	}
}

func WithSendDuration(sendDuration time.Duration) OptionFunc {
	return func(o *options) {
		o.sendDuration = max(sendDuration, minSendDuration)
	}
}
