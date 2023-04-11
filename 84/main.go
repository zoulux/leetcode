package main

import "fmt"

func main() {

	fmt.Println(largestRectangleArea([]int{
		2, 1, 5, 6, 2, 3,
	}))

}

func largestRectangleArea(heights []int) int {

	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	var stack []int

	max := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			// 确定了右边界
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]

			if d := (i - 1 - left) * h; d > max {
				max = d
			}
		}
		stack = append(stack, i)
	}
	return max
}
