package main

import "fmt"

func main() {

	fmt.Println(exchange([]int{1, 2, 3, 4}))
	fmt.Println(exchange([]int{2, 16, 3, 5, 13, 1, 16, 1, 12, 18, 11, 8, 11, 11, 5, 1}))

}

func exchange(nums []int) []int {

	l1, l2 := 0, 1
	// l1 管基数 l2 管偶数

	for l1 < len(nums) && l2 < len(nums) {
		if nums[l2]%2 == 0 {
			l2++
		} else {
			nums[l1], nums[l2] = nums[l2], nums[l1]
			l2++
		}

		if nums[l1]%2 == 1 {
			l1++
		}
	}
	return nums
}
