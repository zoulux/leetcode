package leetcode

import (
	"math"
	"testing"
)

func divide(dividend int, divisor int) int {

	isFu := false
	if dividend > 0 && divisor < 0 {
		isFu = true
	} else if dividend < 0 && divisor > 0 {
		isFu = true
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	i := 0

	if divisor == 1 {
		i = dividend
	} else {
		for dividend > 0 && dividend >= divisor {
			ant := 1
			for dividend > 0 && dividend >= divisor {
				dividend = dividend >> 1
				ant=ant>>1
			dividend -= divisor

			i++


			}
		}
	}

	if isFu {
		if -i < math.MinInt32 {
			return -i + 1
		}

		return -i
	}
	if i > math.MaxInt32 {
		return i - 1
	}
	return i
}

func TestDivide(t *testing.T) {
	t.Log(divide(-5, -1))

	t.Log(divide(10, 2))
	t.Log(divide(10, 3))
	t.Log(divide(10, -3))
	t.Log(divide(7, -3))
	t.Log(divide(-7, -3))
	// t.Log(divide(-2147483648, -1))
	t.Log(divide(-2147483648, -1))
	t.Log(divide(-2147483648, -2))
	t.Log(divide(-2147483648, -3))

}
