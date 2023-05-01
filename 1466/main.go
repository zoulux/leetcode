package main

import "fmt"

func main() {

	//fmt.Println(minReorder(6, [][]int{
	//	{0, 2}, {0, 3}, {4, 1}, {4, 5}, {5, 0},
	//}))
	//
	//fmt.Println(minReorder(5, [][]int{
	//	{4, 3},
	//	{2, 3},
	//	{1, 2},
	//	{1, 0},
	//}))

	fmt.Println(minReorder(6, [][]int{
		{0, 1},
		{1, 3},
		{2, 3},
		{4, 0},
		{4, 5},
	}))
}

func minReorder(n int, connections [][]int) int {
	type Item struct {
		flag bool
		n    int
	}

	adj := make(map[int][]Item)
	for _, c := range connections {
		adj[c[0]] = append(adj[c[0]], Item{
			flag: true,
			n:    c[1],
		})
		adj[c[1]] = append(adj[c[1]], Item{
			flag: false,
			n:    c[0],
		})
	}

	var ans int
	visit := make([]bool, n)
	var dfs func(i int)
	dfs = func(from int) {
		if visit[from] {
			return
		}
		visit[from] = true
		for _, v := range adj[from] {
			if v.flag && !visit[v.n] {
				// 正向的，并且前面还未访问过
				ans++
			}
			dfs(v.n)
		}
	}
	dfs(0)
	return ans
}
