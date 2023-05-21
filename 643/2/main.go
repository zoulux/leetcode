package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(findMaxAverage([]int{
		1, 12, -5, -6, 50, 3,
	}, 4))

	fmt.Println(findMaxAverage([]int{
		-1,
	}, 1))

}

func findMaxAverage(nums []int, k int) float64 {

	var ans float64
	ans = math.MinInt / 2
	var win float64
	right := 0
	left := 0
	for right < len(nums) {
		for right-left < k {
			win += float64(nums[right])
			right++
		}

		if t := win / float64(k); t > ans {
			ans = t
		}
		win -= float64(nums[left])
		left++
	}

	return ans
}
