package linked_list

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

func (l *LinkedList) Pop() *Node {
	node := l.tail
	if node != nil {
		l.tail = node.prev
	}
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return node
}

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
