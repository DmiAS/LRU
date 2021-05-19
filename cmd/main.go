package main

import (
	"fmt"

	"github.com/DmiAS/LRU/internal/cache"
	"github.com/DmiAS/LRU/internal/capacity"
)

func main() {
	//cacheCap := capacity.BytesCap(132) // размер строки 128 байт + размер ключа 4 байта - вместимость 1
	cacheCap := capacity.NewElemCap(3)
	c := cache.NewCache(cacheCap, 1) // {}
	c.Put(1, "str1")                 // {1: "str1"}
	c.Put(2, "str2")                 // {1: "str1", 2: "str2"}
	c.Put(3, "str3")                 // {1: "str1", 2: "str2", 3: "str3"}
	data := c.Get(1)
	fmt.Println("val = ", data)
}
