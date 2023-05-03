package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15}))
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, n+1) //到达当前层需要的花费

	dp[0] = 0 // 第 0 层不需要花费
	dp[1] = 0 // 第 1 层不需要花费

	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2]) // 如果选择从 i-1 层调过来则需要 dp[i-1] + cost
	}

	return dp[n]
}
