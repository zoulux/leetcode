package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(minNumber([]int{121, 12}))
	fmt.Println(minNumber([]int{10, 2}))
	fmt.Println(minNumber([]int{3, 33, 34, 5, 9}))
	fmt.Println(minNumber([]int{830, 8308})) // 8308 830
	// 830 8308
	// 8308 830

}

func minNumber2(nums []int) string {

	numss := make([]string, len(nums))

	for i, v := range nums {
		numss[i] = strconv.Itoa(v)
	}
	sort.SliceStable(numss, func(i, j int) bool {

		//30  3
		// 121,12*
		// 121 121212
		// 8308 830
		// 830830
		if strings.HasPrefix(numss[i], numss[j]) {
			ln := len(numss[i])
			nj := numss[j]
			return numss[i] < strings.Repeat(nj, ln)

		} else if strings.HasPrefix(numss[j], numss[i]) {
			ln := len(numss[j])
			ni := numss[i]
			return strings.Repeat(ni, ln) < numss[j]
		}
		return numss[i] < numss[j]
	})

	return strings.Join(numss, "")
}
func minNumber(nums []int) string {

	numss := make([]string, len(nums))

	for i, v := range nums {
		numss[i] = strconv.Itoa(v)
	}
	sort.Slice(numss, func(i, j int) bool {
		if strings.HasPrefix(numss[i], numss[j]) || strings.HasPrefix(numss[j], numss[i]) {
			return numss[i]+numss[j] < numss[j]+numss[i]
		}
		return numss[i] < numss[j]
	})

	return strings.Join(numss, "")
}
