package main

import "fmt"

//地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，
//它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
//例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

//示例 1：
//输入：m = 2, n = 3, k = 1
//输出：3

//示例 2：
//输入：m = 3, n = 1, k = 0
//输出：1

func main() {
	fmt.Println(movingCount(2, 3, 1))
	//fmt.Println(movingCount(2, 3, 0))
	fmt.Println(movingCount(36, 38, 18))
}

func movingCount(m int, n int, k int) int {
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}

	var travel func(i, j int) int
	travel = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || (i/10+i%10+j/10+j%10) > k || visit[i][j] {
			return 0
		}
		visit[i][j] = true
		return 1 + travel(i-1, j) + travel(i, j-1) + travel(i+1, j) + travel(i, j+1)
	}

	return travel(0, 0)
}
