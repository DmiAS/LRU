package main

import (
	"fmt"

	"github.com/DmiAS/LRU/internal/cache"
	"github.com/DmiAS/LRU/internal/capacity"
)

func main() {
	//cacheCap := capacity.BytesCap(132) // размер строки 128 байт + размер ключа 4 байта - вместимость 1
	cacheCap := capacity.NewElemCap(3) // кэш размером на 3 элемента
	c := cache.NewCache(cacheCap, 10)  // {}
	c.Put(1, "str1")                   // {1: "str1"}
	c.Put(2, "str2")                   // {1: "str1", 2: "str2"}
	c.Put(3, "str3")                   // {1: "str1", 2: "str2", 3: "str3"}
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	c.Put(4, "str4") // {1: “str1”, 3: “str2”, 4: “str4”}
	fmt.Println("After element displacement:")
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
}
