package main

import (
	"fmt"
)

func main() {

	fmt.Println(maxProfit([]int{
		1, 2, 3, 4, 5,
	}))

	fmt.Println(maxProfit([]int{
		3, 3, 5, 0, 0, 3, 1, 4,
	}))
}

func maxProfit(prices []int) int {
	n := len(prices)
	k := 2
	if n/2 < k {
		k = n / 2
	}

	dp0 := make([][]int, n) // 表示不持有股票 第 i 天不持有第 k
	dp1 := make([][]int, n) // 表示持有股票
	for i := range dp0 {
		dp0[i] = make([]int, k+1)
		dp1[i] = make([]int, k+1)
	}

	for i := 1; i <= k; i++ {
		//dp1[0][i] = math.MinInt64 / 2
		//dp0[0][i] = math.MinInt64 / 2
	}
	dp1[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp1[i][0] = max(dp1[i-1][0], dp0[i-1][0]-prices[i])

		for j := 1; j <= k; j++ {
			dp0[i][j] = max(
				dp0[i-1][j],
				dp1[i-1][j-1]+prices[i],
			)

			dp1[i][j] = max(
				dp1[i-1][j],
				dp0[i-1][j]-prices[i],
			)
		}
	}

	return maxArr(dp0[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArr(arr []int) int {
	x := arr[0]
	for _, v := range arr[1:] {
		if v > x {
			x = v
		}
	}
	return x
}
