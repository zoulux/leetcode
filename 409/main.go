package main

func main() {

}

func longestPalindrome(s string) int {
	mm := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		mm[s[i]]++
	}

	ans := 0
	flag := 0
	for _, v := range mm {
		if v%2 == 0 {
			ans += v
		} else {
			flag = 1
			ans += v - 1
		}
	}

	ans += flag
	return ans
}
