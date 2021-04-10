package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

func count(n int) string {

	if n == 1 {
		return "1"
	}

	// if n == 2 {
	// 	return "11"
	// }

	// if n == 3 {
	// 	return "21"
	// }

	// if n == 4 {
	// 	return "1211"
	// }

	// if n == 5 {
	// 	return "111221"
	// }

	return say(count(n - 1))
}

func say(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		temp := 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				temp++
				i++
			} else {
				break
			}
		}
		sb.WriteString(strconv.Itoa(temp))
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	return say(count(n - 1))
}

func TestCountAndSay(t *testing.T) {
	t.Log(countAndSay(11))
}

func TestSay(t *testing.T) {

	t.Log(say("1211"))

}
