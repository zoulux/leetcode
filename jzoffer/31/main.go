package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	i := 0
	for _, pu := range pushed {
		stack = append(stack, pu)
		for len(stack) > 0 {
			if peek := stack[len(stack)-1]; peek == popped[i] {
				stack = stack[:len(stack)-1] // 降 4 弹出
				i++
			} else {
				break
			}
		}
	}
	return i == len(popped)
}
