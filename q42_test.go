package leetcode

import (
	"testing"
)

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	LM := height[0]
	RM := height[n-1]
	for i := 0; i < n; i++ {
		if height[i] > LM {
			LM = height[i]
		}

		if height[n-1-i] > RM {
			RM = height[n-1-i]
		}

		leftMax[i] = LM
		rightMax[n-1-i] = RM
	}

	sum := 0
	for i := 0; i < n; i++ {
		c := mmin(leftMax[i], rightMax[i]) - height[i]
		if c > 0 {
			sum += c
		}
	}

	return sum
}

func mmin(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func TestTrap(t *testing.T) {
	t.Log(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	// t.Log(trap2([]int{4, 2, 0, 3, 2, 5}))
	// t.Log(trap2([]int{}))
}

func trap2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	left, right := 0, n-1

	LM := height[left]
	RM := height[right]
	sum := 0

	for left <= right {

		if height[left] > LM {
			LM = height[left]
		}

		if height[right] > RM {
			RM = height[right]
		}

		if LM < RM {
			c := LM - height[left]

			if c > 0 {
				sum += c
			}
			left++
		} else {
			c := RM - height[right]

			if c > 0 {
				sum += c
			}
			right--
		}
	}

	return sum
}
