package main

import "fmt"

//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {

	dp := make([]int, len(nums))

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		dp[i] = len(nums) + 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(nums)-1]

}

func jump2(nums []int) int {
	pos := len(nums) - 1
	step := 0
	for pos > 0 {
		for i := 0; i < pos; i++ {
			if i+nums[i] >= pos {
				pos = i
				step++
				break
			}
		}
	}
	return step
}
