package leetcode

import (
	"math"
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	ans := math.MaxInt32 + target

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			//去除重复元素
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			tmp := nums[i] + nums[left] + nums[right]

			if tmp < target {
				left++

			} else if tmp > target {
				right--
			} else {
				return target //相等直接返回
			}

			if abs(tmp-target) < abs(ans-target) {
				ans = tmp
			}
		}
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func TestThreeSumClosest(t *testing.T) {
	// t.Log(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	// t.Log(threeSumClosest([]int{1, 1, 1, 1}, 0))
	t.Log(threeSumClosest([]int{-3, -2, -5, 3, -4}, -1))
}
