package wpool

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/glebziz/fs_db/internal/model/core"
)

type Pool struct {
	ctx    context.Context
	cancel context.CancelFunc

	runA      atomic.Bool
	lazySendA atomic.Bool
	listCv    *sync.Cond

	el   core.List[Event]
	pool core.Pool[core.Node[Event]]

	ch     chan Event
	sendWg sync.WaitGroup
	runWg  sync.WaitGroup

	opts options
}

func New(opts ...OptionFunc) *Pool {
	p := Pool{
		opts: options{
			numWorkers:   minNumWorkers,
			sendDuration: minSendDuration,
		},
	}

	for _, f := range opts {
		f(&p.opts)
	}

	return &p
}
