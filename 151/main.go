package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	arr := strings.Fields(s)
	for l, r := 0, len(arr)-1; l < r; {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return strings.Join(arr, " ")
}
