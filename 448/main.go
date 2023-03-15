package main

import (
	"fmt"
	"sort"
)

//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。
//请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[5,6]

func main() {

	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {

	sort.Ints(nums)
	ans := make([]int, 0)

	j := 0

	for i := 0; i < len(nums); i++ {
		j++
		if nums[i] != j {
			ans = append(ans, nums[i])
			//j++
		}
	}
	return ans
}

func findDisappearedNumbers2(nums []int) []int {
	mm := make(map[int]int)
	for _, v := range nums {
		mm[v]++
	}

	ans := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		if _, ok := mm[i]; !ok {
			ans = append(ans, i)
		}
	}
	return ans
}
