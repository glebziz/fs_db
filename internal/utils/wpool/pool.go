package wpool

import (
	"context"
	"sync"

	"github.com/glebziz/fs_db/internal/model/core"
)

type Pool struct {
	ctx    context.Context
	cancel context.CancelFunc

	runM      sync.Mutex
	lazySendM sync.Mutex
	listM     sync.Mutex

	el   core.List[Event]
	pool core.Pool[core.Node[Event]]

	ch     chan Event
	sendWg sync.WaitGroup
	runWg  sync.WaitGroup

	opts Options
}

func New(options Options) *Pool {
	return &Pool{
		opts: Options{
			NumWorkers:   max(options.NumWorkers, minNumWorkers),
			SendDuration: max(options.SendDuration, minSendDuration),
		},
	}
}
