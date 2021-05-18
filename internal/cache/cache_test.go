package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache_Put_One(t *testing.T) {
	c := NewCache(1, 1)
	var key uint32 = 1
	val := "msg"

	c.Put(key, val)
	cacheValue := c.Get(key)
	assert.Equal(t, val, cacheValue)

	time.Sleep(time.Second * 1)
	cacheValue = c.Get(key)
	assert.Equal(t, notFound, cacheValue)
}

func TestCache_Put_Two(t *testing.T) {
	c := NewCache(1, 10)
	var key uint32 = 1
	val := "msg"
	newVal := "newMsg"

	c.Put(key, val)
	cacheValue := c.Get(key)
	assert.Equal(t, val, cacheValue)

	c.Put(key, newVal)
	cacheValue = c.Get(key)
	assert.Equal(t, newVal, cacheValue)
}

func TestCache_Put(t *testing.T) {
	c := NewCache(2, 10)
	var key1, key2 uint32 = 1, 2
	val1, val2 := "str1", "str2"

	c.Put(key1, val1)
	c.Put(key2, val2)

	cacheVal1 := c.Get(key1)
	cacheVal2 := c.Get(key2)
	assert.Equal(t, cacheVal1, val1)
	assert.Equal(t, cacheVal2, val2)
}

func TestCache_Put_Cap(t *testing.T) {
	c := NewCache(3, 10)
	var key1, key2, key3, key4 uint32 = 1, 2, 3, 4
	val1, val2, val3, val4 := "str1", "str2", "str3", "str4"

	c.Put(key1, val1)
	c.Put(key2, val2)
	c.Put(key3, val3)

	cacheVal1 := c.Get(key1)
	cacheVal2 := c.Get(key2)
	cacheVal3 := c.Get(key3)

	assert.Equal(t, cacheVal1, val1)
	assert.Equal(t, cacheVal2, val2)
	assert.Equal(t, cacheVal3, val3)

	_ = c.Get(key1)

	c.Put(key4, val4)

	cacheVal1 = c.Get(key1)
	cacheVal2 = c.Get(key2)
	cacheVal3 = c.Get(key3)
	cacheVal4 := c.Get(key4)

	assert.Equal(t, cacheVal1, val1)
	assert.Equal(t, notFound, cacheVal2)
	assert.Equal(t, cacheVal4, val4)
	assert.Equal(t, cacheVal3, val3)
}

func TestCache_Delete_Nil(t *testing.T) {
	c := NewCache(1, 20)

	assert.NotPanics(t, func() {
		c.Delete(nil)
	})
}

func TestCache_Delete(t *testing.T) {
	c := NewCache(1, 20)
	var key uint32 = 1
	str := "msg1"

	c.Put(key, str)

	node, val := c.getNodeInfo(key)
	assert.NotNil(t, node)
	assert.Equal(t, str, val)

	c.Delete(node)

	val = c.Get(key)

	assert.Equal(t, notFound, val)
}
