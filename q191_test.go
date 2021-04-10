package leetcode

import (
	"testing"
)

func hammingWeight(num uint32) (ans int) {

	for num != 0 {
		ans += int((num & 1))
		num >>= 1
	}

	return
}

func TestHammingWeight(t *testing.T) {
	a := 0b00000000000000000000000000000011
	t.Log(a & (a - 1))
	// t.Log(hammingWeight(00000000000000000000000000001011))
	t.Log(hammingWeight(0b00000000000000000000000000000001))
	t.Log(hammingWeight(0b00000000000000000000000000000010))
	t.Log(hammingWeight(0b00000000000000000000000000000011))
}
