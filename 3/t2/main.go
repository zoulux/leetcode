package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
	fmt.Println(lengthOfLongestSubstring("12"))
	fmt.Println(lengthOfLongestSubstring("1"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("bbbbb"))
	//fmt.Println(lengthOfLongestSubstring("pwwkew"))

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left, right := 0, 0
	mm := make(map[byte]int)

	max := 0
	for right < len(s) {
		if v, ok := mm[s[right]]; ok && left <= v {
			// 表示存在
			left = v + 1 // 那么
		}
		mm[s[right]] = right

		if d := right - left + 1; d > max {
			max = d
		}
		right++
	}
	return max
}
