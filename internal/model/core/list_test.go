package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList_Clear(t *testing.T) {
	for _, tc := range []struct {
		name string
		l    func() *List[int]
	}{
		{
			name: "uninitiated list",
			l: func() *List[int] {
				return &List[int]{}
			},
		},
		{
			name: "empty list",
			l: func() *List[int] {
				l := List[int]{}
				l.root.next = &l.root
				l.root.prev = &l.root
				return &l
			},
		},
		{
			name: "non empty list",
			l: func() *List[int] {
				l := List[int]{}
				l.PushBack(&Node[int]{})
				return &l
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := tc.l()
			l.Clear()

			require.True(t, l.IsEmpty())
		})
	}
}

func TestList_IsEmpty(t *testing.T) {
	for _, tc := range []struct {
		name    string
		l       func() *List[int]
		isEmpty bool
	}{
		{
			name: "empty List with nil root",
			l: func() *List[int] {
				return &List[int]{}
			},
			isEmpty: true,
		},
		{
			name: "empty List",
			l: func() *List[int] {
				l := &List[int]{}

				l.root.next = &l.root
				l.root.prev = &l.root

				return l
			},
			isEmpty: true,
		},
		{
			name: "non empty List",
			l: func() *List[int] {
				l := &List[int]{}
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

			isEmpty := tc.l().IsEmpty()
			require.Equal(t, tc.isEmpty, isEmpty)
		})
	}
}

func TestList_Back(t *testing.T) {
	for _, tc := range []struct {
		name string
		l    func() (*List[int], *Node[int])
	}{
		{
			name: "empty List with nil root",
			l: func() (*List[int], *Node[int]) {
				return &List[int]{}, nil
			},
		},
		{
			name: "empty List",
			l: func() (*List[int], *Node[int]) {
				l := &List[int]{}

				l.root.next = &l.root
				l.root.prev = &l.root

				return l, nil
			},
		},
		{
			name: "List with one node",
			l: func() (*List[int], *Node[int]) {
				l := &List[int]{}
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
			name: "List with multiple nodes",
			l: func() (*List[int], *Node[int]) {
				l := &List[int]{}
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

			back := l.Back()
			require.Equal(t, n, back)
		})
	}
}

func TestList_Front(t *testing.T) {
	for _, tc := range []struct {
		name string
		l    func() (*List[int], *Node[int])
	}{
		{
			name: "empty List with nil root",
			l: func() (*List[int], *Node[int]) {
				return &List[int]{}, nil
			},
		},
		{
			name: "empty List",
			l: func() (*List[int], *Node[int]) {
				l := &List[int]{}

				l.root.next = &l.root
				l.root.prev = &l.root

				return l, nil
			},
		},
		{
			name: "List with one node",
			l: func() (*List[int], *Node[int]) {
				l := &List[int]{}
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
			name: "List with multiple nodes",
			l: func() (*List[int], *Node[int]) {
				l := &List[int]{}
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

			front := l.Front()
			require.Equal(t, n, front)
		})
	}
}

func TestList_PushBack(t *testing.T) {
	for _, tc := range []struct {
		name     string
		l        func() *List[int]
		n        *Node[int]
		requireL func(t *testing.T, l *List[int])
	}{
		{
			name: "push to empty List",
			l: func() *List[int] {
				return &List[int]{}
			},
			n: &Node[int]{
				v: 1,
			},
			requireL: func(t *testing.T, l *List[int]) {
				require.True(t, l.root.next == l.root.next)
				require.True(t, l.root.next.next == &l.root)
				require.True(t, l.root.next.prev == &l.root)
				require.Equal(t, 1, l.root.next.v)
			},
		},
		{
			name: "push to non empty List",
			l: func() *List[int] {
				l := &List[int]{}
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
			requireL: func(t *testing.T, l *List[int]) {
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
			l.PushBack(tc.n)
			tc.requireL(t, l)
		})
	}
}

func TestList_PopBack(t *testing.T) {
	for _, tc := range []struct {
		name     string
		l        func() *List[int]
		requireL func(t *testing.T, l *List[int])
		n        *Node[int]
	}{
		{
			name: "pop from empty List",
			l: func() *List[int] {
				return &List[int]{}
			},
			requireL: func(t *testing.T, l *List[int]) {
				require.True(t, l.root.next == l.root.prev)
				require.Nil(t, l.root.next)
				require.Nil(t, l.root.prev)
			},
			n: nil,
		},
		{
			name: "pop from List with one node",
			l: func() *List[int] {
				l := &List[int]{}
				n := &Node[int]{
					v:    1,
					next: &l.root,
					prev: &l.root,
				}

				l.root.next = n
				l.root.prev = n

				return l
			},
			requireL: func(t *testing.T, l *List[int]) {
				require.True(t, l.root.next == l.root.prev)
				require.True(t, l.root.next == &l.root)
			},
			n: &Node[int]{
				v: 1,
			},
		},
		{
			name: "pop from List with multiple nodes",
			l: func() *List[int] {
				l := &List[int]{}
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
			requireL: func(t *testing.T, l *List[int]) {
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
			n := l.PopBack()
			tc.requireL(t, l)
			require.Equal(t, tc.n, n)
		})
	}
}

func TestList_PopFront(t *testing.T) {
	for _, tc := range []struct {
		name     string
		l        func() *List[int]
		requireL func(t *testing.T, l *List[int])
		n        *Node[int]
	}{
		{
			name: "pop from empty List",
			l: func() *List[int] {
				return &List[int]{}
			},
			requireL: func(t *testing.T, l *List[int]) {
				require.True(t, l.root.next == l.root.prev)
				require.Nil(t, l.root.next)
				require.Nil(t, l.root.prev)
			},
			n: nil,
		},
		{
			name: "pop from List with one node",
			l: func() *List[int] {
				l := &List[int]{}
				n := &Node[int]{
					v:    1,
					next: &l.root,
					prev: &l.root,
				}

				l.root.next = n
				l.root.prev = n

				return l
			},
			requireL: func(t *testing.T, l *List[int]) {
				require.True(t, l.root.next == l.root.prev)
				require.True(t, l.root.next == &l.root)
			},
			n: &Node[int]{
				v: 1,
			},
		},
		{
			name: "pop from List with multiple nodes",
			l: func() *List[int] {
				l := &List[int]{}
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
			requireL: func(t *testing.T, l *List[int]) {
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
			n := l.PopFront()
			tc.requireL(t, l)
			require.Equal(t, tc.n, n)
		})
	}
}
