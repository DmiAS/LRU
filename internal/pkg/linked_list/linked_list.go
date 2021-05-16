package linked_list

type Node struct {
	prev  *Node
	next  *Node
	value string
	key   uint32
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}
