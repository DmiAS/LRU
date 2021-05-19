package cache

func (c *Cache) Get(key uint32) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	var val string
	node, val := c.getNodeInfo(key)
	if node != nil {
		c.makeLeastUsed(node)
	}

	return val
}

func (c *Cache) Put(key uint32, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, _ := c.getNodeInfo(key)
	if node != nil {
		c.updateValue(node, val)
	} else {
		c.push(key, val)
	}
}

func (c *Cache) SetBytesCapacity(size uint64) {

}
