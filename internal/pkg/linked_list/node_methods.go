package linked_list

func (n *Node) Get() string {
	return n.value
}

func (n *Node) Set(val string) {
	n.value = val
}

func (n *Node) Key() uint32 {
	return n.key
}
