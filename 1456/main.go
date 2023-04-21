package main

import (
	"bytes"
	"fmt"
)

func main() {

	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("tryhard", 4))
}

func maxVowels(s string, k int) int {
	keys := []byte{'a', 'e', 'i', 'o', 'u'}
	mx := 0

	var ans int
	for i := 0; i < k; i++ {
		if bytes.Contains(keys, []byte{s[i]}) {
			ans++
		}
	}

	mx = ans

	for i := 1; i+k <= len(s); i++ {
		if bytes.Contains(keys, []byte{s[i-1]}) {
			ans--
		}
		if bytes.Contains(keys, []byte{s[i+k-1]}) {
			ans++
		}

		if ans > mx {
			mx = ans
		}
	}
	return mx
}
