package main

import "fmt"

func main() {

	//fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("111"))
	//fmt.Println(longestPalindrome("11"))
	fmt.Println(longestPalindrome("1121"))
	fmt.Println(longestPalindrome("11211"))
	fmt.Println(longestPalindrome("1221"))
}

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))

	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	// 以 i 结尾的字符是否是回文串

	left, right := 0, 0
	max := 0

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			// 如果左右指针相等，并且 dp 内部也是回文串
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					dp[i][j] = true
				}
			}

			if d := j - i + 1; dp[i][j] && d > max {
				left = i
				right = j
				max = d
			}
		}
	}
	return s[left : right+1]
}

func longestPalindrome2(s string) string {

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	maxlenth := 0
	left, right := 0, 0

	// 因为子问题是 [i+1][j-1] 是否是回文
	// i 从右向左遍历，确保 i+1 已经被计算过
	// j 从左到右遍历，确保 j-1 已经被计算过
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
