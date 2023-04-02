package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
}

func reversePairs(nums []int) int {

	var ls []int
	var ans int
	for _, v := range nums {
		idx := sort.SearchInts(ls, v)
		left := ls[:idx]
		right := ls[idx:]
		ls = append([]int{}, left...)
		ls = append(ls, v)
		ls = append(ls, right...)
		ans += len(ls) - 1 - idx
	}
	return ans
}

func reversePairs2(nums []int) int {

	count := 0
	merge := func(a, b []int) []int {
		result := make([]int, 0, len(a)+len(b))

		for len(a) != 0 && len(b) != 0 {
			if a[0] < b[0] {
				result = append(result, a[0])
				a = a[1:]
			} else {
				result = append(result, b[0])
				b = b[1:]
				count++
			}
		}

		result = append(result, a...)
		result = append(result, b...)
		return result
	}

	var mergetSort func(arr []int) []int

	mergetSort = func(arr []int) []int {
		if len(arr) < 2 {
			return arr
		}

		mid := len(arr) / 2
		left := arr[:mid]
		right := arr[mid:]
		return merge(mergetSort(left), mergetSort(right))
	}

	fmt.Println(mergetSort(nums))

	return count

}
