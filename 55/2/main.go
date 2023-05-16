package main

import "fmt"

func main() {

	fmt.Println(canJump([]int{
		2, 3, 1, 1, 4,
	}))

	fmt.Println(canJump([]int{
		3, 2, 1, 0, 4,
	}))

	fmt.Println(canJump([]int{
		0,
	}))

	fmt.Println(canJump([]int{
		0, 2, 3,
	}))
}

func canJump(nums []int) bool {

	maxPos := 0

	for i := 0; i < len(nums); i++ {
		if i > maxPos {
			return false
		}

		if d := i + nums[i]; d > maxPos {
			maxPos = d
		}
	}

	return maxPos >= len(nums)-1
}

func canJump2(nums []int) bool {
	dp := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		// 以当前起点向后跳，如果能跳到标记 true
		for j := i; j < len(nums) && j <= i+nums[i]; j++ {
			dp[j] = true
		}

		// 能不能跳到下一个点，跳不到返回 false
		if i+1 < len(nums) && !dp[i+1] {
			return false
		}
	}

	return true
}
func canJumpErr(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	dp := make([]int, len(nums))

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums)-1; i++ {
		v := nums[i]
		dp[i] = i + v
		for j := 0; j < i; j++ {
			dp[i] = max(dp[i], dp[j])
		}

		if dp[i] >= len(nums)-1 {
			return true
		}
	}

	return false
}
