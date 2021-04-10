package leetcode

import (
	"math"
	"testing"
)

func myAtoi(s string) int {
	maxInt32, minInt32 := 1<<31-1, -1<<31
	isPlus := 1
	isPlusFlag := false

	for len(s) > 0 {
		c := s[0]
		if (c == ' ') || (c == '+') || (c == '-') {

			if isPlusFlag {
				return 0
			}

			if c == '-' {
				isPlus = -1
				isPlusFlag = true
			}

			if c == '+' {
				isPlus = 1
				isPlusFlag = true
			}

			s = s[1:]
		} else {
			break
		}
	}

	ret := 0
	for i := 0; i < len(s); i++ {

		if s[i] >= '0' && s[i] <= '9' {
			// 数字
			ret = ret*10 + int(s[i]-'0')

			if ret*isPlus >= maxInt32 {
				return math.MaxInt32
			}

			if ret*isPlus <= minInt32 {
				return minInt32
			}

		} else {
			// 非数字
			break
		}
	}
	return ret * isPlus
}

func TestMyAtoi(t *testing.T) {
	t.Log(myAtoi("   -42"))
	t.Log(myAtoi("42"))
	t.Log(myAtoi("4193 with words"))
	t.Log(myAtoi("words and 987"))
	t.Log(myAtoi("-91283472332"))
}
