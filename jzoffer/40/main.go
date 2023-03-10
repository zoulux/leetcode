package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 1))
	fmt.Println(getLeastNumbers([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4))
}

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func getLeastNumbers2(arr []int, k int) []int {
	for i := 0; i < k; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr[:k]
}
