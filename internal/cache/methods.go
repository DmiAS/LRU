package cache

func (c *Cache) Get(key uint32) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	var val string
	node, val := c.getNodeInfo(key)
	// если узел не nil, то нужно переставить его в голову списка
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
		// если узел с таким ключом существует, то нужно обновить значение в нем и
		// переставить его в голову списка
		c.updateValue(node, val)
	} else {
		// в противном случае нужно создавать новый узел, добавить его в мапу и список
		c.push(key, val)
	}
}
