package leetcode

import (
	"testing"
)

func longestCommonSubsequence(text1 string, text2 string) int {

	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for idx := range dp {
		dp[idx] = make([]int, n+1)
	}

	// i,j 这里面充当两部分，一部分是遍历字符串，第二部分是遍历 dp table
	// 字符串是从 0 开始遍历
	// 数组是从 1 开始遍历
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 // 左上角的结果 +1
			} else {
				if dp[i][j-1] > dp[i-1][j] {
					dp[i][j] = dp[i][j-1] // 上边的结果
				} else {
					dp[i][j] = dp[i-1][j] // 左边的结果
				}
			}
		}
	}
	return dp[m][n]
}

func TestLongestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"
	t.Log(longestCommonSubsequence(text1, text2))
}
