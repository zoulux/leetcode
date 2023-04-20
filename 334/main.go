package main

import "fmt"

func main() {
	fmt.Println(increasingTriplet([]int{
		1, 2, 3, 4, 5,
	}))

	fmt.Println(increasingTriplet([]int{
		5, 4, 3, 2, 1,
	}))
	fmt.Println(increasingTriplet([]int{
		2, 1, 5, 0, 4, 6,
	}))

}

func increasingTriplet(nums []int) bool {
	ln := len(nums)
	dpmin := make([]int, ln)
	dpmax := make([]int, ln)

	dpmin[0] = nums[0]
	for i := 1; i < ln; i++ {
		dpmin[i] = nums[i]
		if dpmin[i-1] < dpmin[i] {
			dpmin[i] = dpmin[i-1]
		}
	}

	dpmax[ln-1] = nums[ln-1]
	for i := ln - 2; i >= 0; i-- {
		dpmax[i] = nums[i]
		if dpmax[i+1] > dpmax[i] {
			dpmax[i] = dpmax[i+1]
		}
	}

	for i := 0; i < ln; i++ {
		if nums[i] > dpmin[i] && nums[i] < dpmax[i] {
			return true
		}
	}

	return false
}
