package main

import "fmt"

func main() {

	x := []int{1, 2, 0, 2, 1, 1, 0}
	sortColors(x)
	fmt.Println(x)
}

func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				// p1 也需要向后挪
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

func sortColors3(nums []int) {
	target := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[target] = nums[target], nums[i]
			target++
		}
	}

	for i := target; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[i], nums[target] = nums[target], nums[i]
			target++
		}
	}
}

func sortColors2(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
