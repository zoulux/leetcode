package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kthSmallest([][]int{
		{1, 3, 11},
		{2, 4, 6},
	}, 5))

	fmt.Println(17 == kthSmallest([][]int{
		{1, 3, 11},
		{2, 4, 6},
	}, 9))

}

type tuple [3]int // 和，i下标，j下标
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func kSmallestPairs(nums1, nums2 []int, k int) []int {
	n, m := len(nums1), len(nums2)
	ans := make([]int, 0, k) // 预分配空间
	h := hp{{nums1[0] + nums2[0], 0, 0}}
	for len(h) > 0 && len(ans) < k {
		p := heap.Pop(&h).(tuple)
		i, j := p[1], p[2]
		ans = append(ans, p[0]) // 数对和

		//
		// 10 20 30 可以通过 i,j+1 转化 11 21 31
		if j == 0 && i+1 < n {
			heap.Push(&h, tuple{nums1[i+1] + nums2[j], i + 1, j})
		}

		// i j 出堆时 i j+1 入堆 ，00 01 02
		// 其实要入堆的时 i+1 j 和 i j+1，
		if j+1 < m {
			heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return ans
}

//func kSmallestPairs3(nums1, nums2 []int, k int) []int {
//	n, m := len(nums1), len(nums2)
//	ans := make([]int, 0, k) // 预分配空间
//	//ans := make([]int, 0, min(k, n*m)) // 预分配空间
//	h := hp{{nums1[0] + nums2[0], 0, 0}}
//	for len(h) > 0 && len(ans) < k {
//		p := heap.Pop(&h).(tuple)
//		i, j := p.i, p.j
//		ans = append(ans, p.s) // 数对和
//		if j == 0 && i+1 < n {
//			heap.Push(&h, tuple{nums1[i+1] + nums2[0], i + 1, 0})
//		}
//
//		if j+1 < m {
//			heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
//		}
//	}
//	return ans
//}

func kthSmallest(mat [][]int, k int) int {
	kSmallestPairs := func(nums1, nums2 []int) []int {
		n, m := len(nums1), len(nums2)
		ans := make([]int, 0, k) // 预分配空间
		h := hp{{nums1[0] + nums2[0], 0, 0}}
		for len(h) > 0 && len(ans) < k {
			p := heap.Pop(&h).(tuple)
			i, j := p[1], p[2]
			ans = append(ans, p[0]) // 数对和

			//
			// 10 20 30 可以通过 i,j+1 转化 11 21 31
			if j == 0 && i+1 < n {
				heap.Push(&h, tuple{nums1[i+1] + nums2[j], i + 1, j})
			}

			// i j 出堆时 i j+1 入堆 ，00 01 02
			// 其实要入堆的时 i+1 j 和 i j+1，
			if j+1 < m {
				heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
			}
		}
		return ans
	}

	a := []int{0}
	for _, row := range mat {
		a = kSmallestPairs(row, a)
	}
	return a[k-1]
}

func kthSmallest2(mat [][]int, k int) int {
	a := []int{0}
	for _, row := range mat {
		b := make([]int, 0, len(a)*len(row)) // 预分配空间
		for _, x := range a {
			for _, y := range row {
				b = append(b, x+y)
			}
		}
		sort.Ints(b)
		if len(b) > k { // 保留最小的 k 个
			b = b[:k]
		}
		a = b
	}
	return a[k-1]
}
