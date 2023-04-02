package main

import "fmt"

// 输入：[7,1,5,3,6,4]
// 输出：5
func main() {
	//fmt.Println(maxProfit([]int{}))
	//
	fmt.Println(maxProfit([]int{
		7, 1, 5, 3, 6, 4,
	}))

	//4

	fmt.Println(maxProfit([]int{
		3, 2, 6, 5, 0, 3,
	}))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// 第 i 天获取到的最低价格
	dp := make([]int, len(prices))
	dp[0] = prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		//当前价格比昨天的最低价格还低的话，今天的最低价就是当前价
		if prices[i] < dp[i-1] {
			dp[i] = prices[i]
		} else {
			// 否则就是昨天的最低价
			dp[i] = dp[i-1]
		}

		// 当前价格 - 最低价 = 利润
		// 如果今天的利润 比历史的利润大，则记录最大利润
		if d := prices[i] - dp[i]; d > max {
			max = d
		}
	}
	return max
}

func maxProfit3(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	max := 0
	low := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		}

		if d := prices[i] - low; d > max {
			max = d
		}
	}

	return max
}

func maxProfit2(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if d := prices[j] - prices[i]; d > max {
				max = d
			}

		}
	}
	return max
}
