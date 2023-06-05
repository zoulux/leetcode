package main

import "fmt"

func main() {

	fmt.Println(semiOrderedPermutation([]int{
		2, 1, 4, 3,
	}))

	fmt.Println(semiOrderedPermutation([]int{
		2, 4, 1, 3,
	}))

	fmt.Println(semiOrderedPermutation([]int{
		1, 3, 4, 2, 5,
	}))
}

func semiOrderedPermutation(nums []int) int {
	n := len(nums)
	ii, jj := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			ii = i
		}

		if nums[i] == n {
			jj = i
		}
	}

	if ii < jj {
		return ii + (n - 1 - jj)
	}

	return ii + (n - 1 - jj) - 1
}
