package leetcode

import (
	"testing"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}
	ans := 1

	left, right := 0, 0

	for left < n && right+1 < n {

		if nums[right] < nums[right+1] {
			right++
		} else {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
	}

	return ans
}

func TestFirstMissingPositive(t *testing.T) {
	// t.Log(firstMissingPositive([]int{2, 1}))
	// t.Log(firstMissingPositive([]int{1, 2}))
	// t.Log(firstMissingPositive([]int{1, 2, 0}))
	// t.Log(firstMissingPositive([]int{3, 4, -1, 1}))
	t.Log(firstMissingPositive([]int{7, 8, 9, 11, 12}))

	// t.Log(firstMissingPositive([]int{1, 2, 6, 3, 5, 4}))

}
