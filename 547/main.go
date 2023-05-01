package main

import "fmt"

func main() {
	fmt.Println(findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))

}

type UnionFindSets struct {
	Parent []int
}

func NewUnionFind(n int) *UnionFindSets {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFindSets{Parent: parent}
}

func (uf *UnionFindSets) Find(i int) int {
	if uf.Parent[i] != i {
		uf.Parent[i] = uf.Find(uf.Parent[i])
	}
	return uf.Parent[i]
}
func (uf *UnionFindSets) Union(i, j int) {
	uf.Parent[uf.Find(i)] = uf.Find(j)
}

func findCircleNum(isConnected [][]int) int {
	uf := NewUnionFind(len(isConnected))

	for i, v := range isConnected {
		for j, vv := range v {
			if vv == 1 {
				uf.Union(i, j) // i 和 j 联通加入并查集
			}
		}
	}

	var ans int
	for i, v := range uf.Parent {
		if i == v {
			// 	只有祖先是自己执行自己，其他人其实都是指向祖先
			ans++
		}
	}
	return ans
}

func findCircleNum3(isConnected [][]int) (ans int) {
	n := len(isConnected)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 {
				union(i, j)
			}
		}
	}
	for i, p := range parent {
		if i == p {
			ans++
		}
	}
	return
}

func findCircleNum2(isConnected [][]int) int {

	visited := make([]bool, len(isConnected))

	var dfs func(idx int)

	dfs = func(idx int) {
		if visited[idx] {
			return
		}
		visited[idx] = true

		for j, v := range isConnected[idx] {
			if v == 1 {
				dfs(j)
			}
		}
	}

	var ans int
	for i := range visited {
		if !visited[i] {
			// 当前节点未访问过
			ans++
			dfs(i)
		}
	}

	return ans
}
