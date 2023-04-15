package main

import "fmt"

func main() {

	fmt.Println(longestCommonSubsequence("abcde", "ace"))

}

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonSubsequence(text1 string, text2 string) int {

	//dp[i][j]
	// s1[i]==s1[j] => dp[i-1][j-1]+1
	// s1[i]!=s1[j] => max (dp[i][j-1],dp[i-1][j])

	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := 0

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}

			ans = max(ans, dp[i][j])
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
