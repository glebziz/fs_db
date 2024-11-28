package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/utils/ptr"
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
				free: []*int{ptr.Ptr(1)},
			},
			e: ptr.Ptr(1),
		},
		{
			name: "success with empty pool",
			p: &Pool[int]{
				free: nil,
			},
			e: ptr.Ptr(0),
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
				free: []*int{ptr.Ptr(3)},
			},
			freeElems: []*int{ptr.Ptr(3), ptr.Ptr(0), ptr.Ptr(0)},
		},
		{
			name: "success with empty pool",
			p: &Pool[int]{
				free: nil,
			},
			freeElems: []*int{ptr.Ptr(0), ptr.Ptr(0)},
		},
		{
			name: "success with clearFunc",
			p: NewPool(func(e *int) {
				*e = 10
			}),
			freeElems: []*int{ptr.Ptr(10), ptr.Ptr(10)},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.p.Release(ptr.Ptr(1), ptr.Ptr(2))
			require.Equal(t, tc.freeElems, tc.p.free)
		})
	}
}
