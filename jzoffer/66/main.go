package main

import "fmt"

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
	fmt.Println(constructArr([]int{1, 0, 3, 4, 5}))
}

func constructArr(a []int) []int {
	ans := make([]int, len(a))

	cur := 1
	for i := 0; i < len(a); i++ {
		ans[i] = cur
		cur *= a[i]
	}

	cur = 1
	for i := len(a) - 1; i >= 0; i-- {
		ans[i] *= cur
		cur *= a[i]
	}
	//ans[0] = curm
	return ans
}

// 会超时
func constructArr2(a []int) []int {
	left := []int{1}
	right := []int{1}

	for i := 0; i < len(a); i++ {
		left = append(left, left[len(left)-1]*a[i])
	}

	for i := len(a) - 1; i >= 0; i-- {
		right = append([]int{right[0] * a[i]}, right...)
	}

	for i := 0; i < len(a); i++ {
		a[i] = left[i] * right[i+1]
	}
	return a
}
