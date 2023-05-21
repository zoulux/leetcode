package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 0
	left, right := 0, 0
	win := make(map[byte]int)
	for ; right < len(s); right++ {
		win[s[right]]++
		for win[s[right]] > 1 {
			win[s[left]]--
			left++
		}
		if t := right - left; t > max {
			max = t
		}
	}
	return max + 1
}
