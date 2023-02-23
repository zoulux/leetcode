package main

import "fmt"

func main() {

	fmt.Println(lengthOfLongestSubstring("abcabcbb")) //3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    //1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}

func lengthOfLongestSubstring(s string) int {
	l, r, max, win := 0, 0, 0, make(map[byte]int)
	for r < len(s) {
		c := s[r]
		r++
		win[c]++

		for win[c] > 1 {
			d := s[l]
			l++
			win[d]--
		}

		if r-l > max {
			max = r - l
		}
	}
	return max
}
