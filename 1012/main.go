package main

import (
	"fmt"
	"strconv"
	"strings"
)

//输入：n = 20
//输出：1
//解释：具有至少 1 位重复数字的正数（<= 20）只有 11 。

func main() {

	//fmt.Println(numDupDigitsAtMostN(11))
	//fmt.Println(numDupDigitsAtMostN(20))
	//fmt.Println(numDupDigitsAtMostN(30))
	//fmt.Println(numDupDigitsAtMostN(100))
	//fmt.Println(numDupDigitsAtMostN(1101))
	//fmt.Println(numDupDigitsAtMostN(29947500))
	//fmt.Println(numDupDigitsAtMostN(10e9))
	fmt.Println(numDupDigitsAtMostN(40137599))
}

func numDupDigitsAtMostN(n int) int {
	if n <= 10 {
		return 0
	}

	dup := func(n int) bool {
		//fmt.Println(n)

		mm := [10]int{}

		for n != 0 {
			x := n % 10
			if mm[x] == 1 {
				return true
			}
			mm[x] = 1
			n /= 10
		}

		return false
	}

	visit := make([]bool, n+1)

	var ans int
	for i := 10; i <= n; i++ {
		if !visit[i] && dup(i) {
			ans++
			m
			for j := i + 1; j <= n; j++ {
				if !visit[j] && strings.Contains(strconv.Itoa(j), strconv.Itoa(i)) {
					ans++
					visit[j] = true
				}
			}
		}
		visit[i] = true
	}
	return ans
}
func numDupDigitsAtMostN2(n int) int {
	if n <= 10 {
		return 0
	}

	dup := func(n int) bool {
		fmt.Println(n)

		mm := [10]int{}

		for n != 0 {
			x := n % 10
			if mm[x] == 1 {
				return true
			}
			mm[x] = 1
			n /= 10
		}

		return false
	}

	visit := make([]bool, n+1)

	var ans int
	for i := 10; i <= n; i++ {
		if !visit[i] && dup(i) {
			ans++

			// 前缀
			t := 10
			for j := i * t; j <= n; j *= t {
				for k := 0; k < t; k++ {
					x := j + k
					if x > n {
						break
					}

					if !visit[x] {
						ans++
					}
					visit[x] = true
				}
				t *= 10
			}
		}
		visit[i] = true
	}
	return ans
}
