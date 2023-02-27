package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0}))
	fmt.Println(missingNumber([]int{0, 1, 3}))
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 5, 6, 7, 8, 9}))

}

// 0,1,2,3,5,6,7,8,9
func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) >> 1
		if mid == nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
