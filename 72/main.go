package main

import "fmt"

func main() {

	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i // 如果前面字符串有值，后面没有，那么只能删除 i 次
	}

	for j := range dp[0] {
		dp[0][j] = j // 同理
	}

	min := func(arr ...int) int {
		m := arr[0]
		for _, v := range arr[1:] {
			if v < m {
				m = v
			}
		}
		return m
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j],   //删除
					dp[i][j-1],   //增加
					dp[i-1][j-1], //替换
				) + 1

			}

		}
	}
	return dp[m][n]

}
