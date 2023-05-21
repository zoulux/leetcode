package main

import "fmt"

func main() {

	fmt.Println(longestSubarray([]int{
		0, 1, 1, 1, 0, 1, 1, 0, 1,
	}))
	fmt.Println(longestSubarray([]int{
		1, 1, 0, 1,
	}))
}

func longestSubarray(nums []int) int {

	mx := 0
	left, right := 0, 0
	zero := 0
	for right < len(nums) {
		if nums[right] == 0 {
			zero++
		}
		right++

		for zero > 1 {
			if nums[left] == 0 {
				zero--
			}
			left++
		}

		if t := right - left; t > mx {
			mx = t
		}
	}
	return mx - 1
}
