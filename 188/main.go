package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	k = min(k, n/2)
	dp0 := make([][]int, n)
	dp1 := make([][]int, n)
	for i := range dp0 {
		dp0[i] = make([]int, k+1)
		dp1[i] = make([]int, k+1)
	}

	for i := 1; i <= k; i++ {
		dp0[0][i] = 0
		dp1[0][i] = -prices[0]
	}

	for i := 0; i < n; i++ {
		dp0[i][0] = 0
		dp1[i][0] = math.MinInt64
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp0[i][j] = max(dp0[i-1][j], dp1[i-1][j]+prices[i])
			dp1[i][j] = max(dp1[i-1][j], dp0[i-1][j-1]-prices[i])
		}
	}

	return max(dp0[n-1][k])
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
