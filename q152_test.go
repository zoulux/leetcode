package leetcode

import (
	"testing"
)

func maxProduct(nums []int) int {
	n := len(nums)

	sum := make([]int, n+1)
	sum[0] = 1
	for i := 0; i < n; i++ {
		if i > 0 && sum[i] != 0 {
			sum[i+1] = sum[i] * nums[i]
		} else {
			sum[i+1] = nums[i]
		}
	}
	ans := nums[0]
	for i := 0; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			tmp := 0
			if sum[i] == 0 {
				tmp = sum[j]
			} else {
				tmp = sum[j] / sum[i]
			}

			if tmp > ans {
				ans = tmp
			}
		}
	}

	return ans
}

func TestMaxProduct(t *testing.T) {
	// t.Log(maxProduct([]int{2, 3, -2, 4}))
	// t.Log(maxProduct([]int{-2, 0, -1}))
	// t.Log(maxProduct([]int{-4, -3}))
	// t.Log(maxProduct([]int{-4}))
	// t.Log(maxProduct([]int{0, 2}))
	t.Log(maxProduct([]int{-1, 0, -2, 2}))

}
