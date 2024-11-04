package core

type Node[T any] struct {
	next, prev *Node[T]
	link       *Node[T]

	v T
}

func (n *Node[T]) SetV(val T) *Node[T] {
	n.v = val

	return n
}

func (n *Node[T]) SetLink(link *Node[T]) *Node[T] {
	n.link = link

	return n
}

func (n *Node[T]) V() T {
	if n == nil {
		return *new(T)
	}

	return n.v
}

func (n *Node[T]) Delete() {
	if n == nil {
		return
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	}

	n.next = nil
	n.prev = nil
}

func (n *Node[T]) DeleteLink() *Node[T] {
	if n == nil {
		return nil
	}

	link := n.link
	n.link = nil

	link.Delete()

	return link
}

func (n *Node[T]) insert(new *Node[T]) {
	if n == nil || new == nil {
		return
	}

	new.prev = n
	new.next = n.next

	if n.next != nil {
		n.next.prev = new
	}

	n.next = new
}
