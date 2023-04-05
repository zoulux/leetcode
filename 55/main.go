package main

import "fmt"

// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
func main() {

	fmt.Println(canJump([]int{
		//2, 3, 1, 1,
		3, 2, 1, 0, 4,
	}))

}

func canJump(nums []int) bool {

	maxPos := 0

	for i := 0; i < len(nums); i++ {
		if d := i + nums[i]; d > maxPos {
			maxPos = d
		}
	}

	return maxPos >= len(nums)-1
}

func canJump2(nums []int) bool {
	dp := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i; j <= i+nums[i] && j < len(nums); j++ {
			dp[j] = true
		}

		if i+1 < len(nums) && !dp[i+1] {
			return false
		}
	}
	return true
}
