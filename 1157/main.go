package main

import (
	"fmt"
	"sort"
)

func main() {

	mc := Constructor([]int{
		1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1,
	})

	fmt.Println(mc.Query(3, 12, 6))
	//fmt.Println(mc.Query(0, 3, 3))
	//fmt.Println(mc.Query(2, 3, 2))
}

func main1() {

	mc := Constructor([]int{
		1, 1, 2, 2, 1, 1,
	})

	fmt.Println(mc.Query(0, 5, 4))
	fmt.Println(mc.Query(0, 3, 3))
	fmt.Println(mc.Query(2, 3, 2))
}

type MajorityChecker struct {
	cache []Item
}

type Item struct {
	k  int
	v  []int
	ln int
}

func Constructor(arr []int) MajorityChecker {
	mm := make(map[int][]int)

	for i := 0; i < len(arr); i++ {
		v := arr[i]
		mm[v] = append(mm[v], i)
	}

	cache := make([]Item, 0)

	for k, v := range mm {
		cache = append(cache, Item{
			k:  k,
			v:  v,
			ln: len(v),
		})
	}

	sort.Slice(cache, func(i, j int) bool {
		return cache[i].ln > cache[j].ln
	})

	return MajorityChecker{cache: cache}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	for _, v := range this.cache {
		if v.ln < threshold {
			return -1
		}

		l := getBound(v.v, left, true)
		r := getBound(v.v, right, false)
		if r-l+1 >= threshold {
			return v.k
		}
	}

	return -1
}

func getBound(arr []int, target int, isLeft bool) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) >> 1
		if target > arr[mid] {
			left = mid + 1
		} else if target < arr[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	if isLeft {
		return left
	}
	return right
}

func (this *MajorityChecker) Query2(left int, right int, threshold int) int {
	for _, v := range this.cache {
		if v.ln < threshold {
			return -1
		}

		l, r := 0, v.ln-1
		for ; l < v.ln; l++ {
			if v.v[l] >= left {
				break
			}
		}

		for ; r >= 0; r-- {
			if v.v[r] <= right {
				break
			}
		}

		if r-l+1 >= threshold {
			return v.k
		}

	}

	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
