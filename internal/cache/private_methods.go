package cache

import "github.com/DmiAS/LRU/internal/pkg/linked_list"

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
