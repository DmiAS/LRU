package linked_list

// создание нового элемента списка и добавление его в голову списка
func (l *LinkedList) PushFront(key uint32, val string) *Node {
	NewNode := NewNode(key, val)
	if l.head == nil {
		l.head = NewNode
		l.tail = NewNode
	} else {
		l.head.prev = NewNode
		NewNode.next = l.head
		l.head = NewNode
	}
	return NewNode
}

// добавление имеющейся ноды в голову списка
func (l *LinkedList) PushNodeFront(node *Node) {
	if node == nil {
		return
	}

	node.prev = nil
	node.next = l.head
	if l.head != nil {
		l.head.prev = node
	} else {
		l.tail = node
	}
	l.head = node
}

// удаление элемента с хвоста
func (l *LinkedList) Pop() *Node {
	if l.tail == nil {
		return nil
	}

	node := l.tail
	l.tail = node.prev
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return node
}

// удаление элемента из произвольного места в списке
func (l *LinkedList) Unlink(node *Node) {
	if node == nil {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if l.tail == node {
		l.tail = node.prev
	}
	if l.head == node {
		l.head = node.next
	}
	node.prev, node.next = nil, nil
}
