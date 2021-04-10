package leetcode

import (
	_ "unsafe"

	"testing"
)

//go:linkname formatBits strconv.formatBits
func formatBits(dst []byte, u uint64, base int, neg, append_ bool) (d []byte, s string)

func reverseBits(num uint32) uint32 {
	var ans uint32 = 0
	for i := 0; i < 32; i++ {
		//最终结果左移一位
		ans <<= 1
		//取最低位
		low := num & 1
		//将最低位放到rs中
		ans += low
		//移除[溢出]最低位
		num >>= 1
	}
	return ans
}

func TestReverseBits(t *testing.T) {

	t.Log(reverseBits(43261596))
}
