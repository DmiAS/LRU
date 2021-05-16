package cache

import (
	"sync"

	"github.com/DmiAS/LRU/internal/hash_table"
	"github.com/DmiAS/LRU/internal/pkg/linked_list"
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
type Cache struct {
	cap  uint16
	hash IHash
	list IList
	ttl  int
	mu   *sync.Mutex
}

func NewCache(cap uint16, ttl int) *Cache {
	return &Cache{
		cap:  cap,
		hash: hash_table.NewHash(),
		list: linked_list.NewLinkedList(),
		ttl:  ttl,
		mu:   &sync.Mutex{},
	}
}
