package main

import "fmt"

func main() {

	fmt.Println(maxSubArray([]int{
		-2, 1, -3, 4, -1, 2, 1, -5, 4,
		//-2, -1,
	}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i]
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}

		if dp[i] > max {
			max = dp[i]
		}

	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
