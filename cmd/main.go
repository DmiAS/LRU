package main

import (
	"fmt"
	"time"

	"github.com/DmiAS/LRU/internal/cache"
	"github.com/DmiAS/LRU/internal/capacity"
)

func main() {
	//cap := capacity.BytesCap(128)
	cap := capacity.NewElemCap(3)
	c := cache.NewCache(cap, 1) // {}
	c.Put(1, "str1")            // {1: "str1"}
	c.Put(2, "str2")            // {1: "str1", 2: "str2"}
	c.Put(3, "str3")            // {1: "str1", 2: "str2", 3: "str3"}
	for i := 0; i < 40; i++ {
		fmt.Println("time = ", i, c.Get(2))
		fmt.Println("time = ", i, c.Get(1))
		fmt.Println("time = ", i, c.Get(3))
		time.Sleep(time.Second)
	}
}
