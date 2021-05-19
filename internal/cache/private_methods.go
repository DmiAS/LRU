package cache

import (
	"github.com/DmiAS/LRU/internal/linked_list"
)

const (
	notFound = "not found"
)

func (c *Cache) makeLeastUsed(node *linked_list.Node) {
	// вытаскиваем из списка
	c.list.Unlink(node)
	// вставляем в голову списка
	c.list.PushNodeFront(node)
}

func (c *Cache) push(key uint32, val string) {
	// если кэш полон, то удаляю элемент с хвоста списка(наименее используемый элемент в кэше) и
	// из мапы удаляю элемент, лежащий по текущему ключу
	if c.available == 0 {
		node := c.list.Pop()
		if node == nil {
			return
		}
		c.hash.Delete(node.Key())
		c.available++
	}
	node := c.list.PushFront(key, val)
	c.hash.Set(key, node)
	c.available--
}

func (c *Cache) getNodeInfo(key uint32) (*linked_list.Node, string) {
	node := c.hash.Get(key)
	if node == nil {
		return nil, notFound
	}
	// проверяю испарился ли ttl, если да, то нужно удалить элемент из кэша
	if c.isExpired(node) {
		c.delete(node)
		return nil, notFound
	}
	return node, node.Get()
}

func (c *Cache) updateValue(node *linked_list.Node, newValue string) {
	node.Set(newValue)
	c.makeLeastUsed(node)
}
