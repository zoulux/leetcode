package main

import "fmt"

func main() {

	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{1, 1, 1}))
	fmt.Println(longestSubarray([]int{0}))
	fmt.Println(longestSubarray([]int{0, 0}))

}

func longestSubarray(nums []int) int {

	mx := 0
	zero := 0
	for left, right := 0, 0; right < len(nums); {
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

		if d := right - left; d > mx {
			mx = d
		}
	}
	return mx - 1
}
