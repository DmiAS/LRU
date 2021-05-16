package cache

func (c *Cache) Delete(key uint32) func() {
	return func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		node := c.hash.Get(key)
		if node == nil {
			return
		}
		c.list.Unlink(node)
		c.hash.Delete(key)
	}
}
