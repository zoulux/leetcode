package main

import (
	"fmt"
)

// 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。

// 提示：
// a, b 均可能是负数或 0
// 结果不会溢出 32 位整数

func main() {

	fmt.Println(add(3, 6))
}

func add(a int, b int) int {
	if b == 0 {
		return a
	}
	return add(a^b, (a&b)<<1)
}
