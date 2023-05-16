package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {

	fmt.Println(maxScore([]int{
		1, 3, 3, 2,
	}, []int{
		2, 1, 3, 4,
	}, 3))

}

func maxScore(nums1, nums2 []int, k int) int64 {
	type pair [2]int
	a := make([]pair, len(nums1))
	sum := 0
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i][1] > a[j][1] })

	h := hp{nums2[:k]} // 复用内存
	for i, p := range a[:k] {
		sum += p[0]
		h.IntSlice[i] = p[0]
	}
	ans := sum * a[k-1][1]
	heap.Init(&h)
	for _, p := range a[k:] {
		if p[0] > h.IntSlice[0] {
			sum += p[0] - h.replace(p[0])
			ans = max(ans, sum*p[1])
		}
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (hp) Pop() (_ interface{}) { return }
func (hp) Push(interface{})     {}
func (h hp) replace(v int) int {
	top := h.IntSlice[0]
	h.IntSlice[0] = v
	heap.Fix(&h, 0)
	return top
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
