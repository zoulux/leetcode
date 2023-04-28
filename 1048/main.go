package main

import "fmt"

func main() {
	fmt.Println(longestStrChain([]string{
		"a", "b", "ba", "bca", "bda", "bdca",
	}))

}

func longestStrChain(words []string) int {

	wm := make(map[string]int)
	for _, w := range words {
		wm[w] = 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(s string) int
	dfs = func(s string) int {
		if v := wm[s]; v > 0 {
			return v
		}

		var ans int
		for i := 0; i < len(s); i++ {
			t := s[:i] + s[i+1:]
			if _, ok := wm[t]; ok {
				ans = max(ans, dfs(t))
			}
		}

		wm[s] = ans + 1
		return ans + 1
	}

	var ans int
	for w := range wm {
		ans = max(ans, dfs(w))
	}
	return ans
}
