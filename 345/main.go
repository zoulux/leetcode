package main

import "bytes"

func main() {

}

func reverseVowels(s string) string {
	ln := len(s)
	mm := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	bs := []byte(s)

	for l, r := 0, ln-1; l < r; {
		for l < ln && !bytes.ContainsRune(mm, rune(s[l])) {
			l++
		}

		for r >= 0 && !bytes.ContainsRune(mm, rune(s[r])) {
			r--
		}

		if l < r {
			bs[l], bs[r] = bs[r], bs[l]
			l++
			r--
		}
	}
	return string(bs)
}
