package linked_list

func (l *LinkedList) PushFront(key uint32, val string) *Node {
	newNode := newNode(key, val)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
	}
	return newNode
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
