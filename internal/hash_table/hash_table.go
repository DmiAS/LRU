package hash_table

import "github.com/DmiAS/LRU/internal/linked_list"

type hash = map[uint32]*linked_list.Node

type Hash struct {
	m hash
}

func NewHash() *Hash {
	return &Hash{m: make(hash)}
}
