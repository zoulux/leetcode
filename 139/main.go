package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordBreak("carcars", []string{
		"car", "ca", "rs",
	}))

	fmt.Println(wordBreak("cars", []string{
		"car", "ca", "rs",
	}))
}

func wordBreak(s string, wordDict []string) bool {
	mm := make(map[string]bool) // 为了快速判断字符床是否存在
	for _, w := range wordDict {
		mm[w] = true
	}

	dp := make([]bool, len(s)+1) // dp[i] 表示字符串 s 的前 i 个字符能否被拆分成 wordDict 中的单词
	dp[0] = true                 // 空字符串默认为 true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && mm[s[j:i]] { // 如果前 j 个字符可以拆分，并且从 j 到 i 的子串在 wordDict 中出现过
				dp[i] = true // 那么前 i 个字符也可以拆分
				break
			}
		}
	}

	return dp[len(s)]
}

func wordBreak2(s string, wordDict []string) bool {

	mm := make(map[string]bool) // 为了快速判断字符床是否存在
	for _, w := range wordDict {
		mm[w] = true
	}

	cache := make(map[string]bool) // 缓存一下，避免重复计算

	var dfs func(s string) bool
	dfs = func(s string) (ret bool) {
		if mm[s] {
			return true
		}

		if v, ok := cache[s]; ok {
			return v
		}
		defer func() {
			cache[s] = ret
		}()

		var ans bool //计算结果，只要本轮有一个 true 那就可以结束了
		for i := 1; i < len(s); i++ {
			if mm[s[:i]] && dfs(s[i:]) { // 看当前字符串在不在，看看后续的字符串在不在
				ans = true // 在的话就停止遍历了
				break
			}
		}
		return ans
	}

	return dfs(s)
}
