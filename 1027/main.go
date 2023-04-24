package main

import "fmt"

func main() {

	fmt.Println(longestArithSeqLength([]int{
		3, 6, 9, 12,
	}))
}

func longestArithSeqLength(nums []int) int {
	m := 1001

	dp := make([][]int, len(nums))

	for i := range dp {
		dp[i] = make([]int, m)

	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var ans int
	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			d := nums[i] - nums[j] + m/2
			dp[i][d] = max(dp[i][d], dp[j][d]+1)
			ans = max(ans, dp[i][d])
		}
	}
	return ans + 1
}
