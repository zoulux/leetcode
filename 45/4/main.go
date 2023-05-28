package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(jump([]int{
		2, 3, 1, 1, 4,
	}))

	fmt.Println(jump([]int{
		2, 3, 0, 1, 4,
	}))
}

func jump(nums []int) int {
	dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = math.MaxInt / 2
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp[0] = 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]
}
