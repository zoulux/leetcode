package main

import "fmt"

func main() {

	fmt.Println(combinationSum3(1, 7))
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
	fmt.Println(combinationSum3(2, 1))
	fmt.Println(combinationSum3(9, 46))
}

func combinationSum3(k int, n int) [][]int {

	var ans [][]int
	var dfs func(arr []int, idx int, s int)
	dfs = func(arr []int, idx int, s int) {
		if len(arr) >= k {
			if s == n {
				ans = append(ans, append([]int{}, arr...))
			}
			return
		}

		if idx == 10 || s > n || n > 45 {
			return
		}

		for i := idx; i <= 9; i++ {
			dfs(append(arr, i), i+1, s+i)
		}
	}

	dfs([]int{}, 1, 0)
	return ans
}
