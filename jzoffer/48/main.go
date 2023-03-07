package main

// TODO 这题还是不太会

import (
	"fmt"
	"strings"
)

// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
// 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	//fmt.Println(lengthOfLongestSubstring(""))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, len(s)) // 当前索引下最长不含重复字符的子字符串
	ans := 1
	pos := make(map[byte]int) //用哈希表维持的是遍历过程中各个字符最近一次出现的位置
	dp[0] = 1
	pos[s[0]] = 0

	for i := 1; i < len(s); i++ {
		if j, ok := pos[s[i]]; !ok {
			// 不存在重复数字那就 +1
			dp[i] = dp[i-1] + 1
		} else {
			// 存在重复，看一下 当前位置 - 老的位置 差值是否
			if i-j <= dp[i-1] {
				dp[i] = i - j
			} else {
				dp[i] = dp[i-1] + 1
			}
		}
		pos[s[i]] = i
		ans = max(ans, dp[i])
	}
	return ans
}

func lengthOfLongestSubstring3(s string) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := 1

	res := s[:1]
	for i := 1; i < len(s); i++ {
		v := s[i]
		if !strings.ContainsRune(res, rune(v)) {
			ans = max(ans, len(res)+1)
		} else {
			res = res[strings.IndexByte(res, v)+1:]
		}
		res += string(v)
	}
	return ans
}

func lengthOfLongestSubstring2(s string) int {
	win := make(map[byte]int)
	max := 0
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		win[c]++
		right++

		for win[c] > 1 {
			c := s[left]
			win[c]--
			left++
		}

		if x := right - left; x > max {
			max = x
		}
	}
	return max
}
