package leetcode

import (
	"testing"
)

func checkInclusion2(s1, s2 string) bool {
	// need, window := make([]int, 26), make([]int, 26)
	need, window := [26]int{}, [26]int{}

	for _, c := range s1 {
		need[c-'a']++
	}

	size := 0

	for _, c := range need {
		if c != 0 {
			size++
		}
	}

	left, right := 0, 0
	vaild := 0

	for right < len(s2) {
		c := s2[right] - 'a'
		right++

		if need[c] != 0 {
			window[c]++
			if need[c] == window[c] {
				vaild++
			}
		}

		for right-left >= len(s1) { // 是否开始移动左侧
			if vaild == size {
				return true
			}

			d := s2[left] - 'a'
			left++

			if need[d] != 0 {
				if need[d] == window[d] {
					vaild--
				}
				window[d]--
			}
		}
	}
	return false
}

func checkInclusion(s1, s2 string) bool {
	need, window := make(map[rune]int), make(map[rune]int)

	for _, c := range s1 {
		need[c]++
	}

	left, right := 0, 0
	vaild := 0

	for right < len(s2) {
		c := rune(s2[right])
		right++

		if _, exist := need[c]; exist {
			window[c]++
			if need[c] == window[c] {
				vaild++
			}
		}

		for right-left >= len(s1) { // 是否开始移动左侧
			if vaild == len(need) {
				return true
			}

			d := rune(s2[left])
			left++

			if _, exist := need[d]; exist {
				if need[d] == window[d] {
					vaild--
				}
				window[d]--
			}
		}
	}
	return false
}

func TestCheckInclusion(t *testing.T) {
	t.Log(checkInclusion("abo", "eidbaooo"))
	t.Log(checkInclusion2("abo", "eidbaooo"))
}
