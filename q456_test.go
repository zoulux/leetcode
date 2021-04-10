package leetcode

import (
	"fmt"
	"testing"
)

func find132pattern(nums []int) bool {
	n := len(nums)

	tmp := make([][]bool, n)
	for i := range tmp {
		tmp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				tmp[i][j] = true
			}

			if nums[j] < nums[i] {
				tmp[j][i] = true
			}
		}
	}

	fmt.Println(tmp)

	for i := 0; i < n; i++ {
		for j := range tmp[i] {
			if tmp[i][j] {

			}

		}

	}

	return false
}

func TestFind132pattern(t *testing.T) {
	t.Log(find132pattern([]int{3, 1, 4, 2}))

}

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue struct {
	items []*Item
	less  bool
}

func (pq PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq.less {
		return pq.items[i].priority < pq.items[j].priority
	}
	return pq.items[i].priority > pq.items[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len((*pq).items)
	item := x.(*Item)
	item.index = n
	*&pq.items = append(*&pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len((*pq).items)
	item := (*pq).items[n-1]
	item.index = -1
	(*pq).items = (*pq).items[0 : n-1]
	return item
}
