package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	h := &hp{}
	heap.Push(h, 7)
	heap.Push(h, 5)
	heap.Push(h, 3)
	heap.Push(h, 10)
	fmt.Println(h.IntSlice[0])
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	ln := len(h.IntSlice) - 1
	v := h.IntSlice[ln]
	h.IntSlice = h.IntSlice[:ln]
	return v
}

//Push(x any) // add x as element Len()
//Pop() any   // remove and return element Len() - 1.
