package main

import "fmt"

func main() {
	fmt.Println(removeStars("leet**cod*e"))
}

func removeStars(s string) string {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			bs = bs[:len(bs)-1] // 	删除左侧的一个字符
		} else {
			bs = append(bs, s[i]) // 直接插入
		}
	}
	return string(bs)
}
