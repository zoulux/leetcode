package leetcode

import (
	"testing"
)

func multiply(num1 string, num2 string) string {

	if len(num2) == 1 {

	}

	multiply(num1, num2[:len(num2)-1])

	return ""
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func TestMultiply(t *testing.T) {

	// t.Log(multiply("2", "3"))
	t.Log(multiply("123", "456"))
}
