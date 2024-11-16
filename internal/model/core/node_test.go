package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNode_SetV(t *testing.T) {
	for _, tc := range []struct {
		name   string
		val    int
		newVal int
	}{
		{
			name:   "new positive value",
			val:    1,
			newVal: 2,
		},
		{
			name:   "new negative value",
			val:    1,
			newVal: -2,
		},
		{
			name:   "new zero value",
			val:    1,
			newVal: 0,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			n := Node[int]{
				v: tc.val,
			}
			newNode := n.SetV(tc.newVal)
			require.Equal(t, tc.newVal, newNode.v)
			require.Equal(t, tc.newVal, n.v)
		})
	}
}

func TestNode_SetLink(t *testing.T) {
	for _, tc := range []struct {
		name    string
		link    *Node[struct{}]
		newLink *Node[struct{}]
	}{
		{
			name:    "old nil link new not nil link",
			link:    nil,
			newLink: &Node[struct{}]{},
		},
		{
			name:    "old not nil link new not nil link",
			link:    &Node[struct{}]{},
			newLink: &Node[struct{}]{},
		},
		{
			name:    "old not nil link new nil link",
			link:    &Node[struct{}]{},
			newLink: &Node[struct{}]{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			n := Node[struct{}]{
				link: tc.link,
			}
			newNode := n.SetLink(tc.newLink)
			require.True(t, tc.newLink == newNode.link)
			require.True(t, tc.newLink == n.link)
		})
	}
}

func TestNode_V(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    *Node[int]
		val  int
	}{
		{
			name: "positive value",
			n: &Node[int]{
				v: 1,
			},
			val: 1,
		},
		{
			name: "zero value",
			n:    &Node[int]{},
			val:  0,
		},
		{
			name: "nil node",
			n:    nil,
			val:  0,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			val := tc.n.V()
			require.Equal(t, tc.val, val)
		})
	}
}

func TestNode_Delete(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    *Node[int]
	}{
		{
			name: "nil node",
			n:    nil,
		},
		{
			name: "empty node",
			n:    &Node[int]{},
		},
		{
			name: "node without prev",
			n: &Node[int]{
				next: &Node[int]{},
			},
		},
		{
			name: "node without next",
			n: &Node[int]{
				prev: &Node[int]{},
			},
		},
		{
			name: "node with links",
			n: &Node[int]{
				next: &Node[int]{},
				prev: &Node[int]{},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.n.Delete()

			if tc.n != nil {
				require.Nil(t, tc.n.next)
				require.Nil(t, tc.n.prev)
			}
		})
	}
}

func TestNode_DeleteLink(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    *Node[int]
		link *Node[int]
	}{
		{
			name: "nil node",
			n:    nil,
			link: nil,
		},
		{
			name: "empty node",
			n:    &Node[int]{},
			link: nil,
		},
		{
			name: "node with link",
			n: &Node[int]{
				link: &Node[int]{
					v: 1,
				},
			},
			link: &Node[int]{
				v: 1,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			link := tc.n.DeleteLink()
			require.Equal(t, tc.link, link)

			if tc.n != nil {
				require.Nil(t, tc.n.link)
			}
		})
	}
}

func TestNode_insert(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    func() (n *Node[int], new *Node[int], requireN func(t *testing.T, n *Node[int]))
	}{
		{
			name: "nil node",
			n: func() (n *Node[int], new *Node[int], requireN func(t *testing.T, n *Node[int])) {
				return nil, &Node[int]{}, func(t *testing.T, n *Node[int]) {
					require.Nil(t, n)
				}
			},
		},
		{
			name: "nil new node",
			n: func() (n *Node[int], new *Node[int], requireN func(t *testing.T, n *Node[int])) {
				return &Node[int]{}, nil, func(t *testing.T, n *Node[int]) {
					require.Nil(t, n.next)
					require.Nil(t, n.prev)
				}
			},
		},
		{
			name: "empty node",
			n: func() (n *Node[int], new *Node[int], requireN func(t *testing.T, n *Node[int])) {
				n = &Node[int]{}
				new = &Node[int]{}

				return n, new, func(t *testing.T, n *Node[int]) {
					require.True(t, n.next == new)
					require.True(t, new.prev == n)
					require.Nil(t, n.prev)
					require.Nil(t, new.next)
				}
			},
		},
		{
			name: "non empty node",
			n: func() (n *Node[int], new *Node[int], requireN func(t *testing.T, n *Node[int])) {
				next := &Node[int]{}
				prev := &Node[int]{}

				n = &Node[int]{
					next: next,
					prev: prev,
				}
				new = &Node[int]{}

				return n, new, func(t *testing.T, n *Node[int]) {
					require.True(t, n.next == new)
					require.True(t, new.prev == n)
					require.True(t, n.prev == prev)
					require.True(t, new.next == next)
				}
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			n, newNode, requireN := tc.n()
			n.insert(newNode)
			requireN(t, n)
		})
	}
}
