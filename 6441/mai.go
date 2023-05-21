package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	//fmt.Println(punishmentNumber(10))
	//fmt.Println(punishmentNumber(37))
	//fmt.Println(punishmentNumber(36))
	fmt.Println(punishmentNumber(45))

	// 2025

	//fmt.Println(handle(45, "2025"))

}

func Init() {
	for i := 38; i <= 1000; i++ {
		if handle(i, strconv.Itoa(i*i)) {
			fmt.Println(i)
		}
	}
}

func handle(i int, i2 string) bool {
	ii, _ := strconv.Atoi(i2)
	if ii == i {
		return true
	}

	for l := 1; l <= len(i2); l++ {
		p := i2[:l]
		pi, _ := strconv.Atoi(p)
		if handle(i-pi, i2[l:]) {
			return true
		}
	}
	return false
}

var arr = []int{
	1,
	9,
	10,
	36,
	45,
	55,
	82,
	91,
	99,
	100,
	235,
	297,
	369,
	370,
	379,
	414,
	657,
	675,
	703,
	756,
	792,
	909,
	918,
	945,
	964,
	990,
	991,
	999,
	1000,
}

func punishmentNumber(n int) int {

	idx := sort.SearchInts(arr, n)
	if arr[idx] != n {
		idx--
	}
	var ans int
	for i := 0; i <= idx; i++ {
		ans += arr[i] * arr[i]
	}

	return ans
}
