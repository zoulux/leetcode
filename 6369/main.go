package main

import (
	"fmt"
)

func main() {

	//fmt.Println(maximumOr([]int{
	//	41, 79, 82, 27, 71, 62, 57, 67, 34, 8, 71, 2, 12, 93, 52, 91, 86, 81, 1, 79, 64, 43, 32, 94, 42, 91, 9, 25,
	//}, 10))

	fmt.Println(35 == maximumOr([]int{
		8, 1, 2,
	}, 2))

	fmt.Println(30 == maximumOr([]int{
		12, 9,
	}, 1))
}

func maximumOr(nums []int, k int) int64 {

	suf := make([]int, len(nums)+1)
	pre := make([]int, len(nums)+1)

	for i := len(nums) - 1; i >= 0; i-- {
		suf[i] = suf[i+1] | nums[i]
	}

	//pre[0] = nums[0]
	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] | nums[i]
	}

	fmt.Println(pre, suf)

	var ans int
	for i := 0; i < len(nums); i++ {
		if d := pre[i] | nums[i]<<k | suf[i+1]; d > ans {
			ans = d
		}
	}

	return int64(ans)
}
