package main

import "fmt"

func main() {

	fmt.Println(longestPalindrome("babad"))

}

func longestPalindrome(s string) string {
	n := len(s)

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	var ln int
	var ii, jj int

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
			}

			if dp[i][j] && j-i > ln {
				ln = j - i
				ii, jj = i, j
			}
		}
	}

	return s[ii : jj+1]
}
