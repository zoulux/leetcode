package main

import "fmt"

func main() {

	fmt.Println(dailyTemperatures([]int{
		73, 74, 75, 71, 69, 72, 76, 73,
	}))
}

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 1 {
		return []int{0}
	}

	ans := make([]int, len(temperatures))
	//stack := make([]int, 0)
	var stack []int
	//stack = append(stack, 0)
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			ln := len(stack) - 1
			ans[stack[ln]] = i - stack[ln]
			stack = stack[:ln]
		}
		stack = append(stack, i)
	}
	return ans
}

func dailyTemperatures2(temperatures []int) []int {
	if len(temperatures) == 1 {
		return []int{0}
	}

	ans := make([]int, len(temperatures))
	stack := make([][2]int, 0)
	stack = append(stack, [2]int{0, temperatures[0]})
	for i := 1; i < len(temperatures); i++ {
		for len(stack) > 0 && stack[len(stack)-1][1] < temperatures[i] {
			ans[stack[len(stack)-1][0]] = i - stack[len(stack)-1][0]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, [2]int{i, temperatures[i]})
	}
	return ans
}
