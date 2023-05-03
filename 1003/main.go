package main

import "fmt"

func main() {

	fmt.Println(isValid("aabcbc"))

}

func isValid(s string) bool {
	st := make([]rune, 0)
	for _, v := range s {
		if len(st) < 2 {
			st = append(st, v) // 长度小于 2 ，直接塞进去
		} else if v == 'c' && st[len(st)-1] == 'b' && st[len(st)-2] == 'a' {
			// 如果当前字符为 c，前面有个 ab，则出站
			st = st[:len(st)-2]
		} else {
			// 还是塞进去
			st = append(st, v)
		}
	}
	return len(st) == 0
}
