package main

import (
	"fmt"
	"math/bits"
)

func main() {

	fmt.Println(countBits(2))

}

func countBits(n int) []int {

	//countBit := func(n int) int {
	//	return bits.OnesCount(uint(n))

	//ans := 0
	//for i := 0; i < 32; i++ {
	//	if n != 0 {
	//		if 1&n == 1 {
	//			ans++
	//		}
	//		n >>= 1
	//	}
	//}
	//return ans
	//}

	var ans []int
	for i := 0; i <= n; i++ {
		ans = append(ans, bits.OnesCount(uint(i)))
	}
	return ans
}
