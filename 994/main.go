package main

import "fmt"

func main() {

	fmt.Println(orangesRotting([][]int{
		{0},
	}))

	// -1
	fmt.Println(orangesRotting([][]int{
		{1},
	}))

	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}))

	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}))
}

func orangesRotting(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	queue := make([][3]int, 0)

	oneCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				// 	找坏的
				queue = append(queue, [3]int{i, j, 0})
				grid[i][j] = 0
			} else if grid[i][j] == 1 {
				// 好的
				oneCount++
			}
		}
	}

	//返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1
	if len(queue) == 0 {
		// 没有坏的
		if oneCount > 0 {
			return -1
		}
		return 0
	}

	var ans int
	for len(queue) > 0 {
		top := queue[0]
		x, y, z := top[0], top[1], top[2]
		queue = queue[1:]
		if z > ans {
			ans = z
		}

		for _, d := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && ny >= 0 && nx < n && ny < m && grid[nx][ny] == 1 {
				oneCount--
				queue = append(queue, [3]int{nx, ny, z + 1})
				grid[nx][ny] = 0
			}
		}
	}

	if 0 != oneCount {
		// 存在未处理的 1
		return -1
	}
	return ans
}
