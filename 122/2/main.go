package main

import "fmt"

func main() {

	fmt.Println(maxProfit([]int{
		7, 1, 5, 3, 6, 4,
	}))

}

func maxProfit(prices []int) int {
	var ans int
	for i := 1; i < len(prices); i++ {
		if d := prices[i] - prices[i-1]; d > 0 {
			ans += d
		}
	}
	return ans
}

func maxProfit2(prices []int) int {
	n := len(prices)
	dp0 := make([]int, n) // 表示不持有股票
	dp1 := make([]int, n) // 表示持有股票
	dp1[0] = -prices[0]

	for i := 1; i < n; i++ {
		dp0[i] = max(dp0[i-1], prices[i]+dp1[i-1])
		dp1[i] = max(dp1[i-1], dp0[i-1]-prices[i])
	}

	return dp0[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
