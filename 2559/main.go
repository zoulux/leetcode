package main

import (
	"fmt"
)

func main() {

	fmt.Println(vowelStrings([]string{
		"aba", "bcb", "ece", "aa", "e",
	}, [][]int{
		{0, 2},
		{1, 4},
		{1, 1},
	}))
}

func vowelStrings(words []string, queries [][]int) []int {
	//'a'、'e'、'i'、'o' 和 'u' 。

	equal := func(b byte) bool {
		return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
	}

	pre := make([]int, len(words)+1)

	for i, w := range words {
		pre[i+1] = pre[i]
		if equal(w[0]) && equal(w[len(w)-1]) {
			pre[i+1]++
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = pre[q[1]+1] - pre[q[0]]
	}
	return ans
}
