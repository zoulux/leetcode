package leetcode

import (
	"strings"
	"testing"
)

func lengthOfLongestSubstring2(s string) int {
	window := make(map[rune]int)

	left, right := 0, 0

	ret := 0

	for right < len(s) {

		c := rune(s[right])
		right++
		window[c]++

		for window[c] > 1 {

			d := rune(s[left])
			left++
			window[d]--
		}

		if right-left > ret {
			ret = right - left
		}
	}
	return ret
}

func lengthOfLongestSubstring(s string) int {

	result, left, right := 0, 0, 0

	for right < len(s) {
		ss := s[left:right]
		// 内置函数判断
		if strings.Contains(ss, string(s[right])) {
			left++
		} else {
			right++
		}

		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

func lengthOfLongestSubstringOld(s string) int {

	result, left, right := 0, 0, 0
	var bitSet [256]bool

	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将X标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

func TestLengthOfLongestSubstring(t *testing.T) {

	t.Log(lengthOfLongestSubstring("abcabcbb"))
	t.Log(lengthOfLongestSubstring("bbbbb"))
	t.Log(lengthOfLongestSubstring("pwwkew"))

	t.Log(lengthOfLongestSubstring2("abcabcbb"))
	t.Log(lengthOfLongestSubstring2("bbbbb"))
	t.Log(lengthOfLongestSubstring2("pwwkew"))
}
