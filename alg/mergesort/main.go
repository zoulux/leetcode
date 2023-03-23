package main

import (
	"fmt"
	"math/rand"
)

func main() {

	arr := rand.Perm(11)
	fmt.Println(arr)
	arr = mergeSort(arr)
	fmt.Println(arr)

}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	left, right := arr[:middle], arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}
