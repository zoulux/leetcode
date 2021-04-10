package leetcode

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, elem := range nums {
		if v, ok := m[target-elem]; ok {
			return []int{v, idx}
		}
		m[elem] = idx
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	t.Log(twoSum([]int{2, 7, 11, 15}, 9))
}
