package main

import (
	"fmt"
	"strings"
)

func main() {

	//fmt.Println(decodeString("3[a]2[bc]"))
	//fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("3[a]2[bc]"))

}

func decodeString(s string) string {

	if !strings.Contains(s, "[") {
		return s
	}

	var sb strings.Builder
	x := 0
	var sub string
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			x *= 10
			x += int(s[i] - '0')
		} else if s[i] == '[' {
			st := []byte{'['}
			j := i + 1
			for ; j < len(s); j++ {
				if s[j] == '[' {
					st = append(st, '[')
				} else if s[j] == ']' {
					st = st[:len(st)-1]
				}
				if len(st) == 0 {
					// 找到右括号了
					break
				}
			}

			sub = decodeString(s[i+1 : j])

			for i := 0; i < x; i++ {
				sb.WriteString(sub)
			}

			i = j
			sub = ""
			x = 0
		} else {
			sb.WriteByte(s[i])
		}
	}

	return sb.String()
}
