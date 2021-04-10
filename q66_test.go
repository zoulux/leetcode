package leetcode

import "testing"

func plusOne(digits []int) []int {

	plus := 1

	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + plus
		if tmp >= 10 {
			digits[i] = tmp - 10
			plus = 1
		} else {
			digits[i] = tmp
			plus = 0
			break
		}
	}
	if plus != 0 {
		digits = append([]int{plus}, digits...)
	}

	return digits
}
func TestPlusOne(t *testing.T) {
	t.Log(plusOne([]int{9}))
}
