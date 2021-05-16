package cache

func (c *Cache) Get(key uint32) string {
	node := c.hash.Get(key)
	if node == nil {
		return ""
	}
	c.makeLeastUsed(node)
	return node.Get()
}

func (c *Cache) Put(key uint32, val string) {
	if node := c.hash.Get(key); node != nil {
		node.Set(val)
		c.makeLeastUsed(node)
	} else {
		c.push(key, val)
	}
}
