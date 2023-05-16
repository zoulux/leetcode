package main

import (
	"fmt"
)

func main() {

	//fmt.Println(strconv.FormatInt(3, 2))
	//fmt.Println(strconv.FormatInt(9, 2))
	//fmt.Println(strconv.FormatInt(11, 2))
	//fmt.Println(strconv.FormatInt(111, 2))
	//fmt.Println(strconv.FormatInt(1111, 2))
	//fmt.Println(strconv.FormatInt(11111, 2))
	//fmt.Println(strconv.FormatInt(111111111, 2))

	//fmt.Println(smallestRepunitDivByK(1))
	//fmt.Println(smallestRepunitDivByK(2))
	//fmt.Println(smallestRepunitDivByK(3))
	fmt.Println(smallestRepunitDivByK(24))
	//fmt.Println(smallestRepunitDivByK(23))
	//fmt.Println(smallestRepunitDivByK(5))
	//fmt.Println(smallestRepunitDivByK(9))
	//fmt.Println(smallestRepunitDivByK(10))
	//fmt.Println(smallestRepunitDivByK(10e5 + 1))
}

const max = 1111111111111111111

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	p := 0
	for i := 1; i <= k; i++ {
		p = (p*10 + 1) % k
		if p == 0 {
			return i
		}
	}
	return -1
}
