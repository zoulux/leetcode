package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(cuttingRope(120))
	//fmt.Println(cuttingRope(1000))
}

func cuttingRope(n int) int {
	if n < 3 {
		return 1
	}
	if n == 3 {
		return 2
	}

	res := 1
	var m int = 1e9 + 7
	for n > 4 {
		res *= 3
		res %= m
		n -= 3
	}
	return (n * res) % m
}

func cuttingRope2(n int) int {

	max := func(a, b *big.Int) *big.Int {
		if a != nil && a.Cmp(b) == 1 {
			return a
		}
		return b
	}

	dp := make([]*big.Int, n+1)

	dp[0] = big.NewInt(1)
	dp[1] = big.NewInt(1)
	dp[2] = big.NewInt(1)

	var m int64 = 1e9 + 7

	for i := int64(3); i <= int64(n); i++ {

		for j := int64(1); j <= i/2; j++ {
			//---
			// - --

			// ----
			// - ---
			// -- --

			// -----
			// - ----
			// -- ---

			// 对于某一刀，要么直接相乘是最大的，要么是当前值乘以右边dp结果是最大的
			t1 := big.NewInt(1).Mul(big.NewInt(j), big.NewInt(i-j))
			t2 := big.NewInt(1).Mul(big.NewInt(j), dp[i-j])

			dp[i] = max(dp[i], max(t1, t2))
		}

	}

	return int(big.NewInt(1).Mod(dp[n], big.NewInt(m)).Int64())
}
