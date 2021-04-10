package leetcode

import (
	"math"
	"testing"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverseX := 0
	cx := x
	for cx != 0 {
		reverseX = reverseX*10 + cx%10
		cx = cx / 10
	}
	return reverseX == x
}

func isPalindromeOld(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	count := 1
	cx := x
	for true {
		cx = cx / 10
		if cx == 0 {
			break
		}
		count++
	}
	l := x / int(math.Pow10(count-1))
	r := x % 10
	current := l == r
	sub := int(x-l*int(math.Pow10(count-1))-r) / int(math.Pow10(count-2))
	return current && (isPalindrome(sub))
}

func TestIsPalindrome(t *testing.T) {
	t.Log(isPalindrome(1000021))
	t.Log(isPalindrome(1002001))
	t.Log(isPalindrome(313))
	t.Log(isPalindrome(121))
	t.Log(isPalindrome(10))
	t.Log(isPalindrome(-121))
	t.Log(isPalindrome(0))
}
