package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(minimumTotal([][]int{
		{-1},
		{-2, -3},
	}))

	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}

func minimumTotal(triangle [][]int) int {
	m, n := len(triangle), len(triangle[len(triangle)-1])
	dp := make([][]int, m) // 弄个二维数组，存一下一路上存储的和
	for i := range dp {
		dp[i] = make([]int, len(triangle[i])+2)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt / 2 // 赋一个初始值
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp[0][1] = triangle[0][0]
	for i := 1; i < m; i++ {
		for j := 1; j <= len(triangle[i]); j++ {
			dp[i][j] = triangle[i][j-1] + min(dp[i-1][j], dp[i-1][j-1])
		}
	}

	x := math.MaxInt / 2
	for _, v := range dp[n-1] {
		if v < x {
			x = v
		}
	}
	return x
}

func minimumTotal3(triangle [][]int) int {
	m, n := len(triangle), len(triangle[len(triangle)-1])
	dp := make([][]int, m) // 弄个二维数组，存一下一路上存储的和
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
		for j := range dp[i] {
			dp[i][j] = math.MaxInt / 2 // 赋一个初始值
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp[0][0] = triangle[0][0]

	for i := 1; i < m; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j <= len(triangle[i-1])-1 {
				dp[i][j] = triangle[i][j] + dp[i-1][j]
			}

			if j-1 >= 0 && j-1 <= len(triangle[i-1])-1 {
				dp[i][j] = min(dp[i][j], triangle[i][j]+dp[i-1][j-1])
			}
		}
	}

	x := math.MaxInt / 2
	for _, v := range dp[n-1] {
		if v < x {
			x = v
		}
	}
	return x
}

func minimumTotal2(triangle [][]int) int {
	m, n := len(triangle), len(triangle[len(triangle)-1])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
		for j := range dp[i] {
			dp[i][j] = math.MaxInt / 2
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp[0][0] = triangle[0][0]

	for i := 1; i < m; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j <= len(triangle[i-1])-1 {
				dp[i][j] = triangle[i][j] + dp[i-1][j]
			}

			if j-1 >= 0 && j-1 <= len(triangle[i-1])-1 {
				dp[i][j] = min(dp[i][j], triangle[i][j]+dp[i-1][j-1])
			}
		}
	}

	x := math.MaxInt / 2

	for _, v := range dp[n-1] {
		if v < x {
			x = v
		}
	}
	return x
}
