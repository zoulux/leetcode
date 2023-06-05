package main

import "fmt"

func main() {

	fmt.Println(minimizedStringLength("aaabc"))
	fmt.Println(minimizedStringLength("ipi"))
}
func minimizedStringLength(s string) int {

	mm := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		mm[s[i]] = struct{}{}
	}

	return len(mm)
}

func minimizedStringLength2(s string) int {

	bs := []byte(s)

	for i := 0; i < len(bs); i++ {
		if i+1 < len(bs) && bs[i] == bs[i+1] {
			return minimizedStringLength(string(append(bs[:i], bs[i+1:]...)))
		}
	}

	return len(bs)
}
