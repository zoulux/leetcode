package main

import "fmt"

func main() {

	fmt.Println(lenx(333))
}

func findColumnWidth(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			ans[j] = max(ans[j], lenx(grid[i][j]))
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lenx(a int) int {
	if a == 0 {
		return 1
	}

	if a < 0 {
		return 1 + lenx(-a)
	}

	x := 0
	for a != 0 {
		a /= 10
		x++
	}
	return x
}
