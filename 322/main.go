package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func change(dp []int, i int, v int) {
	fmt.Println("old", dp)
	dp[i] = v
	fmt.Println("new", dp)
}

func coinChange(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	// 0 元 需要 0 个
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] >= amount+1 {
		return -1
	}

	return dp[amount]
}

func coinChange1(coins []int, amount int) int {
	// 题目要求的最终结果是 dp(amount)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	cache := map[int]int{}
	var dp func(n int) int
	dp = func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}

		res := math.MaxInt32
		for _, c := range coins {
			sub := dp(n - c)
			if sub == -1 {
				continue
			}
			res = min(res, sub+1)
		}

		if res == math.MaxInt32 {
			cache[n] = -1
		} else {
			cache[n] = res
		}
		return cache[n]
	}
	return dp(amount)
}
