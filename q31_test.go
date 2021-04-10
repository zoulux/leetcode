package leetcode

import (
	"testing"
)

// 123465

func nextPermutation(nums []int) {

	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// sort.Ints(nums[i+1:])
	reverse3(nums[i+1:])
}

func reverse3(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 3, 4, 6, 5}
	nextPermutation(nums)
	t.Log(nums)
}
