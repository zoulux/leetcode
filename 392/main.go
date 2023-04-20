package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("", "aaa"))

}

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == len(s)
}

func isSubsequence2(s string, t string) bool {

	x := 0
	if x == len(s) {
		return true
	}

	for i := 0; i < len(t); i++ {
		if s[x] == t[i] {
			// 找到第一个
			x++
			if x == len(s) {
				break
			}
		}
	}
	return x == len(s)

}
