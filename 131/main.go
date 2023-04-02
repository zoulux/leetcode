package main

import "fmt"

func main() {
	fmt.Println(partition("ababbbabbaba"))
	fmt.Println(partition("abbab"))
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))

}

func partition(s string) [][]string {

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}

	var ans [][]string
	var travel func(ret []string, i int)
	travel = func(ret []string, i int) {
		if i == len(s) {
			ans = append(ans, append([]string{}, ret...))
			return
		}

		for j := i; j < len(s); j++ {
			if dp[i][j] {
				travel(append(ret, s[i:j+1]), j+1)
			}
		}
	}

	travel([]string{}, 0)
	return ans
}
