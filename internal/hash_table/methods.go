package hash_table

import "github.com/DmiAS/LRU/internal/linked_list"

func (h *Hash) Get(key uint32) *linked_list.Node {
	if node, ok := h.m[key]; !ok {
		return nil
	} else {
		return node
	}
}

func (h *Hash) Set(key uint32, val *linked_list.Node) {
	h.m[key] = val
}

func (h *Hash) Delete(key uint32) {
	delete(h.m, key)
}
