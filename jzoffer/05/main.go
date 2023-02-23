package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceSpace2("We are happy."))
}

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func replaceSpace2(s string) string {
	var sb strings.Builder
	for _, b := range s {
		if b == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteRune(b)
		}
	}

	return sb.String()
}
