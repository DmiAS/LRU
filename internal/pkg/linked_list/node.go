package linked_list

import "time"

func newNode(key uint32, val string) *Node {
	return &Node{
		value:   val,
		key:     key,
		created: time.Now(),
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
