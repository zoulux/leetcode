package main

import "fmt"

func main() {

	fmt.Println(longestOnes([]int{
		1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0,
	}, 2))
}

func longestOnes(nums []int, k int) int {
	mx := 0
	for zeroCount, left, right := 0, 0, 0; right < len(nums); {
		if nums[right] == 0 {
			zeroCount++
		}
		right++

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		if d := right - left; d > mx {
			mx = d
		}

	}
	return mx
}

func longestOnes2(nums []int, k int) int {
	mx := 0
	zeroCount := 0
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}
		right++

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		if d := right - left; d > mx {
			mx = d
		}

	}
	return mx
}
