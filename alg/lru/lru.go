package main

import (
	"fmt"
	"leetcode/alg/lru/byslices"
)

func main() {

	lru := byslices.Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 2)
	lru.Put(2, 2)
	lru.Put(2, 2)

	fmt.Println(lru.Get(1))

}
