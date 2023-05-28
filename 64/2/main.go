package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minPathSum([][]int{
		{1},
	}))

	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = math.MaxInt / 2
	}

	for j := range dp[0] {
		dp[0][j] = math.MaxInt / 2
	}

	dp[0][1], dp[1][0] = 0, 0 // 其实转移
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = grid[i-1][j-1] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m][n]
}

func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	cache := make(map[[3]int]int)

	var travel func(i, j, s int) int
	travel = func(i, j, s int) (ret int) {
		if v, ok := cache[[3]int{i, j, s}]; ok {
			return v
		}

		defer func() {
			cache[[3]int{i, j, s}] = ret
		}()

		if i == m-1 && j == n-1 {
			return grid[i][j] + s
		}
		if i >= m || j >= n {
			return math.MaxInt
		}
		return min(
			travel(i+1, j, s+grid[i][j]),
			travel(i, j+1, s+grid[i][j]),
		)
	}
	return travel(0, 0, 0)
}
func minPathSum3(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var travel func(i, j, s int) int
	travel = func(i, j, s int) (ret int) {
		if i == m-1 && j == n-1 {
			return grid[i][j] + s
		}
		if i >= m || j >= n {
			return math.MaxInt
		}
		return min(
			travel(i+1, j, s+grid[i][j]),
			travel(i, j+1, s+grid[i][j]),
		)
	}
	return travel(0, 0, 0)
}
