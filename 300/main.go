package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{
		10, 9, 2, 5, 3, 7, 101, 18,
	}))

	//输入：nums = [10,9,2,5,3,7,101,18]
}

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var ans int

	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		ans = max(ans, dp[i])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
