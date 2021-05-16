package linked_list

import "fmt"

func (l *LinkedList) PushFront(key uint32, val string) *Node {
	newNode := &Node{value: val, key: key}
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
	node.prev = nil
	if l.head != nil && l.head != node {
		node.next = l.head
		l.head.prev = node
	} else {
		node.next = nil
	}
	l.head = node
}

func (l *LinkedList) Pop() *Node {
	node := l.tail
	if l.tail != nil && l.tail.prev != nil {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	} else {
		l.head, l.tail = nil, nil
	}

	return node
}

func (l LinkedList) PrintList() {
	for curr := l.head; curr != nil; curr = curr.next {
		fmt.Println("val = ", curr.Get())
	}
}

func (l *LinkedList) Unlink(node *Node) {
	if l.head == node {
		l.head, l.head = nil, nil
	} else {
		if l.tail == node {
			l.tail = node.prev
		}
		node.prev.next = node.next
	}
	node.prev, node.next = nil, nil
}
