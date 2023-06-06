package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//输入：grid = [[3,2,1],[1,7,6],[2,7,7]]
	//输出：1
	//解释：存在一对相等行列对：
	//- (第 2 行，第 1 列)：[2,7,7]

	fmt.Println(equalPairs([][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}))

}

func equalPairs(grid [][]int) int {
	n := len(grid)
	mm := make(map[string]int)
	for i := 0; i < n; i++ {
		var t []string
		for j := 0; j < n; j++ {
			t = append(t, strconv.Itoa(grid[i][j]))
		}
		mm[strings.Join(t, ",")]++
	}

	var ans int
	for j := 0; j < n; j++ {
		var t []string
		for i := 0; i < n; i++ {
			t = append(t, strconv.Itoa(grid[i][j]))
		}
		ans += mm[strings.Join(t, ",")]
	}
	return ans
}
