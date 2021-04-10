package leetcode

import (
	"testing"
)

func maxSubArray(nums []int) (ans int) {
	n := len(nums)
	ans = nums[0]

	f := make([]int, n+1) // 前缀和

	for i := 1; i <= n; i++ {
		num := nums[i-1]
		f[i] = f[i-1] + num
	}

	for i := 0; i < len(f); i++ {
		for j := i + 1; j < len(f); j++ {
			tmp := f[j] - f[i]
			if tmp > ans {
				ans = tmp
			}

		}

	}
	return
}

func maxSubArray2(nums []int) int {
	l, max := len(nums), nums[0]

	for i := 1; i < l; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func TestMaxSubArray(t *testing.T) {
	t.Log(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
