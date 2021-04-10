package leetcode

import (
	"testing"
)

func isPalindrome2(c []byte) bool {

	i, j := 0, len(c)-1
	for i <= j {
		if c[i] != c[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func shortestPalindrome(s string) string {

	n := len(s)
	cs := make([]byte, n*2)
	copy(cs[n:], []byte(s))

	for i := 0; i < n; i++ {
		cs[i] = s[n-i-1]
	}

	min := cs
	for i := 1; i < n; i++ {
		if isPalindrome2(cs[i:]) {
			min = cs[i:]
		}
	}

	return string(min)
}

func TestShortestPalindrome(t *testing.T) {
	t.Log(shortestPalindrome("abca"))
	t.Log(shortestPalindrome("aacecaaa"))
	t.Log(shortestPalindrome("abcd"))
}
