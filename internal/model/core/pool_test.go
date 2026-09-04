package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPool_Acquire(t *testing.T) {
	for _, tc := range []struct {
		name string
		p    *Pool[int]
		e    *int
	}{
		{
			name: "success",
			p: &Pool[int]{
				free: []*int{new(1)},
			},
			e: new(1),
		},
		{
			name: "success with empty pool",
			p: &Pool[int]{
				free: nil,
			},
			e: new(0),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			e := tc.p.Acquire()
			require.Equal(t, tc.e, e)
		})
	}
}

func TestPool_Release(t *testing.T) {
	for _, tc := range []struct {
		name      string
		p         *Pool[int]
		freeElems []*int
	}{
		{
			name: "success",
			p: &Pool[int]{
				free: []*int{new(3)},
			},
			freeElems: []*int{new(3), new(0), new(0)},
		},
		{
			name: "success with empty pool",
			p: &Pool[int]{
				free: nil,
			},
			freeElems: []*int{new(0), new(0)},
		},
		{
			name: "success with clearFunc",
			p: NewPool(func(e *int) {
				*e = 10
			}),
			freeElems: []*int{new(10), new(10)},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.p.Release(new(1), new(2))
			require.Equal(t, tc.freeElems, tc.p.free)
		})
	}
}
