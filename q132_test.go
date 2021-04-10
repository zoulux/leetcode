package leetcode

import (
	"math"
	"testing"
)

func minCut(s string) int {
	n := len(s)

	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
		for j := range g[i] {
			g[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			g[i][j] = s[i] == s[j] && g[i+1][j-1]
		}
	}

	dp := make([]int, n)
	// 表示前缀 s[0:i] 最少分割次数
	// dp[n-1] 就是最后结果

	for i := range dp {
		if g[0][i] {
			continue
		}
		dp[i] = math.MaxInt64
		//  s[j+1:i] 如果是回文
		for j := 0; j < i; j++ {
			if g[j+1][i] && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	return dp[n-1]
}

func TestMinCut(t *testing.T) {
	// t.Log(minCut("aab"))
	// t.Log(minCut("a"))
	// t.Log(minCut("ab"))
	// t.Log(minCut("ababababababababababababcbabababababababababababa"))
	// t.Log(minCut("fff"))
	t.Log(minCut("aaabaa"))
	t.Log(minCut("eegiicgaeadbcfacfhifdbiehbgejcaeggcgbahfcajfhjjdgj"))
}
