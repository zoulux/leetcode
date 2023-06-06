package main

import "fmt"

func main() {

	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
	fmt.Println(longestPalindromeSubseq("aabaa"))
}

func longestPalindromeSubseq(s string) int {

	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
