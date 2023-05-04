package main

func main() {

}

func maxProfit(prices []int, fee int) int {

	dp := make([][2]int, len(prices)) // 0 表示不持有的状态下获得的利润，1表示持有状态下获得的利润

	dp[0][1] = -prices[0]

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee) // 如果当前不持有，1.可以将持有的股票卖出，2.可以是上一次不持有
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])     // 如果当前持有，1. 可以购买新的股票，2.拿住之前的股票
	}

	return dp[len(prices)][0]
}
