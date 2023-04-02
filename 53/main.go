package main

import "fmt"

//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

func main() {

	fmt.Println(maxSubArray([]int{
		-2, 1, -3, 4, -1, 2, 1, -5, 4,
	}))
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	max := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
		if i-1 >= 0 && dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func maxSubArray2(nums []int) int {

	sum := 0
	maxSum := nums[0]
	for _, v := range nums {
		sum += v

		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}
