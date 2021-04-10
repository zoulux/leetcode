package leetcode

import (
	"testing"
)

func isMatch(s string, p string) bool {

	// memo := make(map[string]bool)

	var dp func(i, j int) bool
	dp = func(i, j int) bool {

		if j >= len(p) {
			return i == len(s)
		}

		if i >= len(s) {
			if (len(p)-j)%2 == 1 {
				return false
			}

			for ; j+1 < len(p); j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true
		}

		// key := strconv.Itoa(i) + "," + strconv.Itoa(j)

		// if elem, exist := memo[key]; exist {
		// 	return elem
		// }

		var ret bool

		if s[i] == p[j] || p[j] == '.' {
			if (j+1 < len(p)) && p[j+1] == '*' {
				ret = dp(i, j+2) || dp(i+1, j)
			} else {
				ret = dp(i+1, j+1)
			}
		} else {
			if (j+1 < len(p)) && p[j+1] == '*' {
				ret = dp(i, j+2)
			} else {
				ret = false
			}
		}

		return ret
		// return memo[key]
	}
	return dp(0, 0)
}

func TestIsMatch(t *testing.T) {
	t.Log(isMatch("", "c*c*"))
	// t.Log(isMatch("aa", "a"))
	// t.Log(isMatch("aa", "a*"))
	// t.Log(isMatch("ab", ".*"))
	// t.Log(isMatch("aab", "c*a*b"))
	// t.Log(isMatch("mississippi", "mis*is*p*."))
}
