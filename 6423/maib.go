package main

import (
	"sort"
)

func main() {

}

func sumOfPower(nums []int) (ans int) {
	const mod int = 1e9 + 7
	sort.Ints(nums)
	s := 0
	for _, x := range nums {
		ans = (ans + x*x%mod*(x+s)) % mod // 中间模一次防止溢出
		s = (s*2 + x) % mod
	}
	return
}
