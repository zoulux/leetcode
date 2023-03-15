package main

import (
	"fmt"
)

func main() {

	fmt.Println(singleNumber([]int{1, 1, 6, 1}))
}

func singleNumber(nums []int) int {
	var ans int
	for i := 0; i < 32; i++ {
		x := 0
		for j, v := range nums {
			if v != 0 {
				x += v & 1
				nums[j] = v >> 1
			}
		}
		ans += (x % 3) << i
	}
	return ans
}
