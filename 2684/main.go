package main

import "fmt"

func main() {
	//输入：grid = [[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]

	//fmt.Println(maxMoves([][]int{
	//	{2, 4, 3, 5},
	//	{5, 4, 9, 3},
	//	{3, 4, 2, 11},
	//	{10, 9, 13, 15},
	//}))

	fmt.Println(maxMoves([][]int{
		{3, 2, 4},
		{2, 1, 9},
	}))
}

func maxMoves(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	cache := make(map[[2]int]int)

	var travel func(i, j int) int
	travel = func(row, col int) (ret int) {
		//(row - 1, col + 1)、(row, col + 1) 和 (row + 1, col + 1)
		if v, ok := cache[[2]int{row, col}]; ok {
			return v
		}
		defer func() {
			if ret != 0 {
				cache[[2]int{row, col}] = ret
			}
		}()

		max := 0
		if col+1 < n {
			if row-1 >= 0 && grid[row-1][col+1] > grid[row][col] {
				if t := 1 + travel(row-1, col+1); t > max {
					max = t
				}
			}

			if grid[row][col+1] > grid[row][col] {
				if t := 1 + travel(row, col+1); t > max {
					max = t
				}
			}

			if row+1 < m && grid[row+1][col+1] > grid[row][col] {
				if t := 1 + travel(row+1, col+1); t > max {
					max = t
				}
			}
		}
		return max
	}

	max := 0
	for r := 0; r < m; r++ {
		if t := travel(r, 0); t > max {
			max = t
		}
	}
	return max
}
