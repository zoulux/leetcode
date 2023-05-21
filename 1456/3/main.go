package main

func main() {

}

var keys = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func maxVowels(s string, k int) int {
	left, right := 0, 0

	mx := 0
	win := 0
	for right < len(s) {
		if keys[s[right]] {
			win++
		}
		right++

		for right-left > k {
			if keys[s[left]] {
				win--
			}
			left++
		}

		if win > mx {
			mx = win
		}
	}
	return mx
}
