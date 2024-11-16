package wpool

import (
	"context"
	"time"
)

const (
	minNumWorkers   = 1
	minSendDuration = 1
)

type Options struct {
	NumWorkers   int
	SendDuration time.Duration
}

type RunFunc func(ctx context.Context) error

type Event struct {
	ctx    context.Context
	Caller string
	Fn     RunFunc
}
