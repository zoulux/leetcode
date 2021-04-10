package leetcode

import "testing"

func mySqrt(x int) int {

	left, right := 0, x/2+1

	for left <= right {
		mid := (right-left)/2 + left
		tmp := mid * mid
		if tmp == x {
			return mid
		} else if tmp > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func TestMySqrt(t *testing.T) {
	t.Log(mySqrt(0))

	// t.Log(mySqrt(5))
	// t.Log(mySqrt(8))
	// t.Log(mySqrt(4))
	// t.Log(mySqrt(9))
	// t.Log(mySqrt(8))
	// t.Log(mySqrt(100000))
	// t.Log(mySqrt(1))
	// t.Log(mySqrt(5))
}
