package main

import "math"

func main() {

}

func jump(nums []int) int {

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = math.MaxInt / 2
	}

	dp[0] = 0 // 以 i 为终点需要多少步
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]
}
