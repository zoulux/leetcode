package main

import "fmt"

func main() {

	fmt.Println(maxArea([]int{
		1, 8, 6, 2, 5, 4, 8, 3, 7,
	}))

}

func maxArea(height []int) int {

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var ans int
	left, right := 0, len(height)-1
	for left < right {
		h := min(height[left], height[right])
		w := right - left
		ans = max(ans, h*w)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ans
}

func maxArea2(height []int) int {
	leftdp := make([][2]int, len(height))
	rightdp := make([][2]int, len(height))

	for i := 0; i < len(height); i++ {
		leftdp[i] = [2]int{height[i], i}
		if i-1 >= 0 && leftdp[i-1][0] > leftdp[i][0] {
			leftdp[i] = leftdp[i-1]
		}
	}

	for i := len(height) - 1; i >= 0; i-- {
		rightdp[i] = [2]int{height[i], i}
		if i+1 < len(height) && rightdp[i+1][0] > rightdp[i][0] {
			rightdp[i] = rightdp[i+1]
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	mx := 0
	for i := 0; i < len(height); i++ {
		h := min(leftdp[i][0], rightdp[i][0])
		w := rightdp[i][1] - leftdp[i][1]

		mx = max(mx, w*h)
		fmt.Println(w, h, mx)

	}
	return mx
}
