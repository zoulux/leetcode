package main

import "fmt"

func main() {

	fmt.Println(canVisitAllRooms([][]int{
		{1},
		{2},
		{3},
		{},
	}))

	fmt.Println(canVisitAllRooms([][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	}))
}

func canVisitAllRooms(rooms [][]int) bool {
	visit := make([]bool, len(rooms))

	var dfs func(idx int)

	dfs = func(idx int) {
		if visit[idx] {
			return
		}

		visit[idx] = true
		room := rooms[idx]
		for _, v := range room {
			dfs(v)
		}
	}

	dfs(0)

	for _, v := range visit {
		if !v {
			return false
		}
	}
	return true
}
