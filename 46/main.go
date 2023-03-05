package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}

func permute(nums []int) [][]int {

	contains := func(nums []int, k int) bool {
		mm := make(map[int]bool)
		for _, v := range nums {
			mm[v] = true
		}
		return mm[k]
	}

	ans := make([][]int, 0)
	var travel func(selectNums []int)
	travel = func(selectNums []int) {
		if len(selectNums) == len(nums) {
			ans = append(ans, selectNums)
			return
		}

		for _, v := range nums {
			if !contains(selectNums, v) {
				travel(append(selectNums, v))
			}
		}
	}

	travel([]int{})

	return ans
}
