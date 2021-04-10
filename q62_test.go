package leetcode

import "testing"

func uniquePaths2(m int, n int) int {

	var dp func(i, j int) int

	mem := make([][]int, m+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dp = func(i, j int) int {
		if i == m && j == n {
			return 1
		}

		if i > m || j > n {
			return 0
		}

		if mem[i][j] != -1 {
			return mem[i][j]
		}

		mem[i][j] = dp(i+1, j) + dp(i, j+1)

		return mem[i][j]
	}

	return dp(1, 1)
}

func uniquePaths(m int, n int) int {

	visited := make([][]bool, m)

	for i := range visited {
		visited[i] = make([]bool, n)
	}

	ans := 0
	var dfs func(starti, startj int)
	dfs = func(starti, startj int) {

		if starti == m-1 && startj == n-1 {
			ans++
		}

		if starti == m || startj == n {
			return
		}

		if visited[starti][startj] {
			return
		}

		visited[starti][startj] = true
		dfs(starti+1, startj)
		dfs(starti, startj+1)
		visited[starti][startj] = false

	}

	dfs(0, 0)
	return ans
}

func TestUniquePaths(t *testing.T) {

	t.Log(uniquePaths2(3, 7))
	t.Log(uniquePaths2(3, 2))
	// t.Log(uniquePaths(3, 2))
	t.Log(uniquePaths2(23, 12))
}
