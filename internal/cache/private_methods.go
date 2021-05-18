package cache

import (
	"time"

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

func (c *Cache) isExpired(node *linked_list.Node) bool {
	return node != nil && time.Now().Sub(node.Ttl()).Seconds() > c.ttl
}

func (c *Cache) getNodeInfo(key uint32) (*linked_list.Node, string) {
	var val string
	c.mu.RLock()

	node := c.hash.Get(key)
	shouldDelete := c.isExpired(node)
	defer func() {
		if shouldDelete {
			c.Delete(node)
			node = nil
		}
	}()

	if node != nil {
		val = node.Get()
	} else {
		val = notFound
	}
	c.mu.RUnlock()
	return node, val
}

func (c *Cache) updateValue(node *linked_list.Node, newValue string) {
	node.Set(newValue)
	c.makeLeastUsed(node)
}
