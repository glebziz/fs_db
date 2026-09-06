package model

import (
	"io"
)

type SeekDirection int

const (
	SeekStart SeekDirection = iota
	SeekCurrent
	SeekEnd
)

type Seek struct {
	Pos int64
	Dir SeekDirection
}

func NewSeek(offset int64, whence int) Seek {
	return Seek{
		Pos: offset,
		Dir: newSeekDirection(whence),
	}
}

func (s Seek) Apply(sr io.Seeker) (int64, error) {
	return sr.Seek(s.Pos, s.Dir.whence())
}

func (d SeekDirection) whence() int {
	switch d {
	case SeekStart:
		return io.SeekStart
	case SeekCurrent:
		return io.SeekCurrent
	case SeekEnd:
		return io.SeekEnd
	default:
		return io.SeekStart
	}
}

func newSeekDirection(whence int) SeekDirection {
	switch whence {
	case io.SeekStart:
		return SeekStart
	case io.SeekCurrent:
		return SeekCurrent
	case io.SeekEnd:
		return SeekEnd
	default:
		return SeekStart
	}
}
