package main

import (
	"fmt"
	"sort"
)

// bucketShift returns 1<<b, optimized for code generation.
func bucketShift(b uint8) uintptr {
	// Masking the shift amount allows overflow checks to be elided.
	return uintptr(1) << (b & (4<<(^uintptr(0)>>63)*8 - 1))
}

func main() {
	fmt.Println(11 >> 56)
	m := make(map[string]int, 9000)
	fmt.Println(m)
	fmt.Println(bucketShift(8))

	fmt.Println(102500 & (1024 - 1))

	fmt.Println(findMaxK([]int{
		-1, 2, 3, -3,
	}))
}

func findMaxK(nums []int) int {
	mm := make(map[int]struct{})

	for _, v := range nums {
		mm[v] = struct{}{}
	}

	ans := -1
	for _, v := range nums {
		if _, ok := mm[-v]; ok && v > ans {
			ans = v
		}
	}
	return ans
}

func findMaxK2(nums []int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	sort.Slice(nums, func(i, j int) bool {
		if abs(nums[i]) == abs(nums[j]) {
			return nums[i] < 0
		}
		return abs(nums[i]) > abs(nums[j])
	})

	mm := make(map[int]bool)
	for _, v := range nums {
		if _, ok := mm[-v]; ok {
			return v
		}
		mm[v] = true
	}
	return -1
}
