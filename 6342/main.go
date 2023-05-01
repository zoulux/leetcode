package main

import (
	"fmt"
)

func main() {

	fmt.Println(firstCompleteIndex([]int{
		1, 4, 5, 2, 6, 3,
	}, [][]int{
		{4, 3, 5},
		{1, 2, 6},
	}))

	fmt.Println(firstCompleteIndex([]int{
		1, 3, 4, 2,
	}, [][]int{
		{1, 4},
		{2, 3},
	}))

	fmt.Println(firstCompleteIndex([]int{
		2, 8, 7, 4, 1, 3, 5, 6, 9,
	}, [][]int{
		{3, 2, 5},
		{1, 4, 6},
		{8, 7, 9},
	}))

}

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	maprow := make(map[int]int)
	mapcol := make(map[int]int)

	pos := make([][2]int, m*n+1)

	for i, r := range mat {
		for j, v := range r {
			pos[v] = [2]int{i, j} // 数的位置存储一下
		}
	}

	for ai, a := range arr {
		p := pos[a]
		i, j := p[0], p[1]
		maprow[i]++
		mapcol[j]++

		if maprow[i] == len(mat[i]) {
			return ai
		}

		if mapcol[j] == len(mat) {
			return ai
		}
	}
	return 0
}
