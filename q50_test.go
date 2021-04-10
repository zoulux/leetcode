package leetcode

import (
	"testing"
)

func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow(1/x, -n)
	}

	if x == 1 || n == 0 {
		return 1
	}

	if n%2 == 0 {
		w := myPow(x, n>>1)
		return w * w
	}

	return x * myPow(x, n-1)
}

func TestMyPow(t *testing.T) {
	// t.Log(myPow(2, 10))
	// t.Log(myPow(2.1, 3))
	// t.Log(myPow(2, -2))
	// t.Log(myPow(0.00001, 2147483647))
	t.Log(myPow(0.00001, 2147483647))

}
