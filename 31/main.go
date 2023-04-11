package main

import (
	"fmt"
	"sort"
)

func main() {

	arr1 := []int{1, 3, 2}
	nextPermutation(arr1)
	fmt.Println(arr1)

	arr1 = []int{1, 2, 3}
	nextPermutation(arr1)
	fmt.Println(arr1)

	arr2 := []int{3, 2, 1}
	nextPermutation(arr2)
	fmt.Println(arr2)

	arr3 := []int{1, 1, 5}
	nextPermutation(arr3)
	fmt.Println(arr3)
}

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			// 等于找到一个断裂区间
			sort.Ints(nums[i:]) // 将i后面的排序

			for j := i; j < len(nums); j++ {
				if nums[j] > nums[i-1] {
					// 因为排序，所有找到第一个比  nums[i-1] 就退出
					nums[j], nums[i-1] = nums[i-1], nums[j]
					return
				}
			}
		}
	}
	sort.Ints(nums)
}
