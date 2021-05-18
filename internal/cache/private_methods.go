package cache

import (
	"github.com/DmiAS/LRU/internal/linked_list"
)

const (
	notFound = "notFound"
)

func (c *Cache) makeLeastUsed(node *linked_list.Node) {
	c.list.Unlink(node)
	c.list.PushNodeFront(node)
}

func (c *Cache) push(key uint32, val string) {
	if c.cap == 0 {
		node := c.list.Pop()
		c.hash.Delete(node.Key())
		c.cap++
	}
	node := c.list.PushFront(key, val)
	c.hash.Set(key, node)
	c.cap--
}

func (c *Cache) getNodeInfo(key uint32) (*linked_list.Node, string) {
	node := c.hash.Get(key)
	if node != nil {
		if c.isExpired(node) {
			return nil, notFound
		}
		return node, node.Get()
	}
	return nil, notFound
}

func (c *Cache) updateValue(node *linked_list.Node, newValue string) {
	node.Set(newValue)
	c.makeLeastUsed(node)
}
