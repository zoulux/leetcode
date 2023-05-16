package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {

	fmt.Println(findKthLargest([]int{
		3, 2, 1, 5, 6, 4,
	}, 2))

	fmt.Println(findKthLargest([]int{
		3, 2, 3, 1, 2, 4, 5, 5, 6,
	}, 4))
}

type kheap struct {
	sort.IntSlice
}

func (h *kheap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *kheap) Pop() interface{} {

	ans := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return ans
}

//Push(x any) // add x as element Len()
//Pop() any   // remove and return element Len() - 1.

func findKthLargest(nums []int, k int) int {
	kh := &kheap{}
	for _, v := range nums {
		heap.Push(kh, v)
		if len(kh.IntSlice) > k {
			heap.Pop(kh)
		}
	}
	return kh.IntSlice[0]
}
