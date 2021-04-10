package leetcode

import "testing"

func searchRange2(nums []int, target int) []int {

	min := bSearch2(nums, target, true)
	max := bSearch2(nums, target, false)

	return []int{min, max}
}

func bSearch2(nums []int, target int, min bool) int {

	ans, left, right := -1, 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2

		// 1234567 mid:4 target 3
		if nums[mid] < target {
			left = mid + 1

		} else if nums[mid] > target {
			right = mid - 1

		} else {
			ans = mid

			if min {
				right = mid - 1
			} else {
				left = mid + 1
			}

		}

	}

	return ans
}
func searchRange(nums []int, target int) []int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			l, r := mid, mid
			for l-1 >= 0 && nums[l-1] == nums[l] {
				l--
			}

			for r+1 < len(nums) && nums[r] == nums[r+1] {
				r++
			}

			return []int{l, r}

		} else if target > nums[mid] {

			left = mid + 1

		} else {
			right = mid - 1

		}

	}
	return []int{-1, -1}
}

func TestSearchRange(t *testing.T) {
	t.Log(searchRange2([]int{5, 7, 7, 8, 8, 10}, 8))
	t.Log(searchRange2([]int{5, 7, 7, 8, 8, 10}, 6))
	t.Log(searchRange2([]int{}, 0))
	t.Log(searchRange2([]int{2, 2}, 2))
}
