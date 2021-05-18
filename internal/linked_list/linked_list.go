package linked_list

import "time"

type Node struct {
	prev    *Node
	next    *Node
	value   string
	key     uint32
	created time.Time
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}
