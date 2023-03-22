package main

import "fmt"

func main() {
	fmt.Println(printNumbers(1))
	fmt.Println(printNumbers(2))
	fmt.Println(printNumbers(3))
}

func printNumbers(n int) []int {

	max := 0
	for i := 0; i < n; i++ {
		max = (max * 10) + 9
	}

	var ans []int

	for i := 1; i <= max; i++ {
		ans = append(ans, i)
	}
	return ans
}
