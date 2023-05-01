package main

import "fmt"

func main() {

	fmt.Println(numOfMinutes(6, 2, []int{
		2, 2, -1, 2, 2, 2,
	}, []int{
		0, 0, 1, 0, 0, 0,
	}))
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(root int) int
	dfs = func(root int) int {
		if headID == root {
			return informTime[root]
		}

		return dfs(manager[root]) + informTime[root]
	}

	var ans int
	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i))
	}
	return ans
}

func numOfMinutes2(n int, headID int, manager []int, informTime []int) int {

	mm := make(map[int][]int)

	for i, m := range manager {
		if m >= 0 {
			mm[m] = append(mm[m], i) // manage 下面的儿子
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(root int) int
	dfs = func(root int) int {
		var ch int
		for _, v := range mm[root] {
			ch = max(ch, dfs(v))
		}
		return ch + informTime[root]
	}
	return dfs(headID)
}
