package main

import "fmt"

//输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

//输入：target = 9
//输出：[[2,3,4],[4,5]]

func main() {
	fmt.Println(findContinuousSequence(50252))
	//fmt.Println(findContinuousSequence(9))
}

func findContinuousSequence(target int) [][]int {
	var ans [][]int
	var travel func(arr []int, s, idx int) bool
	travel = func(arr []int, s, idx int) bool {
		// s 算求和
		if s == target {
			ans = append(ans, append([]int{}, arr...))
			return true
		}

		if s > target {
			return true
		}

		travel(append(arr, idx), s+idx, idx+1)
		return false
	}
	for i := 1; i < target/2+2; i++ {
		travel([]int{}, 0, i)
	}

	return ans
}
