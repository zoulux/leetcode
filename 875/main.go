package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(minEatingSpeed([]int{
		30, 11, 23, 4, 20,
	}, 6))

	fmt.Println(minEatingSpeed([]int{
		3,
	}, 8))

	fmt.Println(minEatingSpeed([]int{
		3, 6, 7, 11,
	}, 8))

	fmt.Println(minEatingSpeed([]int{
		30, 11, 23, 4, 20,
	}, 5))

}

func minEatingSpeed(piles []int, h int) int {
	left := piles[0]
	right := piles[0]
	for _, p := range piles[1:] {
		if p < left {
			left = p
		}
		if p > right {
			right = p
		}
	}

	for left <= right {
		k := float64(left+right) / 2.0

		var t int
		for _, v := range piles {
			t += int(math.Ceil(float64(v) / k))
		}
		if t < h {
			right = int(k - 1)
		} else if t > h {
			left = int(k + 1)
		} else {
			return int(k)
		}
	}
	return 0
}
