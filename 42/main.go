package main

import "fmt"

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{0}))
}

func trap(height []int) int {
	n := len(height)
	dpleft := make([]int, len(height))
	dpright := make([]int, len(height))

	dpleft[0] = height[0]
	for i := 1; i < len(height); i++ {
		dpleft[i] = max(height[i], dpleft[i-1])
	}

	dpright[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		dpright[i] = max(height[i], dpright[i+1])
	}

	var ans int
	for i := 0; i < n; i++ {
		ans += min(dpleft[i], dpright[i]) - height[i]
	}
	return ans
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
