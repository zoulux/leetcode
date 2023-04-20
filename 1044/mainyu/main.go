package main

import "fmt"

func main() {

	fmt.Println(longestDupSubstring("aa"))
}

func longestDupSubstring(s string) string {
	n := len(s)
	l, r := 0, n-1
	ans := ""
	for l < r {
		mid := l + (r-l+1)>>1
		//二分长度
		if res, ok := longestDupSubstringCheck(s, mid); ok {
			ans = res
			l = mid
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 使用map记录子串是否出现过
func longestDupSubstringCheck(s string, length int) (string, bool) {
	cache := map[string]struct{}{}
	for i := 0; i <= len(s)-length; i++ {
		subStr := s[i : i+length]
		if _, ok := cache[subStr]; ok {
			return subStr, true
		}
		cache[subStr] = struct{}{}
	}
	return "", false
}
