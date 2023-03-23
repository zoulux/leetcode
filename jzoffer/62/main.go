package main

import "fmt"

// 0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。
// 求出这个圆圈里剩下的最后一个数字。
// 例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。

//输入: n = 5, m = 3
//输出: 3

func main() {

	fmt.Println(lastRemaining(5, 3))
}

func lastRemaining(n int, m int) int {
	pos := 0 // 最终活下来那个人的初始位置
	for i := 2; i <= n; i++ {
		pos = (pos + m) % i // 每次循环右移
	}
	return pos
}
