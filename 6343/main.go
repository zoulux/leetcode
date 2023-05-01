package main

import (
	"fmt"
	"math"
)

func main() {

	// 1000 => 10
	fmt.Println(minimumCost([]int{1, 1}, []int{10, 8}, [][]int{
		{6, 4, 9, 7, 1},
		{5, 2, 2, 1, 3},
		{3, 2, 5, 5, 2},
	}))

	// 999 => 12
	fmt.Println(minimumCost([]int{1, 1}, []int{10, 7}, [][]int{
		{3, 7, 5, 3, 3},
		{10, 2, 1, 2, 3},
		{7, 7, 5, 5, 4},
		{2, 6, 10, 2, 1},
	}))

	//961 => 11
	fmt.Println(minimumCost([]int{1, 1}, []int{3, 8}, [][]int{
		{2, 7, 3, 1, 5},
		{2, 8, 2, 8, 2},
		{3, 2, 2, 5, 1},
	}))

	//129 =>8
	fmt.Println(minimumCost([]int{1, 1}, []int{3, 8}, [][]int{
		{2, 7, 3, 1, 5},
		{2, 8, 2, 8, 2},
		{3, 2, 2, 5, 1},
	}))

	//33
	fmt.Println(minimumCost([]int{1, 1}, []int{4, 6}, [][]int{
		{1, 5, 3, 5, 5},
		{3, 4, 1, 5, 4},
		{3, 3, 2, 5, 5},
		{1, 5, 3, 4, 5},
	}))

	fmt.Println(minimumCost([]int{1, 1}, []int{7, 9}, [][]int{
		{1, 3, 1, 9, 1},
		{1, 9, 7, 8, 5},
		{1, 3, 4, 2, 5},
		{5, 5, 7, 5, 4},
	}))

	fmt.Println(minimumCost([]int{1, 1}, []int{4, 5}, [][]int{
		{1, 2, 3, 3, 2},
		{3, 4, 4, 5, 1},
	}))

	fmt.Println(minimumCost([]int{1, 1}, []int{4, 6}, [][]int{
		{3, 4, 2, 4, 1},
		{2, 5, 4, 2, 5},
		{3, 2, 1, 6, 3},
	}))
}

func minimumCost(start []int, target []int, specialRoads [][]int) int {

	// 将 specialRoads 中一些不合法的去掉
	for {
		flag := true
		for i, v := range specialRoads {
			p := Path{v[0], v[1], v[2], v[3]}
			if v[4] > p.cost() {
				flag = false
				specialRoads = append(specialRoads[:i], specialRoads[i+1:]...)
				break
			}
		}
		if flag {
			break
		}
	}

	mem := make(map[Path]int)
	var dfs func(path Path, specialIdx int) int
	dfs = func(p Path, specialIdx int) int {
		if v, ok := mem[p]; ok {
			return v
		}
		cost := p.cost() // 本来 (x1,y1) -> (x2,y2) 需要的花费
		for i := specialIdx; i < len(specialRoads); i++ {
			v := specialRoads[i]
			// (x1,y1) -> (x2,y2)  =>  (x1,y1) -> (v1,v2) -> (v3,v4) -> (x2,y2)
			// 如果有特殊点可以经过特殊点中转一下并且计算
			c := dfs(Path{p[0], p[1], v[0], v[1]}, i+1) +
				dfs(Path{v[2], v[3], p[2], p[3]}, i+1) +
				v[4]
			cost = min(cost, c)
		}
		mem[p] = cost
		return cost
	}

	return dfs([...]int{
		start[0],
		start[1],
		target[0],
		target[1],
	}, 0)
}

type Path [4]int

func (p *Path) cost() int {
	return abs(p[2]-p[0]) + abs(p[3]-p[1])
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(arr ...int) int {
	x := arr[0]
	for _, v := range arr[1:] {
		if v < x {
			x = v
		}
	}
	return x
}

func minimumCostErr(start []int, target []int, specialRoads [][]int) int {
	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}

	specialMap := make(map[[2]int][][3]int)
	for _, v := range specialRoads {
		curCost := abs(v[2]-v[0]) + abs(v[3]-v[1])
		if v[4] < curCost {
			specialMap[[2]int{v[0], v[1]}] = append(specialMap[[2]int{v[0], v[1]}], [...]int{v[2], v[3], v[4]})
		}
	}

	min := func(arr ...int) int {
		x := arr[0]
		for _, v := range arr[1:] {
			if v < x {
				x = v
			}
		}

		return x
	}
	cache := make([][]int, target[0]+1)
	for i := range cache {
		cache[i] = make([]int, target[1]+1)
	}
	visitSpecial := make(map[[3]int]bool)

	var dfs func(i, j int) int

	dfs = func(i, j int) int {

		if i > target[0] || j > target[1] || i < 0 || j < 0 {
			return math.MaxInt / 2
		}

		//if visit[[2]int{i, j}] {
		//	return math.MaxInt / 2
		//}

		if i == target[0] && j == target[1] {
			// 到达终点
			return 0
		}

		if v := cache[i][j]; v != 0 {
			return v
		}

		m := min(dfs(i+1, j), dfs(i, j+1)) + 1

		if vs, ok := specialMap[[2]int{i, j}]; ok {
			for vi, v := range vs {
				if visitSpecial[[3]int{i, j, vi}] {
					continue
				}
				visitSpecial[[...]int{i, j, vi}] = true // 原则上 specialMap 是可以重复使用，但是重复使用是没有意义的，并且会引入死循环

				// 存在特殊路径，特殊路径是要通往未访问过的点
				m = min(m, v[2]+dfs(v[0], v[1]))

				visitSpecial[[...]int{i, j, vi}] = false
			}
		}
		cache[i][j] = m
		return m
	}

	return dfs(start[0], start[1])
}
