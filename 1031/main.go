package main

import "fmt"

func main() {

	fmt.Println(maxSumTwoNoOverlap([]int{
		2, 1, 5, 6, 0, 9, 5, 0, 3, 8,
	}, 4, 2))

}

// 0,6,5,2,2,5,1,9,4
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	s := make([]int, n+1)
	for i, v := range nums {
		s[i+1] = s[i] + v
	}

	f := func(firstLen, secondLen int) int {
		var ans int
		maxSumA := 0
		for i := firstLen + secondLen; i <= n; i++ {
			maxSumA = max(maxSumA, s[i-secondLen]-s[i-secondLen-firstLen]) //  s[i-secondLen]-s[i-secondLen-firstLen] 为 a 数组的和
			ans = max(ans, maxSumA+s[i]-s[i-secondLen])                    // 为 s[i]-s[i-secondLen] 为b 数组的和
		}
		return ans
	}

	return max(
		f(firstLen, secondLen),
		f(secondLen, firstLen),
	)
}
