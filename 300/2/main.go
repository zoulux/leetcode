package main

func main() {

}

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	mx := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		mx = max(mx, dp[i])
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
