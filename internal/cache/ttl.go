package cache

import (
	"time"

	"github.com/DmiAS/LRU/internal/linked_list"
)

// удаление конкретной ноды из кэша
func (c *Cache) delete(node *linked_list.Node) {
	if node == nil {
		return
	}
	key := node.Key()
	c.list.Unlink(node)
	c.hash.Delete(key)
}

func (c *Cache) isExpired(node *linked_list.Node) bool {
	// количество секунд, которое прошло с момента создания записи
	liveSeconds := time.Now().Sub(node.Ttl()).Seconds()
	return node != nil && liveSeconds > c.ttl
}
