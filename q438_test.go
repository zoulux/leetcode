package leetcode

import (
	"testing"
)

func findAnagrams(s, p string) []int {
	need, window := make(map[rune]int), make(map[rune]int)

	for _, c := range p {
		need[c]++
	}

	left, right := 0, 0
	vaild := 0

	ret := []int{}

	for right < len(s) {

		c := rune(s[right])
		right++

		if _, exist := need[c]; exist {
			window[c]++
			if need[c] == window[c] {
				vaild++
			}
		}

		for right-left >= len(p) {
			if vaild == len(need) {
				ret = append(ret, left)
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
	return ret
}

func TestFindAnagrams(t *testing.T) {
	// t.Log(findAnagrams("cbaebabacd", "abc"))
	// t.Log(findAnagrams("abab", "ab"))
	t.Log(findAnagrams("baa", "aa"))
}
