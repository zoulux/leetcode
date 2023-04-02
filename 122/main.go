package main

import "fmt"

func main() {

	fmt.Println(maxProfit([]int{
		7, 1, 5, 3, 6, 4,
	}))

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

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dpbuy := make([]int, len(prices))
	// 在 i 天 买获得的利润
	dpsell := make([]int, len(prices))
	// 在 i 天 卖获得的利润

	// 第一天持仓的话，收益为 -prices[0]
	dpbuy[0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dpbuy[i] = max(
			dpbuy[i-1],            // 股票昨天买的，今天继续持有
			dpsell[i-1]-prices[i], // 今天开买，昨天空仓的收益 - 今天的支出
		)

		dpsell[i] = max(
			dpsell[i-1],          // 今天继续空仓
			dpbuy[i-1]+prices[i], // 昨天持仓收益 + 今天股票卖出收益
		)
	}

	return dpsell[len(prices)-1]
}
