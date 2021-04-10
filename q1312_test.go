package leetcode


func minInsertions(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return 0
	}
	dp := make([]int, n)

	for i := n - 2; i >= 0; i-- {
		pre := 0
		for j := i + 1; j < n; j++ {
			temp := dp[j]
			if s[i] == s[j] {
				dp[j] = pre
			} else {
				dp[j] = Min(dp[j-1], dp[j]) + 1
			}
			pre = temp
		}
	}
	return dp[n-1]
}
