package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit2(prices []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := math.MinInt
	dp := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			dp[i] = max(dp[i], prices[j]-prices[i])
		}

		if dp[i] > ans {
			ans = dp[i]
		}
	}

	if ans > 0 {
		return ans
	}
	return 0
}

func maxProfit(prices []int) int {

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if len(prices) <= 1 {
		return 0
	}

	buy := prices[0] //买入价格
	profit := 0      //利润

	for _, p := range prices {
		if x := p - buy; x > profit {
			profit = x
		}
		buy = min(buy, p)
	}

	return profit
}
