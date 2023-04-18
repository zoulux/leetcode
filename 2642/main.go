package main

import (
	"fmt"
	"math"
)

func main() {

	g := Constructor(4, [][]int{
		{0, 2, 5},
		{0, 1, 2},
		{1, 2, 1},
		{3, 0, 3},
	})

	fmt.Println(g.ShortestPath(3, 7))

}

const inf = math.MaxInt / 2 // 防止更新最短路时加法溢出

type Graph struct {
	g [][]int
	n int
}

func Constructor(n int, edges [][]int) Graph {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)

		for j := range g[i] {
			g[i][j] = inf
		}
	}

	for _, e := range edges {
		g[e[0]][e[1]] = e[2] // x->y = cost
	}
	return Graph{g: g, n: n}
}

func (this *Graph) AddEdge(edge []int) {
	this.g[edge[0]][edge[1]] = edge[2]
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	visited := make([]bool, this.n)
	dis := make([]int, this.n)
	for i := range dis {
		dis[i] = inf
	}

	dis[node1] = 0 // 起点最短路为 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for {
		x := -1
		for i := 0; i < this.n; i++ {
			if !visited[i] && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}

		if x == -1 || dis[x] == inf {
			return -1
		}

		if x == node2 {
			return dis[x]
		}

		visited[x] = true

		for y, w := range this.g[x] {
			dis[y] = min(dis[y], dis[x]+w)
		}
	}
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
