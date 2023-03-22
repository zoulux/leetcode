package main

import "fmt"

func main() {

	fmt.Println(cuttingRope(2))
	fmt.Println(cuttingRope(3))
	fmt.Println(cuttingRope(10))

}

func cuttingRope(n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}
