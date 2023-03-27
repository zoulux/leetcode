package main

import "fmt"

func main() {

	fmt.Println(longestPalindrome("111"))
	fmt.Println(longestPalindrome("11"))
	fmt.Println(longestPalindrome("1121"))
	fmt.Println(longestPalindrome("11211"))
	fmt.Println(longestPalindrome("1221"))
}

func longestPalindrome(s string) string {

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	maxlenth := 0
	left, right := 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
			}

			if dp[i][j] && j-i+1 > maxlenth {
				maxlenth = j - i + 1
				left = i
				right = j
			}
		}
	}

	return s[left : right+1]
}
