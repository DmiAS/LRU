package cache

import (
	"time"

	"github.com/DmiAS/LRU/internal/linked_list"
)

func (c *Cache) Delete(node *linked_list.Node) {
	if node == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	key := node.Key()
	c.list.Unlink(node)
	c.hash.Delete(key)
}

func (c *Cache) isExpired(node *linked_list.Node) bool {
	return node != nil && time.Now().Sub(node.Ttl()).Seconds() > c.ttl
}
