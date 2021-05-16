package main

import (
	"fmt"

	"github.com/DmiAS/LRU/internal/cache"
)

func main() {
	c := cache.NewCache(3) // {}
	c.Put(1, "str1")       // {1: "str1"}
	c.Put(2, "str2")       // {1: "str1", 2: "str2"}
	c.Put(3, "str3")       // {1: "str1", 2: "str2", 3: "str3"}
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	c.Put(4, "str4") // {1: "str1", 3: "str2", 4: "str4"}
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
}
