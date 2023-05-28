package main

import "fmt"

func main() {
	fmt.Println(jump([]int{
		1, 2,
	}))

	fmt.Println(jump([]int{
		0,
	}))

	fmt.Println(jump([]int{
		2, 3, 1, 1, 4,
	}))
	fmt.Println(jump([]int{
		2, 3, 0, 1, 4,
	}))
}

func jump(nums []int) int {
	step := 0  // 步数
	right := 0 // 最远达到的距离
	end := 0   //
	for i, n := range nums[:len(nums)-1] {
		if t := n + i; t > right {
			right = t
		}

		if i == end {
			step++
			end = right
		}
	}

	return step
}
