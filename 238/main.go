package main

import "fmt"

func main() {

	fmt.Println(productExceptSelf([]int{
		1, 2, 3, 4,
	}))

}

func productExceptSelf(nums []int) []int {

	ln := len(nums)
	left := make([]int, ln+1)
	right := make([]int, ln+1)
	left[0] = 1
	for i := 0; i < ln; i++ {
		left[i+1] = left[i] * nums[i]
	}

	right[ln] = 1
	for j := ln - 1; j >= 0; j-- {
		right[j] = right[j+1] * nums[j]
	}

	fmt.Println(left, right)

	ans := make([]int, ln)
	for i := range ans {
		ans[i] = left[i] * right[i+1]
	}
	return ans
}
