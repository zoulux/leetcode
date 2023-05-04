package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(minDistance("", ""))
	fmt.Println(minDistance("", "ros"))
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	min := func(arr ...int) int {
		sort.Ints(arr)
		return arr[0]
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if word1[i] != word2[j] {
				dp[i+1][j+1] = min(
					dp[i+1][j],
					dp[i][j+1],
					dp[i][j],
				) + 1
			} else {
				dp[i+1][j+1] = dp[i][j]
			}
		}
	}

	return dp[n][m]
}
