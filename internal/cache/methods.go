package cache

import "time"

func (c *Cache) Get(key uint32) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	node := c.hash.Get(key)
	if node == nil {
		return ""
	}
	c.makeLeastUsed(node)
	return node.Get()
}

func (c *Cache) Put(key uint32, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if node := c.hash.Get(key); node != nil {
		node.Set(val)
		c.makeLeastUsed(node)
	} else {
		c.push(key, val)
	}

	time.AfterFunc(time.Duration(c.ttl)*time.Second, c.Delete(key))
}
