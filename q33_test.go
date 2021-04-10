package leetcode

import "testing"

func search2(nums []int, target int) int {

	rotatedI := 0
	for rotatedI+1 < len(nums) {
		if nums[rotatedI] < nums[rotatedI+1] {
			rotatedI++
		} else {
			break
		}
	}

	// [0 -rotatedI] 左闭右闭
	// [rotatedI+1 - n-1] 左闭右闭

	if target > nums[rotatedI] {
		return -1
	} else if target == nums[rotatedI] {
		return rotatedI
	}

	ans := bSearch(nums, 0, rotatedI, target)
	if ans == -1 {
		if rotatedI+1 < len(nums) {
			ans = bSearch(nums, rotatedI+1, len(nums)-1, target)
		}
	}

	return ans
}

func bSearch(nums []int, left, right, target int) int {

	for left <= right {
		mid := (left + right) / 2

		// 1234567 mid:4 target 3
		if nums[mid] < target {
			left = mid + 1

		} else if nums[mid] > target {
			right = mid - 1

		} else {
			return mid
		}
	}

	return -1
}

func TestSearch2(t *testing.T) {
	// t.Log(search2([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	// t.Log(search2([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	// t.Log(search2([]int{1}, 0))
	t.Log(search2([]int{1}, 1))
}


