package main

import "fmt"

func main() {

	fmt.Println(minReorder(6, [][]int{
		{0, 1},
		{1, 3},
		{2, 3},
		{4, 0},
		{4, 5},
	}))

	//fmt.Println(minReorder(5, [][]int{
	//	{4, 3},
	//	{2, 3},
	//	{1, 2},
	//	{1, 0},
	//}))
}

func minReorder(n int, connections [][]int) int {
	adj := map[int][]int{}
	for _, conn := range connections {
		adj[conn[0]] = append(adj[conn[0]], conn[1])
		adj[conn[1]] = append(adj[conn[1]], -conn[0])
	}
	ans := 0
	var dfs func(u, p int)
	dfs = func(u, p int) {
		for _, v := range adj[u] {
			if v != p && -v != p {
				if v > 0 {
					ans++
				} else {
					v = -v
				}
				dfs(v, u)
			}
		}
	}
	dfs(0, n)
	return ans
}
