package cache

import "github.com/DmiAS/LRU/internal/linked_list"

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
