package main

import (
	"fmt"
)

func main() {

	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
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
	return maxProfitK(prices, 2)
}

func maxProfitK(prices []int, k int) int {

	if len(prices) <= 1 {
		return 0
	}

	k = min(k, len(prices)/2)

	// 在 i 天持有完成 k 笔交易的收入
	dpbuy := make([][]int, len(prices))
	// 在 i 天空仓完成 k 笔交易的收入
	dpsell := make([][]int, len(prices))

	//k 表示完成 k 笔交易
	for i := range dpbuy {
		dpbuy[i] = make([]int, k+1)
		dpsell[i] = make([]int, k+1)
	}

	// 第一天持仓的话，收益为 -prices[0]
	// 	第一天空仓的话，收益为 0
	dpbuy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		dpbuy[0][i] = -100
		dpsell[0][i] = -100
	}

	for i := 1; i < len(prices); i++ {
		dpbuy[i][0] = max(
			dpbuy[i-1][0],            // 股票昨天买的，今天继续持有
			dpsell[i-1][0]-prices[i], // 今天开买，昨天空仓的收益 - 今天的支出
		)

		for j := 1; j <= k; j++ {
			dpbuy[i][j] = max(
				dpbuy[i-1][j],            // 股票昨天买的，今天继续持有，【交易的完成的笔数肯定还是上次的】
				dpsell[i-1][j]-prices[i], // 今天开买，昨天空仓的收益 - 今天的支出 【交易完成的笔数今天是+1了，但是对于空仓收益而言笔数是不变的】
			)
			dpsell[i][j] = max(
				dpsell[i-1][j],            // 今天继续空仓【交易笔数不变】
				dpbuy[i-1][j-1]+prices[i], // 今天可以将昨天持有的票卖掉变成空仓状态，昨天持仓收益 + 今天股票卖出收益 【今天开卖的话，】
			)
		}
	}

	fmt.Println(dpbuy)
	fmt.Println(dpsell)
	return max(dpsell[len(dpsell)-1]...)
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
