package main

import "fmt"

func main() {

	fmt.Println(firstMissingPositive([]int{
		//1, 2, 0,
		//3, 2, 4, -1, 1,
		1, 1,
	}))
}

func firstMissingPositive(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[i]-1 >= 0 {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		t1 := nums[i] // 3
		if t1 != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1

}
func firstMissingPositive2(nums []int) int {

	m := make(map[int]bool)
	min := 1
	for _, v := range nums {
		m[v] = true
		if v == min {
			min++
			for {
				if _, ok := m[min]; ok {
					min++
				} else {
					break
				}
			}
		}
	}
	return min
}
