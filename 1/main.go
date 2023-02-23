package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i, v := range nums {

		if j, ok := m[v]; ok {
			return []int{j, i}
		}

		m[target-v] = i

	}

	return []int{0, 0}
}
