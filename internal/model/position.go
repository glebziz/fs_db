package model

type Position struct {
	current int64
	end     int64
}

func (p *Position) Seek(seek Seek) error {
	newPos := p.current
	switch seek.Dir {
	case SeekStart:
		newPos = seek.Pos
	case SeekCurrent:
		newPos += seek.Pos
	case SeekEnd:
		newPos = p.end + seek.Pos
	}

	if newPos < 0 {
		return ErrInvalidPosition
	}

	p.current = newPos
	p.end = max(p.end, p.current)

	return nil
}

func (p *Position) SetEnd(end int64) {
	p.end = max(p.end, end)
}

func (p *Position) Pos() int64 {
	return p.current
}

func (p *Position) NeedSeek(seek Seek) bool {
	switch seek.Dir {
	case SeekStart:
		return seek.Pos != p.current
	case SeekCurrent:
		return seek.Pos != 0
	case SeekEnd:
		return seek.Pos != (p.current - p.end)
	}

	return true
}
