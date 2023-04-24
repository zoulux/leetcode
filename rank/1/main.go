package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(supplyWagon([]int{
		7, 3, 6, 1, 8,
	}))

	fmt.Println(supplyWagon([]int{
		6, 2, 2, 6, 9, 8, 5, 7,
	}))

}

func supplyWagon(arr []int) []int {
	ln := len(arr)
	b := ln / 2

	var _supplyWagon func(supplies []int) []int
	_supplyWagon = func(supplies []int) []int {
		if len(supplies) == b {
			return supplies
		}

		min := math.MaxInt / 2
		li := 0
		for i := 0; i+1 < len(supplies); i++ {
			if d := supplies[i] + supplies[i+1]; d < min {
				min = d
				li = i
			}
		}
		l := supplies[:li]

		var r []int
		if li+2 < len(supplies) {
			r = supplies[li+2:]
		}

		ans := append([]int{}, l...)
		ans = append(ans, supplies[li]+supplies[li+1])
		ans = append(ans, r...)
		return _supplyWagon(ans)
	}
	return _supplyWagon(arr)
}
