package leetcode

import (
	"math"
	"testing"
)

func minEatingSpeed(piles []int, h int) int {

	min := math.MaxInt64
	max := math.MinInt64
	for _, p := range piles {
		if p > max {
			max = p
		}

		if p < min {
			min = p
		}
	}

	for k := min; k <= max; k++ {
		h_ := 0
		for p := range piles {
			h_ += (p / k)
		}

		if h == h_ {
			return k
		}

	}

	return 0
}

func TestMinEatingSpeed(t *testing.T) {
	t.Log(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}
