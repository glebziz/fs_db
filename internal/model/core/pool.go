package core

import (
	"sync"
)

type ClearFunc[T any] func(e *T)

type Pool[T any] struct {
	m      sync.Mutex
	free   []*T
	clearF ClearFunc[T]
}

func NewPool[T any](clearF ClearFunc[T]) *Pool[T] {
	return &Pool[T]{
		clearF: clearF,
	}
}

func (p *Pool[T]) Acquire() *T {
	p.m.Lock()
	defer p.m.Unlock()

	if len(p.free) == 0 {
		return new(T)
	}

	e := p.free[len(p.free)-1]
	p.free = p.free[:len(p.free)-1]

	return e
}

func (p *Pool[T]) Release(els ...*T) {
	p.m.Lock()
	defer p.m.Unlock()

	for _, e := range els {
		if p.clearF != nil {
			p.clearF(e)
		} else {
			*e = *new(T)
		}
	}

	p.free = append(p.free, els...)
}
