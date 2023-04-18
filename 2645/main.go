package main

import "fmt"

func main() {
	fmt.Println(addMinimum("b"))
	fmt.Println(addMinimum("bca"))
}

func addMinimum(word string) int {
	t := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] >= word[i] {
			// 只要有一个逆序，就需要单独组一个 abc
			// ba aa 都算逆序
			t++
		}
	}
	return t*3 - len(word)
}
