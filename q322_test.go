package leetcode

import (
	"math"
	"testing"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func dp(coins []int, amount int, dpMap map[int]int, result []int) int {

	if v, ok := dpMap[amount]; ok {
		return v
	}

	if amount == 0 {
		dpMap[amount] = 0
		return dpMap[amount]
	}

	if amount < 0 {
		dpMap[amount] = -1
		return dpMap[amount]
	}

	res := math.MaxInt64
	for _, coin := range coins {
		sub := dp(coins, amount-coin, dpMap, result)
		if sub == -1 {
			continue
		}
		if 1+sub < res {
			res = 1 + sub
			result[amount] = coin
			// result = append(result, coin)
		}
	}
	if res != math.MaxInt64 {
		dpMap[amount] = res
		return dpMap[amount]
	}
	dpMap[amount] = -1
	return dpMap[amount]
}

func coinChange(coins []int, amount int) (int, []int, map[int]int) {
	dpMap := make(map[int]int)
	result := make([]int, amount+1)
	return dp(coins, amount, dpMap, result), result, dpMap
}

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11

	// coins := []int{2}
	// amount := 3
	t.Log(coinChange(coins, amount))
}
