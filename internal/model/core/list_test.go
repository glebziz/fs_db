package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList_isEmpty(t *testing.T) {
	for _, tc := range []struct {
		name    string
		l       func() *list[int]
		isEmpty bool
	}{
		{
			name: "empty list with nil root",
			l: func() *list[int] {
				return &list[int]{}
			},
			isEmpty: true,
		},
		{
			name: "empty list",
			l: func() *list[int] {
				l := &list[int]{}

				l.root.next = &l.root
				l.root.prev = &l.root

				return l
			},
			isEmpty: true,
		},
		{
			name: "non empty list",
			l: func() *list[int] {
				l := &list[int]{}
				n := &Node[int]{
					next: &l.root,
					prev: &l.root,
				}

				l.root.next = n
				l.root.prev = n

				return l
			},
			isEmpty: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			isEmpty := tc.l().isEmpty()
			require.Equal(t, tc.isEmpty, isEmpty)
		})
	}
}

func TestList_back(t *testing.T) {
	for _, tc := range []struct {
		name string
		l    func() (*list[int], *Node[int])
	}{
		{
			name: "empty list with nil root",
			l: func() (*list[int], *Node[int]) {
				return &list[int]{}, nil
			},
		},
		{
			name: "empty list",
			l: func() (*list[int], *Node[int]) {
				l := &list[int]{}

				l.root.next = &l.root
				l.root.prev = &l.root

				return l, nil
			},
		},
		{
			name: "list with one node",
			l: func() (*list[int], *Node[int]) {
				l := &list[int]{}
				n := &Node[int]{
					next: &l.root,
					prev: &l.root,
				}

				l.root.next = n
				l.root.prev = n

				return l, n
			},
		},
		{
			name: "list with multiple nodes",
			l: func() (*list[int], *Node[int]) {
				l := &list[int]{}
				n1 := &Node[int]{}
				n2 := &Node[int]{}

				l.root.next = n1
				n1.prev = &l.root
				n1.next = n2
				n2.prev = n1
				n2.next = &l.root
				l.root.prev = n2

				return l, n2
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l, n := tc.l()

			back := l.back()
			require.Equal(t, n, back)
		})
	}
}

func TestList_front(t *testing.T) {
	for _, tc := range []struct {
		name string
		l    func() (*list[int], *Node[int])
	}{
		{
			name: "empty list with nil root",
			l: func() (*list[int], *Node[int]) {
				return &list[int]{}, nil
			},
		},
		{
			name: "empty list",
			l: func() (*list[int], *Node[int]) {
				l := &list[int]{}

				l.root.next = &l.root
				l.root.prev = &l.root

				return l, nil
			},
		},
		{
			name: "list with one node",
			l: func() (*list[int], *Node[int]) {
				l := &list[int]{}
				n := &Node[int]{
					next: &l.root,
					prev: &l.root,
				}

				l.root.next = n
				l.root.prev = n

				return l, n
			},
		},
		{
			name: "list with multiple nodes",
			l: func() (*list[int], *Node[int]) {
				l := &list[int]{}
				n1 := &Node[int]{}
				n2 := &Node[int]{}

				l.root.next = n1
				n1.prev = &l.root
				n1.next = n2
				n2.prev = n1
				n2.next = &l.root
				l.root.prev = n2

				return l, n1
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l, n := tc.l()

			front := l.front()
			require.Equal(t, n, front)
		})
	}
}

func TestList_pushBack(t *testing.T) {
	for _, tc := range []struct {
		name     string
		l        func() *list[int]
		n        *Node[int]
		requireL func(t *testing.T, l *list[int])
	}{
		{
			name: "push to empty list",
			l: func() *list[int] {
				return &list[int]{}
			},
			n: &Node[int]{
				v: 1,
			},
			requireL: func(t *testing.T, l *list[int]) {
				require.True(t, l.root.next == l.root.next)
				require.True(t, l.root.next.next == &l.root)
				require.True(t, l.root.next.prev == &l.root)
				require.Equal(t, 1, l.root.next.v)
			},
		},
		{
			name: "push to non empty list",
			l: func() *list[int] {
				l := &list[int]{}
				n := &Node[int]{
					v:    1,
					next: &l.root,
					prev: &l.root,
				}

				l.root.next = n
				l.root.prev = n

				return l
			},
			n: &Node[int]{
				v: 2,
			},
			requireL: func(t *testing.T, l *list[int]) {
				require.True(t, l.root.next != l.root.prev)
				require.True(t, l.root.next.prev == &l.root)
				require.True(t, l.root.prev.next == &l.root)
				require.Equal(t, l.root.prev.v, 2)
				require.Equal(t, l.root.next.v, 1)
				require.Equal(t, l.root.prev.prev.v, 1)
				require.Equal(t, l.root.next.next.v, 2)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			l.pushBack(tc.n)
			tc.requireL(t, l)
		})
	}
}

