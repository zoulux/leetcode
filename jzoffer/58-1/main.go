package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	arr := strings.Fields(s)
	var sb strings.Builder
	for i := len(arr) - 1; i >= 0; i-- {
		sb.WriteString(arr[i])
		if i != 0 {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}

func reverseWords2(s string) string {
	arr := strings.Fields(s)
	ln := len(arr)
	for i := 0; i < ln/2; i++ {
		arr[i], arr[ln-i-1] = arr[ln-i-1], arr[i]
	}
	return strings.Join(arr, " ")
}
