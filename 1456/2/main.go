package main

import "fmt"

func main() {

	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("aeiou", 2))
}

func maxVowels(s string, k int) int {
	target := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	left, right := 0, 0

	mx := 0
	win := 0
	for right < len(s) {
		for right-left < k {
			if target[s[right]] {
				win++
			}
			right++
		}

		if win > mx {
			mx = win
		}

		if target[s[left]] {
			win--
		}
		left++
	}

	return mx
}
