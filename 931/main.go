package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minFallingPathSum([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))
	//
	fmt.Println(minFallingPathSum([][]int{
		{17, 82},
		{1, -44},
	}))

}

func minFallingPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+2)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt / 2
		}
	}

	min := func(arr ...int) int {
		x := arr[0]
		for _, v := range arr[1:] {
			if v < x {
				x = v
			}
		}
		return x
	}

	for i := 0; i < len(matrix[0]); i++ {
		dp[1][i] = matrix[0][i]
		//dp[1][i] = matrix[0][i]
	}

	for i := 2; i <= len(matrix); i++ {
		for j := 0; j < len(matrix[i-1]); j++ {
			dp[i][j] = dp[i-1][j] + matrix[i-1][j]
			if j-1 >= 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+matrix[i-1][j])
			}

			if j+1 < n {
				dp[i][j] = min(dp[i][j], dp[i-1][j+1]+matrix[i-1][j])
			}
		}
	}

	return min(dp[len(dp)-1]...)
}
