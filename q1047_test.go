package leetcode

import (
	"strings"
	"testing"
)

func removeDuplicates4(S string) string {

	for i := 0; i+1 < len(S); i++ {
		if S[i] == S[i+1] {
			var sb strings.Builder
			sb.WriteString(S[0:i])
			if i+2 < len(S) {
				sb.WriteString(S[i+2:])
			}
			return removeDuplicates4(sb.String())
		}
	}

	return S
}

func removeDuplicates5(S string) string {

	var sb strings.Builder
	for true {
		sb.Reset()
		target := 0
		flag := false
		for i := 0; i+1 < len(S); i++ {
			if S[i] == S[i+1] {
				sb.WriteString(S[target:i])
				i++
				target = i
				target++
				flag = true
			}
		}
		if target < len(S) {
			sb.WriteString(S[target:])
		}
		S = sb.String()

		if !flag {
			break
		}
	}

	return S
}

func removeDuplicates6(S string) string {
	chars := []byte(S)
	for i := 0; i+1 < len(S); i++ {
		if chars[i] == chars[i+1] {
			if i+2 < len(S) {
				copy(chars[i:], chars[i+2:])
			}
			return removeDuplicates6(string(chars[:len(S)-2]))
		}
	}
	return S
}

func removeDuplicates7(S string) string {
	var helper func(chars []byte, len int) int
	helper = func(chars []byte, len int) int {
		for i := 0; i+1 < len; i++ {
			if chars[i] == chars[i+1] {
				copy(chars[i:len], chars[i+2:len])
				return helper(chars, len-2)
			}
		}
		return len
	}

	chars := []byte(S)
	newlen := helper(chars, len(S))
	return string(chars[0:newlen])
}

func TestRemoveDuplicates6(t *testing.T) {
	t.Log(removeDuplicates6("abbaca"))
	t.Log(removeDuplicates7("abbaca"))
}
