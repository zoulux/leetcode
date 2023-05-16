package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	fmt.Println(matrixSum([][]int{
		{7, 2, 1},
		{6, 4, 2},
		{6, 5, 3},
		{3, 2, 1},
	}))

	fmt.Println(matrixSum([][]int{{}}))

}

func matrixSum(nums [][]int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums[0]) == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		sort.Ints(nums[i]) // 对于每一行排序
	}

	var ans int
	for j := 0; j < len(nums[0]); j++ {
		t := 0
		for i := 1; i < len(nums); i++ {
			if nums[i][j] > nums[t][j] {
				t = i
			}
		}
		ans += nums[t][j]
	}

	return ans
}

func matrixSum2(nums [][]int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums[0]) == 0 {
		return 0
	}

	max := math.MinInt / 2
	for i := 0; i < len(nums); i++ {
		t := 0
		for j := 1; j < len(nums[i]); j++ {
			if nums[i][j] > nums[i][t] {
				t = j
			}
		}

		if nums[i][t] > max {
			max = nums[i][t]
		}

		nums[i] = append(nums[i][:t], nums[i][t+1:]...)
	}
	return max + matrixSum(nums)
}
