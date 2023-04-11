package main

import "fmt"

func main() {

	fmt.Println(minPathSum([][]int{
		{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
	}))

}

func minPathSum(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = grid[i][j]
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] += min(dp[i-1][j], dp[i][j-1])
			} else if i-1 >= 0 {
				dp[i][j] += dp[i-1][j]
			} else if j-1 >= 0 {
				dp[i][j] += dp[i][j-1]
			}

		}
	}

	return dp[m-1][n-1]

}
