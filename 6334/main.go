package main

func main() {

}

func findPrefixScore(nums []int) []int64 {

	dp := make([]int64, len(nums))

	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		dp[i] = int64(max + nums[i])
		if i > 0 {
			dp[i] += dp[i-1]
		}
	}

	return dp
}
