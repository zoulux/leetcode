package leetcode

import "math"

func minWindow(s, t string) string {
	need, window := make(map[rune]int), make(map[rune]int)

	for _, c := range t {
		need[c]++
	}

	left, right := 0, 0
	mleft, mright := 0, math.MaxInt64
	vaild := 0

	for right < len(s) {
		c := rune(s[right])
		right++

		if _, exist := need[c]; exist {
			window[c]++
			if need[c] == window[c] {
				vaild++
			}
		}

		// 找到符合为止
		for vaild == len(need) {
			if right-left < mright-mleft {
				mright = right
				mleft = left
			}

			d := rune(s[left])
			left++

			if _, exist := need[d]; exist {
				if need[d] == window[d] {
					vaild--
				}
				window[d]--
			}
		}
	}

	if mright != math.MaxInt64 {
		return s[mleft:mright]
	}
	return ""
}
