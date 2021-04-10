package leetcode

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"testing"
)

func reverse2(x int) int {

	value, err := 0, errors.New("")
	if x >= 0 {
		value, err = strconv.Atoi(reverseStr(strconv.Itoa(x)))

	} else {
		value, err = strconv.Atoi(reverseStr(strconv.Itoa(-x)))
		value = (-value)
	}

	if err != nil {
		return 0
	}

	if math.Abs(float64(value)) > float64(math.MaxInt32) {
		return 0
	}

	return value
}

func reverseStr(s string) string {

	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func TestReverse2(t *testing.T) {
	t.Log(reverse2(123))
	t.Log(reverse2(-123))
	t.Log(reverse2(120))
	t.Log(reverse2(0))
	t.Log(reverse2(1534236469))
}
