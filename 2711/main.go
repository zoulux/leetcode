package main

import "fmt"

func main() {

	fmt.Println(differenceOfDistinctValues([][]int{
		{1, 2, 3},
		{3, 1, 5},
		{3, 2, 1},
	}))
}

func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	get := func(i, j, flag int) int {
		mm := make(map[int]bool)
		var c int
		for ii, jj := i+flag, j+flag; ii >= 0 && jj >= 0 && ii < m && jj < n; ii, jj = ii+flag, jj+flag {
			if !mm[grid[ii][jj]] {
				c++
				mm[grid[ii][jj]] = true
			}
		}
		return c
	}

	for i, row := range grid {
		for j, _ := range row {
			ans[i][j] = abs(get(i, j, -1) - get(i, j, 1))
		}
	}
	return ans
}
