package leetcode

import "testing"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	var dp func(i, j int) int

	mem := make([][]int, m+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dp = func(i, j int) int {
		if obstacleGrid[i][j] == 1 {
			return 0
		}

		if i == m-1 && j == n-1 {
			return 1
		}

		if i >= m || j >= n {
			return 0
		}

		if mem[i][j] != -1 {
			return mem[i][j]
		}

		tmp := 0
		if i+i < m && obstacleGrid[i+1][j] == 0 {
			tmp += dp(i+1, j)
		}

		if j+1 < n && obstacleGrid[i][j+1] == 0 {
			tmp += dp(i, j+1)
		}

		mem[i][j] = tmp
		return tmp
	}

	return dp(0, 0)
}

func TestUniquePathsWithObstacles(t *testing.T) {
	m := [][]int{
		{0, 1},
		{0, 0},
	}

	// m := [][]int{
	// 	{1},
	// }
	t.Log(uniquePathsWithObstacles(m))
}
