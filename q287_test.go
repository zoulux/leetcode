package leetcode

import (
	"testing"
)

func findDuplicate(nums []int) int {
	fast, slow := 0, 0 // 索引

	slow = nums[slow]       // 指向下一个
	fast = nums[nums[fast]] // 指向下下个

	for slow != fast { // 一直到相遇
		slow = nums[slow]       // 指向下一个
		fast = nums[nums[fast]] // 指向下下个
	}

	fast = 0
	for slow != fast { // 继续前进 直到相遇就是环的入口点
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func TestFindDuplicate(t *testing.T) {
	// t.Log(findDuplicate([]int{1, 3, 4, 2, 2}))
	// t.Log(findDuplicate([]int{3,1,3,4,2}))
	t.Log(findDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 4, 7, 1}))
	// t.Log(findDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}))
}
