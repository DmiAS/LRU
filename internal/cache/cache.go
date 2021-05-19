package cache

import (
	"sync"

	"github.com/DmiAS/LRU/internal/hash_table"
	"github.com/DmiAS/LRU/internal/linked_list"
)

type IHash interface {
	Get(key uint32) *linked_list.Node
	Set(key uint32, val *linked_list.Node)
	Delete(key uint32)
}

type IList interface {
	Unlink(node *linked_list.Node)
	PushNodeFront(node *linked_list.Node)
	PushFront(key uint32, val string) *linked_list.Node
	Pop() *linked_list.Node
}

type ICap interface {
	GetSize() uint16
}

type Cache struct {
	cap  uint16
	hash IHash
	list IList
	ttl  float64
	mu   *sync.Mutex
}

// т.к. размер может задаваться либо количеством элементов,
// либо размером потребляемой памяти, то я использовал типы, которые
// реализуют интерфейс с методом, возвращающим capacity = количеству возможных элементов в кэше
func NewCache(cap ICap, ttl float64) *Cache {
	return &Cache{
		cap:  cap.GetSize(),
		hash: hash_table.NewHash(),
		list: linked_list.NewLinkedList(),
		ttl:  ttl,
		mu:   &sync.Mutex{},
	}
}
