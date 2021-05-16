package main

import (
	"fmt"
	"time"

	"github.com/DmiAS/LRU/internal/cache"
)

func main() {
	c := cache.NewCache(3, 30) // {}
	c.Put(1, "str1")           // {1: "str1"}
	c.Put(2, "str2")           // {1: "str1", 2: "str2"}
	c.Put(3, "str3")           // {1: "str1", 2: "str2", 3: "str3"}
	for i := 0; i < 40; i++ {
		fmt.Println("time = ", i, c.Get(3))
		time.Sleep(time.Second)
	}
}