func TestList_popBack(t *testing.T) {
	for _, tc := range []struct {
		name     string
		l        func() *list[int]
		requireL func(t *testing.T, l *list[int])
		n        *Node[int]
	}{
		{
			name: "pop from empty list",
			l: func() *list[int] {
				return &list[int]{}
			},
			requireL: func(t *testing.T, l *list[int]) {
				require.True(t, l.root.next == l.root.prev)
				require.Nil(t, l.root.next)
				require.Nil(t, l.root.prev)
			},
			n: nil,
		},
		{
			name: "pop from list with one node",
			l: func() *list[int] {
				l := &list[int]{}
				n := &Node[int]{
					v:    1,
					next: &l.root,
					prev: &l.root,
				}

				l.root.next = n
				l.root.prev = n

				return l
			},
			requireL: func(t *testing.T, l *list[int]) {
				require.True(t, l.root.next == l.root.prev)
				require.True(t, l.root.next == &l.root)
			},
			n: &Node[int]{
				v: 1,
			},
		},
		{
			name: "pop from list with multiple nodes",
			l: func() *list[int] {
				l := &list[int]{}
				n1 := &Node[int]{
					v: 1,
				}
				n2 := &Node[int]{
					v: 2,
				}

				l.root.next = n1
				n1.prev = &l.root
				n1.next = n2
				n2.prev = n1
				n2.next = &l.root
				l.root.prev = n2

				return l
			},
			requireL: func(t *testing.T, l *list[int]) {
				require.True(t, l.root.next == l.root.prev)
				require.True(t, l.root.next.next == &l.root)
				require.True(t, l.root.prev.prev == &l.root)
				require.Equal(t, 1, l.root.prev.v)
			},
			n: &Node[int]{
				v: 2,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			n := l.popBack()
			tc.requireL(t, l)
			require.Equal(t, tc.n, n)
		})
	}
}

func TestList_popFront(t *testing.T) {
	for _, tc := range []struct {
		name     string
		l        func() *list[int]
		requireL func(t *testing.T, l *list[int])
		n        *Node[int]
	}{
		{
			name: "pop from empty list",
			l: func() *list[int] {
				return &list[int]{}
			},
			requireL: func(t *testing.T, l *list[int]) {
				require.True(t, l.root.next == l.root.prev)
				require.Nil(t, l.root.next)
				require.Nil(t, l.root.prev)
			},
			n: nil,
		},
		{
			name: "pop from list with one node",
			l: func() *list[int] {
				l := &list[int]{}
				n := &Node[int]{
					v:    1,
					next: &l.root,
					prev: &l.root,
				}

				l.root.next = n
				l.root.prev = n

				return l
			},
			requireL: func(t *testing.T, l *list[int]) {
				require.True(t, l.root.next == l.root.prev)
				require.True(t, l.root.next == &l.root)
			},
			n: &Node[int]{
				v: 1,
			},
		},
		{
			name: "pop from list with multiple nodes",
			l: func() *list[int] {
				l := &list[int]{}
				n1 := &Node[int]{
					v: 1,
				}
				n2 := &Node[int]{
					v: 2,
				}

				l.root.next = n1
				n1.prev = &l.root
				n1.next = n2
				n2.prev = n1
				n2.next = &l.root
				l.root.prev = n2

				return l
			},
			requireL: func(t *testing.T, l *list[int]) {
				require.True(t, l.root.next == l.root.prev)
				require.True(t, l.root.next.next == &l.root)
				require.True(t, l.root.prev.prev == &l.root)
				require.Equal(t, 2, l.root.prev.v)
			},
			n: &Node[int]{
				v: 1,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			n := l.popFront()
			tc.requireL(t, l)
			require.Equal(t, tc.n, n)
		})
	}
}
