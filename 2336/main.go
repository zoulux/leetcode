package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{1, 2, 5}
	fmt.Println(sort.SearchInts(arr, 3))

}

type miniHeap struct {
	sort.IntSlice
}

func (h *miniHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *miniHeap) Pop() interface{} {
	ret := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return ret
}

//Push(x any) // add x as element Len()
//Pop() any   // remove and return element Len() - 1.

type SmallestInfiniteSet struct {
	arr []int
}

func Constructor() SmallestInfiniteSet {
	// 因为数字是确定的，那就先将数字对入队列
	arr := make([]int, 0, 1000)
	for i := 1; i <= 1000; i++ {
		arr = append(arr, i)
	}
	return SmallestInfiniteSet{
		arr: arr,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	x := this.arr[0]        // 对头元素
	this.arr = this.arr[1:] // 删除对头元素
	return x

}

func (this *SmallestInfiniteSet) AddBack(num int) {
	x := this.arr[0]
	if num < x {
		this.arr = append([]int{num}, this.arr...) // 比对头元素小，直接插入在对头
	} else if num > x {
		t := sort.SearchInts(this.arr, num) // 需要插在队列中，那就搜位置
		if t < len(this.arr) {
			// 没有越界
			if this.arr[t] != num { // 相等的话表示数字在队列中，不等的话那就开始处理
				// 数组中不存在，将原队列以 t 为界限化为两段，并将 t 位置换成 num
				this.arr = append(this.arr[:t], append([]int{num}, this.arr[t:]...)...)
			}
		} else {
			this.arr = append(this.arr, num) //如果需要插入在队尾
		}
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
