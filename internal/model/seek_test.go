package model

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type seeker struct {
	t *testing.T

	offset int64
	whence int

	n   int64
	err error
}

func (s seeker) Seek(offset int64, whence int) (int64, error) {
	require.Equal(s.t, s.offset, offset)
	require.Equal(s.t, s.whence, whence)

	return s.n, s.err
}

func TestNewSeek(t *testing.T) {
	t.Parallel()

	const (
		offset = iota + 1
	)

	for _, tc := range []struct {
		name   string
		offset int64
		whence int
		seek   Seek
	}{
		{
			name:   "success",
			offset: 1,
			whence: io.SeekCurrent,
			seek: Seek{
				Pos: offset,
				Dir: SeekCurrent,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			seek := NewSeek(tc.offset, tc.whence)
			require.Equal(t, tc.seek, seek)
		})
	}
}

func TestSeek_Apply(t *testing.T) {
	t.Parallel()

	const (
		pos int64 = iota + 1
		offset
	)

	for _, tc := range []struct {
		name   string
		seeker seeker
		seek   Seek
	}{
		{
			name: "success",
			seeker: seeker{
				offset: offset,
				whence: io.SeekStart,
				n:      pos,
			},
			seek: Seek{
				Pos: offset,
				Dir: SeekStart,
			},
		},
		{
			name: "seeker error",
			seeker: seeker{
				offset: offset,
				whence: io.SeekStart,
				err:    assert.AnError,
			},
			seek: Seek{
				Pos: offset,
				Dir: SeekStart,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.seeker.t = t
			n, err := tc.seek.Apply(tc.seeker)
			require.ErrorIs(t, tc.seeker.err, err)
			require.Equal(t, tc.seeker.n, n)
		})
	}
}

func TestSeekDirection_whence(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		d      SeekDirection
		whence int
	}{
		{
			name:   "SeekStart",
			d:      SeekStart,
			whence: io.SeekStart,
		},
		{
			name:   "SeekCurrent",
			d:      SeekCurrent,
			whence: io.SeekCurrent,
		},
		{
			name:   "SeekEnd",
			d:      SeekEnd,
			whence: io.SeekEnd,
		},
		{
			name:   "default",
			d:      -1,
			whence: io.SeekStart,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			whence := tc.d.whence()
			require.Equal(t, tc.whence, whence)
		})
	}
}

func Test_newSeekDirection(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		whence int
		dir    SeekDirection
	}{
		{
			name:   "SeekStart",
			whence: io.SeekStart,
			dir:    SeekStart,
		},
		{
			name:   "SeekCurrent",
			whence: io.SeekCurrent,
			dir:    SeekCurrent,
		},
		{
			name:   "SeekEnd",
			whence: io.SeekEnd,
			dir:    SeekEnd,
		},
		{
			name:   "default",
			whence: -1,
			dir:    SeekStart,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			dir := newSeekDirection(tc.whence)
			require.Equal(t, tc.dir, dir)
		})
	}
}
