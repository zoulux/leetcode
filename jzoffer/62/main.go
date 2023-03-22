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

	arr := make([]int, n)
	for i := range arr {
		arr[i] = i
	}

	s := -1
	for len(arr) != 1 {
		ln := len(arr)
		s += m
		if x := s % ln; x != 0 {
			arr = append(arr[:x], arr[x+1:]...)
		}
	}

	return 0
}
