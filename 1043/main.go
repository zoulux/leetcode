package main

import "fmt"

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{
		1, 15, 7, 9, 2, 5, 10,
	}, 3))

}

func maxSumAfterPartitioning(arr []int, k int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = arr[i]
		for j := max(i-k+1, 0); j < min(i+k, len(arr)); j++ {
			dp[i] = max(dp[i], arr[j])
		}
	}

	fmt.Println(dp)

	var sum int

	for _, v := range dp {
		sum += v

	}

	return sum
}
