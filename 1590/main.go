package main

import "fmt"

func main() {

	//fmt.Println(minSubarray([]int{1, 2, 3}, 3))
	fmt.Println(minSubarray([]int{1, 1, 2, 1, 3}, 6))
}

func minSubarray(nums []int, p int) int {
	n := len(nums)
	lastMod := 0
	for _, v := range nums {
		lastMod += v % p
		lastMod %= p
	}
	if lastMod == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	mm := make(map[int]int) // 余数以及余数对应的索引
	mm[0] = -1
	j := 0
	ans := n

	for i := 0; i < n; i++ {
		j = (j + nums[i]) % p
		k := (j - lastMod + p) % p // 期望的余数
		if v, ok := mm[k]; ok {
			// 存在期望的余数
			ans = min(ans, i-v)
		}
		mm[j] = i
	}
	if ans == n {
		return -1
	}
	return ans
}
