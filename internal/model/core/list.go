package core

type list[T any] struct {
	root Node[T]
}

func (l *list[T]) isEmpty() bool {
	return l.root.next == nil || l.root.next == &l.root
}

func (l *list[T]) back() *Node[T] {
	if l.isEmpty() {
		return nil
	}

	return l.root.prev
}

func (l *list[T]) front() *Node[T] {
	if l.isEmpty() {
		return nil
	}

	return l.root.next
}

func (l *list[T]) pushBack(n *Node[T]) {
	if l.root.next == nil {
		l.root.next = &l.root
		l.root.prev = &l.root
	}

	l.root.prev.insert(n)
}

func (l *list[T]) popBack() *Node[T] {
	n := l.back()
	n.Delete()

	return n
}

func (l *list[T]) popFront() *Node[T] {
	n := l.front()
	n.Delete()

	return n
}
