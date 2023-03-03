package main

import (
	"fmt"
)

func main() {

	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-1}))

}

func maxSubArray(nums []int) int {
	max := func(arr ...int) int {
		mm := arr[0]
		for _, v := range arr {
			if v > mm {
				mm = v
			}
		}
		return mm
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i] + max(dp[i-1], 0) // 左边的哥们能给我贡献多少，不能贡献那我就算自己的
	}

	ans := nums[0]
	for _, v := range dp {
		ans = max(ans, v) // 看一下最大值
	}
	return ans
}
