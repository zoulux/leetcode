package main

import "fmt"

func main() {

	fmt.Println(countVowelStrings(1))
	fmt.Println(countVowelStrings(2))
	fmt.Println(countVowelStrings(3))
	fmt.Println(countVowelStrings(4))
	//fmt.Println(countVowelStrings(33))
}

func countVowelStrings(n int) int {

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 5)
	}

	for i := 0; i < 5; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] += dp[i-1][k]
			}
		}
	}

	ans := 0
	for i := 0; i < 5; i++ {
		ans += dp[n-1][i]
	}
	return ans
}

func countVowelStrings1(n int) int {

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 5)
	}

	for i := 0; i < 5; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = dp[i-1][1] + dp[i-1][0]
		dp[i][2] = dp[i-1][2] + dp[i-1][1] + dp[i-1][0]
		dp[i][3] = dp[i-1][3] + dp[i-1][2] + dp[i-1][1] + dp[i-1][0]
		dp[i][4] = dp[i-1][4] + dp[i-1][3] + dp[i-1][2] + dp[i-1][1] + dp[i-1][0]
	}

	ans := 0
	for i := 0; i < 5; i++ {
		ans += dp[n-1][i]
	}
	return ans
}
