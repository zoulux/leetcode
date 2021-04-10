package leetcode

import "testing"

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

func removeElement2(nums []int, val int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i] = nums[slow]
			slow++

		}
	}
	return slow
}

func TestRemoveElement2(t *testing.T) {
	t.Log(removeElement2([]int{3, 2, 2, 3}, 3))
	t.Log(removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
