package main

import "fmt"

func main() {

	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 其实对于一行或者一列都只有一种情况
	dp[1][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = max(
				dp[i][j],
				dp[i-1][j]+dp[i][j-1],
			)
		}
	}

	return dp[m][n]
}
