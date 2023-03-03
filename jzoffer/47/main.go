package main

import "fmt"

func main() {
	fmt.Println(maxValue3([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))

}

func maxValue2(grid [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][]int, len(grid)+1)
	for i := 0; i < len(grid)+1; i++ {
		dp[i] = make([]int, len(grid[0])+1)
	}
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			// 向下走
			a := grid[i][j] + dp[i+1][j]
			// 向右走
			b := grid[i][j] + dp[i][j+1]
			dp[i][j] = max(a, b)
		}
	}
	return dp[0][0]
}

func maxValue3(grid [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][]int, len(grid)+1)
	for i := 0; i < len(grid)+1; i++ {
		dp[i] = make([]int, len(grid[0])+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 向下走
			d := 0
			if i-1 >= 0 {
				d += dp[i-1][j]
			}

			// 向右走
			r := 0
			if j-1 >= 0 {
				r += dp[i][j-1]
			}

			dp[i][j] = grid[i][j] + max(d, r)
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

// 暴力递归
func maxValue(grid [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == len(grid) || j == len(grid[0]) {
			return 0
		}
		return grid[i][j] + max(dfs(i+1, j), dfs(i, j+1))
	}
	return dfs(0, 0)
}
