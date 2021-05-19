package linked_list

import "time"

func NewNode(key uint32, val string) *Node {
	return &Node{
		value:   val,
		key:     key,
		created: time.Now(), // используется при вычислении испарения
	}
}

func (n *Node) Get() string {
	return n.value
}

func (n *Node) Set(val string) {
	n.value = val
}

func (n *Node) Key() uint32 {
	return n.key
}

func (n *Node) Ttl() time.Time {
	return n.created
}
