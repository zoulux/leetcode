package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println()

}

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	mm := make(map[int]struct{})
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		mm[nums[left]+nums[right]] = struct{}{}
	}
	return len(mm)
}

func distinctAverages2(nums []int) int {

	sort.Ints(nums)
	mm := make(map[float64]struct{})
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		k := (float64(nums[left]) + float64(nums[right])) / 2.0
		mm[k] = struct{}{}
	}
	return len(mm)
}
