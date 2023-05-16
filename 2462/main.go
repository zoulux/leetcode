package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {

	fmt.Println(totalCost2([]int{
		17, 12, 10, 2, 7, 2, 11, 20, 8,
	}, 3, 4))
	fmt.Println(totalCost2([]int{
		31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58,
	}, 11, 2))
}

type miniHeap struct {
	sort.IntSlice
}

func (h *miniHeap) replace(v int) int {
	top := h.IntSlice[0]
	h.IntSlice[0] = v
	heap.Fix(h, 0)
	return top
}

func (h *miniHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *miniHeap) Pop() interface{} {
	x := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	vis := make([]bool, len(costs))

	hleft := &miniHeap{}
	hright := &miniHeap{}

	lidx, ridx := 0, len(costs)-1
	for ; lidx < candidates; lidx++ {
		vis[lidx] = true
		hleft.Push(costs[lidx])
	}

	for ; ridx >= len(costs)-candidates; ridx-- {
		if !vis[ridx] {
			vis[ridx] = true
			hright.Push(costs[ridx])
		}
	}

	heap.Init(hleft)
	heap.Init(hright)

	var ans int
	for k != 0 {
		if len(hleft.IntSlice) != 0 && len(hright.IntSlice) != 0 {
			// 两个都不为空则比较
			if hleft.IntSlice[0] <= hright.IntSlice[0] {
				if lidx < len(costs) && !vis[lidx] {
					ans += hleft.replace(costs[lidx])
					vis[lidx] = true
					lidx++
				} else {
					ans += hleft.IntSlice[0]
					heap.Pop(hleft)
				}
			} else {
				if ridx >= 0 && !vis[ridx] {
					ans += hright.replace(costs[ridx])
					vis[ridx] = true
					ridx--
				} else {
					ans += hright.IntSlice[0]
					heap.Pop(hright)
				}
			}
		} else if len(hleft.IntSlice) != 0 {
			// 左边不为空
			if lidx < len(costs) && !vis[lidx] {
				ans += hleft.replace(costs[lidx])
				vis[lidx] = true
				lidx++
			} else {
				ans += hleft.IntSlice[0]
				heap.Pop(hleft)
			}
		} else {
			// 右边不为空
			if ridx >= 0 && !vis[ridx] {
				ans += hright.replace(costs[ridx])
				vis[ridx] = true
				ridx--
			} else {
				ans += hright.IntSlice[0]
				heap.Pop(hright)
			}
		}

		k--
	}

	return int64(ans)
}

func totalCost2(costs []int, k int, candidates int) int64 {

	hleft := &miniHeap{}  // 左边的堆
	hright := &miniHeap{} // 右边的堆

	lidx, ridx := 0, len(costs)-1 // 左边的索引，右边的索引¬
	for ; lidx < candidates; lidx++ {
		hleft.Push(costs[lidx])
	}

	for ; ridx >= len(costs)-candidates && lidx < ridx; ridx-- {
		hright.Push(costs[ridx])
	}

	heap.Init(hleft)
	heap.Init(hright)

	var ans int
	for k != 0 {
		if len(hleft.IntSlice) != 0 && len(hright.IntSlice) != 0 {
			// 两个都不为空则比较
			if hleft.IntSlice[0] <= hright.IntSlice[0] {
				ans += hleft.IntSlice[0]
				heap.Pop(hleft)
				if lidx < len(costs) && lidx < ridx {
					heap.Push(hleft, costs[lidx])
					lidx++
				}
			} else {
				ans += hright.IntSlice[0]
				heap.Pop(hright)

				if ridx >= 0 && lidx < ridx {
					heap.Push(hright, costs[ridx])
					ridx--
				}

			}
		} else if len(hleft.IntSlice) != 0 {
			// 左边不为空
			ans += hleft.IntSlice[0]
			heap.Pop(hleft)

			if lidx < len(costs) && lidx < ridx {
				heap.Push(hleft, costs[lidx])
				lidx++
			}
		} else {
			// 右边不为空
			ans += hright.IntSlice[0]
			heap.Pop(hright)

			if ridx >= 0 && lidx < ridx {
				heap.Push(hright, costs[ridx])
				ridx--
			}
		}
		k--
	}
	return int64(ans)
}
