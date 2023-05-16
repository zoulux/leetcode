package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(runeReserve([]int{
		1,
	}))

	fmt.Println(runeReserve([]int{
		1,
	}))

	fmt.Println(runeReserve([]int{
		1, 3, 5, 4, 1, 7,
	}))

	fmt.Println(runeReserve([]int{
		1, 1, 3, 3, 2, 4,
	}))

}

func runeReserve(runes []int) int {
	sort.Ints(runes)

	ans := 1
	dp := make([]int, len(runes))

	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(runes); i++ {
		if runes[i]-runes[i-1] <= 1 {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}
