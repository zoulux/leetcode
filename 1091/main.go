package main

import (
	"fmt"
)

func main() {

	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 1},
	}))

	// test63
	fmt.Println(14 == shortestPathBinaryMatrix([][]int{
		{0, 1, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 0},
	}))

	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1},
		{1, 0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
	}))
	//
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 1, 1, 0, 0, 0},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 0, 0, 1, 1, 0},
		{1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 0},
	}))

	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 1},
		{1, 0},
	}))

	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}))

	fmt.Println(shortestPathBinaryMatrix([][]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}))
}

func shortestPathBinaryMatrix(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	q := make([][3]int, 0)

	if grid[0][0] == 0 {
		q = append(q, [...]int{0, 0, 1})
		grid[0][0] = 1
	}

	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		i, j, c := top[0], top[1], top[2]
		if i == m-1 && j == n-1 {
			return c
		}

		for _, v := range [][2]int{{0, 1}, {1, 1}, {-1, 1}, {0, -1}, {1, -1}, {-1, -1}, {1, 0}, {-1, 0}} {
			x, y := i+v[0], j+v[1]
			if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != 0 {
				continue
			}

			q = append(q, [...]int{x, y, c + 1})
			grid[x][y] = 1
		}
	}
	return -1
}

func shortestPathBinaryMatrix2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	q := make([][3]int, 0)

	if grid[0][0] == 0 {
		q = append(q, [3]int{0, 0, 1})
	} else {
		return -1
	}

	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		i, j, c := top[0], top[1], top[2]
		if i == m-1 && j == n-1 {
			return c
		}

		grid[i][j] = 1

		if j+1 < n {
			if grid[i][j+1] == 0 {
				q = append(q, [3]int{i, j + 1, c + 1})
			}

			if i+1 < m && grid[i+1][j+1] == 0 {
				q = append(q, [3]int{i + 1, j + 1, c + 1})
			}

			if i-1 >= 0 && grid[i-1][j+1] == 0 {
				q = append(q, [3]int{i - 1, j + 1, c + 1})
			}
		}

		if j-1 >= 0 {
			if grid[i][j-1] == 0 {
				q = append(q, [3]int{i, j - 1, c + 1})
			}

			if i+1 < m && grid[i+1][j-1] == 0 {
				q = append(q, [3]int{i + 1, j - 1, c + 1})
			}

			if i-1 >= 0 && grid[i-1][j-1] == 0 {
				q = append(q, [3]int{i - 1, j - 1, c + 1})
			}
		}

		if i+1 < m && grid[i+1][j] == 0 {
			q = append(q, [3]int{i + 1, j, c + 1})
		}

		if i-1 >= 0 && grid[i-1][j] == 0 {
			q = append(q, [3]int{i - 1, j, c + 1})
		}
	}
	return -1
}
