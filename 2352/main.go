package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(equalPairs([][]int{
		{250, 78, 253},
		{334, 252, 253},
		{250, 253, 253},
	}))

}

func equalPairs(grid [][]int) int {
	mm := make(map[string]int)
	for i := 0; i < len(grid); i++ {
		var t []byte
		for j := 0; j < len(grid[i]); j++ {
			t = append(t, byte(grid[i][j]+'0'), ',')
		}
		mm[string(t)]++
	}

	var ans int
	for j := 0; j < len(grid[0]); j++ {
		var t []byte
		for i := 0; i < len(grid); i++ {
			t = append(t, byte(grid[i][j]+'0'), ',')
		}

		if v, ok := mm[string(t)]; ok {
			ans += v
		}
	}
	return ans
}

func equalPairs3(grid [][]int) int {
	mm := make(map[string]int)
	for i := 0; i < len(grid); i++ {
		var t []string
		for j := 0; j < len(grid[i]); j++ {
			t = append(t, string(grid[i][j]))
		}
		mm[strings.Join(t, ",")]++
	}

	var ans int
	for j := 0; j < len(grid[0]); j++ {
		var t []string
		for i := 0; i < len(grid); i++ {
			t = append(t, string(grid[i][j]))
		}

		if v, ok := mm[strings.Join(t, ",")]; ok {
			ans += v
		}
	}
	return ans
}

func equalPairs2(grid [][]int) int {
	colMap := make(map[int][]int) // 缓存列
	handle := func(i, j int) bool {
		row := grid[i] // 获取所在行
		var col []int
		if c, ok := colMap[j]; ok {
			col = c
		} else {
			for i := 0; i < len(grid); i++ {
				col = append(col, grid[i][j])
			}
			colMap[j] = col // 获取所在列
		}
		return reflect.DeepEqual(row, col) // 行列是否相等
	}
	var ans int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if handle(i, j) {
				ans++
			}
		}
	}
	return ans
}
