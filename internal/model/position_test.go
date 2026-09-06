package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPosition_Seek(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
		pos  Position
		seek Seek
		res  Position
		err  error
	}{
		{
			name: "positive SeekStart",
			pos: Position{
				current: 200,
				end:     200,
			},
			seek: Seek{
				Pos: 100,
				Dir: SeekStart,
			},
			res: Position{
				current: 100,
				end:     200,
			},
		},
		{
			name: "negative SeekStart",
			pos: Position{
				current: 100,
				end:     100,
			},
			seek: Seek{
				Pos: -100,
				Dir: SeekStart,
			},
			res: Position{
				current: 100,
				end:     100,
			},
			err: ErrInvalidPosition,
		},
		{
			name: "zero SeekStart",
			pos: Position{
				current: 100,
				end:     200,
			},
			seek: Seek{
				Pos: 0,
				Dir: SeekStart,
			},
			res: Position{
				current: 0,
				end:     200,
			},
		},
		{
			name: "positive SeekEnd",
			pos: Position{
				current: 0,
				end:     100,
			},
			seek: Seek{
				Pos: 100,
				Dir: SeekEnd,
			},
			res: Position{
				current: 200,
				end:     200,
			},
		},
		{
			name: "negative SeekEnd",
			pos: Position{
				current: 50,
				end:     100,
			},
			seek: Seek{
				Pos: -100,
				Dir: SeekEnd,
			},
			res: Position{
				current: 0,
				end:     100,
			},
		},
		{
			name: "negative SeekEnd with error",
			pos: Position{
				current: 0,
				end:     50,
			},
			seek: Seek{
				Pos: -100,
				Dir: SeekEnd,
			},
			res: Position{
				current: 0,
				end:     50,
			},
			err: ErrInvalidPosition,
		},
		{
			name: "zero SeekEnd",
			pos: Position{
				current: 0,
				end:     100,
			},
			seek: Seek{
				Pos: 0,
				Dir: SeekEnd,
			},
			res: Position{
				current: 100,
				end:     100,
			},
		},
		{
			name: "positive SeekCurrent",
			pos: Position{
				current: 50,
				end:     100,
			},
			seek: Seek{
				Pos: 100,
				Dir: SeekCurrent,
			},
			res: Position{
				current: 150,
				end:     150,
			},
		},
		{
			name: "negative SeekCurrent",
			pos: Position{
				current: 150,
				end:     200,
			},
			seek: Seek{
				Pos: -100,
				Dir: SeekCurrent,
			},
			res: Position{
				current: 50,
				end:     200,
			},
		},
		{
			name: "zero SeekCurrent",
			pos: Position{
				current: 100,
				end:     200,
			},
			seek: Seek{
				Pos: 0,
				Dir: SeekCurrent,
			},
			res: Position{
				current: 100,
				end:     200,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := tc.pos.Seek(tc.seek)

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, tc.res, tc.pos)
		})
	}
}

func TestPosition_SetEnd(t *testing.T) {
	t.Parallel()

	const (
		end int64 = 10
	)

	for _, tc := range []struct {
		name   string
		pos    Position
		resEnd int64
	}{
		{
			name: "end less than newEnd",
			pos: Position{
				end: end - 1,
			},
			resEnd: end,
		},
		{
			name: "end equals newEnd",
			pos: Position{
				end: end,
			},
			resEnd: end,
		},
		{
			name: "end greater than newEnd",
			pos: Position{
				end: end + 1,
			},
			resEnd: end + 1,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.pos.SetEnd(end)

			require.Equal(t, tc.resEnd, tc.pos.end)
		})
	}
}

func TestPosition_Pos(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
		pos  Position
		p    int64
	}{
		{
			name: "non-zero position",
			pos: Position{
				current: 200,
				end:     200,
			},
			p: 200,
		},
		{
			name: "zero position",
			pos: Position{
				current: 0,
				end:     200,
			},
			p: 0,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := tc.pos.Pos()
			require.Equal(t, tc.p, p)
		})
	}
}

func TestPosition_NeedSeek(t *testing.T) {
	t.Parallel()

	const (
		current int64 = 1 << (iota + 1)
		end
	)

	for _, tc := range []struct {
		name     string
		seek     Seek
		needSeek bool
	}{
		{
			name: "need seek for start direction",
			seek: Seek{
				Pos: end + 1,
				Dir: SeekStart,
			},
			needSeek: true,
		},
		{
			name: "not need seek for start direction",
			seek: Seek{
				Pos: current,
				Dir: SeekStart,
			},
			needSeek: false,
		},
		{
			name: "need seek for current direction",
			seek: Seek{
				Pos: current,
				Dir: SeekCurrent,
			},
			needSeek: true,
		},
		{
			name: "not need seek for current direction",
			seek: Seek{
				Pos: 0,
				Dir: SeekCurrent,
			},
			needSeek: false,
		},
		{
			name: "need seek for end direction",
			seek: Seek{
				Pos: 1,
				Dir: SeekEnd,
			},
			needSeek: true,
		},
		{
			name: "not need seek for end direction",
			seek: Seek{
				Pos: current - end,
				Dir: SeekEnd,
			},
			needSeek: false,
		},
		{
			name: "unknown direction",
			seek: Seek{
				Pos: current,
				Dir: -1,
			},
			needSeek: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := Position{
				current: current,
				end:     end,
			}

			needSeek := p.NeedSeek(tc.seek)
			require.Equal(t, tc.needSeek, needSeek)
		})
	}
}
