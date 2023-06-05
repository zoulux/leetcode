package main

import "fmt"

func main() {

	//fmt.Println(applyOperations([]int{
	//	1, 2, 2, 1, 1, 0,
	//}))
	//
	fmt.Println(applyOperations([]int{
		847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272,
	}))
}

func applyOperations(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	left, right := 0, 0

	//1,4,0,2,0,0
	for left < len(nums) {

		for left < len(nums) && nums[left] != 0 {
			left++
		}

		for right = left + 1; right < len(nums) && nums[right] == 0; right++ {
		}

		if right < len(nums) && left < len(nums) {
			nums[left], nums[right] = nums[right], nums[left]
		}
		left++

	}

	return nums
}
