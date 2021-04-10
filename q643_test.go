package leetcode

import "math"

func findMaxAverage(nums []int, k int) float64 {

	winSum := 0
	maxSum := math.MinInt64

	for i, v := range nums {
		winSum += v
		if i >= k {
			winSum -= nums[i-k]
		}

		if i >= k-1 {
			maxSum = max(maxSum, winSum)
		}
	}

	return float64(maxSum) / float64(k)
}
