package main

import "fmt"

func main() {

	fmt.Println(pivotIndex([]int{
		1, 7, 3, 6, 5, 6,
	}))

	fmt.Println(pivotIndex([]int{
		1, 2, 3,
	}))

	fmt.Println(pivotIndex([]int{
		2, 1, -1,
	}))
}

func pivotIndex(nums []int) int {
	ln := len(nums)
	leftSum := make([]int, ln)
	rightSum := make([]int, ln)

	for i := 0; i < ln; i++ {
		leftSum[i] = nums[i]
		if i > 0 {
			leftSum[i] += leftSum[i-1]
		}
	}

	for i := ln - 1; i >= 0; i-- {
		rightSum[i] = nums[i]
		if i < ln-1 {
			rightSum[i] += rightSum[i+1]
		}
	}

	for i := 0; i < ln; i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}
	return -1
}
