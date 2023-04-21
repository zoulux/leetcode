package main

import "sort"

func main() {

}

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	var ans int
	for l, r := 0, len(nums)-1; l < r; {
		if nums[l]+nums[r] < k {
			l++
		} else if nums[l]+nums[r] > k {
			r--
		} else {
			ans++ // 找到了
			l++
			r--
		}
	}
	return ans
}
