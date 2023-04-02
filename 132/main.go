package main

import (
	"fmt"
)

func main() {
	//fmt.Println(minCut("ababababababababababababcbabababababababababababa"))
	fmt.Println(minCut("leet"))
	fmt.Println(minCut("ccaacabacb"))
	fmt.Println(minCut("cc|aa|cabac|b"))
	//fmt.Println(minCut("abbab"))
	//fmt.Println(minCut("aab"))
	//fmt.Println(minCut("a"))

}

func minCut(s string) int {

	g := make([][]bool, len(s))
	for i := range g {
		g[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || g[i+1][j-1]) {
				g[i][j] = true
			}
		}
	}

	n := len(s)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, n)

	for i := range dp {
		dp[i] = n
	}

	for i := range dp {
		if g[0][i] {
			dp[i] = 0
		} else {
			for j := 0; j < i; j++ {
				if g[j+1][i] {
					dp[i] = min(dp[i], dp[j]+1)
				}
			}
		}
	}

	return dp[n-1]
}
