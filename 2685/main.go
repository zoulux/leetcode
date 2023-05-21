package main

import "fmt"

func main() {
	fmt.Println(countCompleteComponents(3, [][]int{
		{1, 0},
		{2, 1},
	}))

	fmt.Println(countCompleteComponents(6, [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{3, 4},
	}))
}

func countCompleteComponents(n int, edges [][]int) int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	for _, e := range edges {
		grid[e[0]][e[1]] = 1
		grid[e[1]][e[0]] = 1
	}

	mm := make(map[int]bool)

	var ans int

	for i := 0; i < n; i++ {
		if mm[i] {
			continue
		}

		//var jj []int
		jj := []int{}
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// 从 i 出发的
				jj = append(jj, j)
			}
			if grid[j][i] == 1 {
				// 从 i 出发的
				jj = append(jj, j)
			}
		}

		if handle(jj, grid) {
			ans++
		}

		for _, v := range jj {
			mm[v] = true
		}
	}

	return ans
}

func handle(jj []int, grid [][]int) bool {

	for ji := 0; ji < len(jj); ji++ {
		for jt := ji + 1; jt < len(jj); jt++ {
			if grid[jj[ji]][jj[jt]] == 0 {
				return false
			}
		}
	}
	return true
}
