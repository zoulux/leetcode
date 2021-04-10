package leetcode

import "testing"

func TestSearch(t *testing.T) {

	// [-1,0,3,5,9,12]
	// 9

	// a := []int{-1, 0, 3, 5, 9, 12}
	// target := 9
	// t.Log(search(a, target))

	// a = []int{5}
	// target = 5
	// t.Log(search(a, target))

	// t.Log(search([]int{2, 5}, 5))
	// t.Log(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	// t.Log(search([]int{5}, 5))
	t.Log(search([]int{2, 5}, 0))
}

func sb(arr []int, target int, start int, end int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		if arr[0] == target {
			return 0
		}
		return -1
	}

	if start < 0 || end < 0 {
		return -1
	}

	if arr[start] == target {
		return start
	}

	if arr[end] == target {
		return end
	}

	if start >= end {
		return -1
	}

	mid := start + (end-start)/2
	mValue := arr[mid]

	if target == mValue {
		return mid
	}

	if target < mValue {
		return sb(arr, target, start, mid-1)
	}

	return sb(arr, target, mid+1, end)
}

func search(nums []int, target int) int {
	return sb(nums, target, 0, len(nums)-1)
}
