package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(storeWater([]int{1, 3}, []int{6, 8}))
	//fmt.Println(storeWater([]int{9, 0, 1}, []int{0, 2, 2}))
	fmt.Println(storeWater([]int{3, 2, 5}, []int{0, 0, 0}))

}

func storeWater(bucket []int, vat []int) int {
	n := len(bucket)
	maxk := 0
	for _, v := range vat {
		if v > maxk {
			maxk = v
		}
	}
	if maxk == 0 {
		return 0
	}
	res := math.MaxInt32
	for k := 1; k <= maxk && k < res; k++ {
		t := 0
		for i := 0; i < n; i++ {
			t += max(0, (vat[i]+k-1)/k-bucket[i])
		}
		res = min(res, t+k)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func storeWater2(bucket []int, vat []int) int {

	idx := 0
	var p float64 = 0
	for i := 0; i < len(vat); i++ {
		if bucket[i] != 0 && vat[i] != 0 {
			if t := float64(vat[i]) / float64(bucket[i]); t < p {
				p = t
				idx = i
			}
		}

	}

	pi := int(math.Ceil(p))

	var ans int
	for i, b := range bucket {
		if vat[i] != 0 && pi != 0 {
			if i != idx {
				ans += vat[i]/pi - b
			}
		}

	}

	return ans + pi
}
