package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(minArray([]int{4, 5, 1, 2, 3}))
	fmt.Println(minArray([]int{1, 1, 1, 1}))
}

func minArray2(numbers []int) int {
	minFun := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	ans := math.MaxInt
	for _, v := range numbers {
		ans = minFun(ans, v)
	}
	return ans
}

func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		mid := low + (high-low)/2
		if numbers[mid] < numbers[high] {
			high = mid
		} else if numbers[mid] > numbers[high] {
			low = mid + 1
		} else {
			high -= 1
		}
	}
	return numbers[low]

}
