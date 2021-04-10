package leetcode

import (
	"testing"
)

func lengthOfLastWord2(s string) int {

	end := len(s) - 1

	for end >= 0 {
		if s[end] != ' ' {
			break
		}
		end--
	}

	start := end

	for start >= 0 {
		if s[start] == ' ' {
			break
		}
		start--
	}

	return end - start
}

func lengthOfLastWord(s string) int {

	startWord := false
	start, end := -1, -1

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] != ' ' {
			startWord = true
			if end == -1 {
				end = i
			}
		} else {
			if startWord {
				start = i
				break
			}
		}
	}

	return end - start
}

func Test(t *testing.T) {
	// t.Log(lengthOfLastWord2("Hello World"))
	// t.Log(lengthOfLastWord2("Hello World "))
	// t.Log(lengthOfLastWord2(" "))
	// t.Log(lengthOfLastWord2(""))
	t.Log(lengthOfLastWord2("a"))
}
