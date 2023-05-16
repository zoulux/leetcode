package main

import "fmt"

func main() {

	fmt.Println(jump22([]int{
		2, 3, 1, 1, 4,
	}))
}

func jump(nums []int) int {
	//max := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}

	n := len(nums)
	maxPos := 0
	step := 0

	for i := 0; i < n-1; i++ {
		if d := i + nums[i]; d > maxPos {
			maxPos = d
		}

		if i == maxPos {
			step++
		}
	}
	return step
}

func jump22(nums []int) int {

	//输入: nums = [2,3,1,1,4]
	//	输出: 2
	//	解释: 跳到最后一个位置的最小跳跃数是 2。
	//	从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

	n := len(nums)
	dp := make([]int, n)
	// 以 i 位置需要跳多少步

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := range dp {
		dp[i] = n + 1
	}

	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}
