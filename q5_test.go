package leetcode

import (
	"testing"
)

func longestPalindrome(s string) string {
	n := len(s)

	if n <= 1 {
		return s
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	maxStr := ""

	for j := 1; j < n; j++ {
		for i := 0; i <= j && i < n; i++ {
			if s[i] != s[j] {
				dp[i][j] = 0
			} else {
				if j-i <= 2 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] == 1 {
				if len(s[i:j+1]) > len(maxStr) {
					maxStr = s[i : j+1]
				}

			}
		}
	}
	return maxStr
}

func TestLongestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("a"))
	t.Log(longestPalindrome("ac"))
	t.Log(longestPalindrome("babad"))
}
