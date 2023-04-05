package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}

func wordPattern(pattern string, s string) bool {
	parr := strings.Split(pattern, "")
	sarr := strings.Split(s, " ")

	if len(parr) != len(sarr) {
		return false
	}

	for i := 0; i < len(parr); i++ {
		for j := i + 1; j < len(parr); j++ {
			if parr[i] == parr[j] {
				if sarr[i] != sarr[j] {
					return false
				}
			} else {
				if sarr[i] == sarr[j] {
					return false
				}
			}

		}
	}

	return true
}
