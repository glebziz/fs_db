package core

type List[T any] struct {
	root Node[T]
}

func (l *List[T]) Clear() {
	l.root.next = &l.root
	l.root.prev = &l.root
}

func (l *List[T]) IsEmpty() bool {
	return l.root.next == nil || l.root.next == &l.root
}

func (l *List[T]) Back() *Node[T] {
	if l.IsEmpty() {
		return nil
	}

	return l.root.prev
}

func (l *List[T]) Front() *Node[T] {
	if l.IsEmpty() {
		return nil
	}

	return l.root.next
}

func (l *List[T]) PushBack(n *Node[T]) {
	if l.root.next == nil {
		l.Clear()
	}

	l.root.prev.insert(n)
}

func (l *List[T]) PopBack() *Node[T] {
	n := l.Back()
	n.Delete()

	return n
}

func (l *List[T]) PopFront() *Node[T] {
	n := l.Front()
	n.Delete()

	return n
}
