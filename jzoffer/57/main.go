package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 17))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))

}

func twoSum2(nums []int, target int) []int {
	maxRight := sort.SearchInts(nums, target)

	for i := 0; i < len(nums) && i <= maxRight; i++ {
		x := target - nums[i]
		j := sort.SearchInts(nums[:maxRight], x)
		if j < len(nums) && nums[j] == x {
			return []int{nums[i], nums[j]}
		}
	}
	return nil
}
func twoSum(nums []int, target int) []int {
	maxRight := sort.SearchInts(nums, target)
	if maxRight == len(nums) {
		maxRight--
	}

	left, right := 0, maxRight
	for left <= right {
		sum := nums[left] + nums[right]

		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{
				nums[left], nums[right],
			}
		}
	}
	return nil
}

func twoSum3(nums []int, target int) []int {

	left, right := 0, len(nums)-1
	for left <= right {
		sum := nums[left] + nums[right]

		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{
				nums[left], nums[right],
			}
		}
	}
	return nil
}
