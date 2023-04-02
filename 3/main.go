package main

import "fmt"

func main() {

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	win := make(map[byte]int)
	var left, right, ans int

	for right < len(s) {
		c := s[right]
		if l, ok := win[c]; ok && left <= l {
			left = l + 1
		}
		win[c] = right
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
