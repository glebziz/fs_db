package core

import (
	"sync"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
	"github.com/glebziz/fs_db/internal/utils/ptr"
)

type NextFunc func() *Node[model.File]

var (
	emptyNextFunc = func() *Node[model.File] {
		return nil
	}
)

type file struct {
	m             sync.RWMutex
	l             List[model.File]
	arr           []*Node[model.File]
	withoutSearch bool
}

func (f *file) Lock() {
	if f == nil {
		return
	}

	f.m.Lock()
}

func (f *file) RLock() {
	if f == nil {
		return
	}

	f.m.RLock()
}

func (f *file) Unlock() {
	if f == nil {
		return
	}

	f.m.Unlock()
}

func (f *file) RUnlock() {
	if f == nil {
		return
	}

	f.m.RUnlock()
}

func (f *file) PushBack(n *Node[model.File]) {
	if f == nil || n == nil {
		return
	}

	if !f.withoutSearch {
		f.arr = append(f.arr, n)
	}

	f.l.PushBack(n)
}

func (f *file) PopBack() *Node[model.File] {
	if f == nil {
		return nil
	}

	n := f.l.PopBack()
	if n != nil && !f.withoutSearch {
		f.arr = f.arr[:len(f.arr)-1]
	}

	return n
}

func (f *file) PopFront() *Node[model.File] {
	if f == nil {
		return nil
	}

	n := f.l.PopFront()
	if n != nil && !f.withoutSearch {
		copy(f.arr, f.arr[1:])
		f.arr = f.arr[:len(f.arr)-1]
	}

	return n
}

func (f *file) Latest() model.File {
	if f == nil {
		return model.File{}
	}

	return f.l.Back().V()
}

func (f *file) LastBefore(seq sequence.Seq) model.File {
	if f == nil || len(f.arr) == 0 {
		return model.File{}
	}

	return ptr.Val(binarySearch(f.arr, seq)).v
}

func (f *file) IterateBeforeSeq(seq sequence.Seq) NextFunc {
	if f == nil || f.l.IsEmpty() {
		return emptyNextFunc
	}

	n := f.l.Front()
	return func() *Node[model.File] {
		v := n.next.v
		if v.Seq.Zero() || v.Seq.After(seq) {
			return nil
		}
		defer func() {
			n = n.next
		}()

		return n
	}
}

func binarySearch(arr []*Node[model.File], seq sequence.Seq) *Node[model.File] {
	var (
		n int
	)
	for len(arr) > 0 {
		n = len(arr) / 2
		if !arr[n].v.Seq.Before(seq) {
			arr = arr[:n]
		} else {
			if n == len(arr)-1 || !arr[n+1].v.Seq.Before(seq) {
				return arr[n]
			}

			arr = arr[n+1:]
		}
	}

	return nil
}
