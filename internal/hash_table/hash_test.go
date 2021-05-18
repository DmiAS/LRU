package hash_table

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DmiAS/LRU/internal/linked_list"
)

func TestHash_Set(t *testing.T) {
	hash := NewHash()
	var key uint32 = 1
	str := "msg"
	val := linked_list.NewNode(key, str)
	hash.Set(key, val)
	valHash, ok := hash.m[key]
	assert.True(t, ok)
	assert.Equal(t, val, valHash)
}

func TestHash_Get(t *testing.T) {
	hash := NewHash()
	var key uint32 = 1
	str := "msg"
	val := linked_list.NewNode(key, str)
	hash.Set(key, val)
	node := hash.Get(key)
	assert.Equal(t, node, val)
}

func TestHash_Delete(t *testing.T) {
	hash := NewHash()
	var key uint32 = 1
	str := "msg"
	val := linked_list.NewNode(key, str)
	hash.Set(key, val)
	hash.Delete(key)
	node := hash.Get(key)
	assert.Nil(t, node)

	_, ok := hash.m[key]
	assert.False(t, ok)
}
