package sequence

import "sync/atomic"

type Seq uint64

var seq uint64

func Set(s Seq) {
	atomic.CompareAndSwapUint64(&seq, 0, uint64(s))
}

func Next() Seq {
	return Seq(atomic.AddUint64(&seq, 1))
}

func (s Seq) After(o Seq) bool {
	return s > o
}

func (s Seq) Before(o Seq) bool {
	return s < o
}

func (s Seq) Zero() bool {
	return s == 0
}
