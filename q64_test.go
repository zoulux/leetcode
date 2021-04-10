package leetcode

import (
	"math"
	"testing"
)

func minPathSum(grid [][]int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	m, n := len(grid), len(grid[0])

	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i < 0 || j < 0 {
			return math.MaxInt64
		}

		if i == 0 && j == 0 {
			return grid[i][j]
		}

		if mem[i][j] != -1 {
			return mem[i][j]
		}

		ret := min(dp(i, j-1), dp(i-1, j)) + grid[i][j]

		mem[i][j] = ret

		return ret
	}

	return dp(m-1, n-1)
}

func TestMinPathSum(t *testing.T) {
	m := [][]int{
		// {1, 3, 1},
		// {1, 5, 1},
		// {4, 2, 1},

		{1, 2, 3},
		{4, 5, 6},
	}
	t.Log(minPathSum(m))

}
