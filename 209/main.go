package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(2 == minSubArrayLen(7, []int{
		2, 3, 1, 2, 4, 3,
	}))

	fmt.Println(1 == minSubArrayLen(4, []int{
		1, 4, 4,
	}))

	fmt.Println(0 == minSubArrayLen(11, []int{
		1, 1, 1, 1, 1, 1, 1, 1,
	}))
}

func minSubArrayLen(target int, nums []int) int {

	sum := make([]int, len(nums)+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}

	ans := math.MaxInt / 2

	left, right := 0, 0

	win := 0

	for right < len(nums) {
		win += nums[right]
		right++

		for win >= target {
			if d := right - left; d < ans {
				ans = d
			}

			win -= nums[left]

			left++
		}
	}

	if ans == math.MaxInt/2 {
		return 0
	}

	return ans
}

func minSubArrayLen2(target int, nums []int) int {

	sum := make([]int, len(nums)+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}

	// 暴力会超时
	ans := math.MaxInt / 2
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if sum[j]-sum[i] >= target {
				if j-i < ans {
					ans = j - i
				}
			}
		}
	}
	if ans == math.MaxInt/2 {
		return 0
	}
	return ans
}
