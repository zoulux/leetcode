package main

import "fmt"

func main() {

	x := []int{
		1, 0,
	}
	moveZeroes(x)
	fmt.Println(x)

}

func moveZeroes(nums []int) {
	for left, right := 0, 0; left < len(nums) && right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
