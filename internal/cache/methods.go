package cache

func (c *Cache) Get(key uint32) string {
	var val string
	node, val := c.getNodeInfo(key)

	c.mu.Lock()
	defer c.mu.Unlock()

	if node != nil {
		c.makeLeastUsed(node)
	}

	return val
}

func (c *Cache) Put(key uint32, val string) {
	node, _ := c.getNodeInfo(key)
	c.mu.Lock()
	defer c.mu.Unlock()
	if node != nil {
		c.updateValue(node, val)
	} else {
		c.push(key, val)
	}
}
