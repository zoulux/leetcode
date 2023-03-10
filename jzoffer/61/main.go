package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(isStraight([]int{1, 2, 3, 4, 5}))
	//fmt.Println(isStraight([]int{0, 0, 1, 2, 5}))
	fmt.Println(isStraight([]int{0, 0, 2, 2, 5}))
}

//输入: [1,2,3,4,5]
//输出: True

//输入: [0,0,1,2,5]
//输出: True

func isStraight(nums []int) bool {
	zeroCount := 0
	for _, v := range nums {
		if v == 0 {
			zeroCount++
		}
	}

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 || nums[i-1] == 0 {
			continue
		}

		if nums[i] == nums[i-1] {
			return false
		}

		if x := nums[i] - nums[i-1]; x != 1 {
			if zeroCount > 0 {
				zeroCount -= x
				zeroCount++
				if zeroCount < 0 {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}
