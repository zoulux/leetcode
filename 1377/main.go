package main

import (
	"fmt"
	"math"
)

func main() {

	intSize := 32 << (^uint(0) >> 63) // 32 or 64

	MaxInt := 1<<(intSize-1) - 1
	fmt.Println(intSize)
	fmt.Println(MaxInt)

	fmt.Println(math.Pow(2, 63))
	fmt.Println(math.MaxInt)
	fmt.Println(math.MinInt)

	fmt.Println(frogPosition(3, [][]int{
		{2, 1},
		{3, 2},
	}, 1,
		2))

	fmt.Println(frogPosition(7, [][]int{
		{1, 2},
		{1, 3},
		{1, 7},
		{2, 4},
		{2, 6},
		{3, 5},
	}, 1,
		7))

	fmt.Println(frogPosition(7, [][]int{
		{1, 2},
		{1, 3},
		{1, 7},
		{2, 4},
		{2, 6},
		{3, 5},
	}, 2,
		4))
}

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	mm := make(map[int][]int)
	for _, e := range edges {
		mm[e[0]] = append(mm[e[0]], e[1])
		mm[e[1]] = append(mm[e[1]], e[0])
	}

	var ans int
	visit := make(map[int]bool)
	var travel func(idx int, count int, total int)
	travel = func(idx int, count int, total int) {
		if count > t {
			// 次数到了

			return
		}
		if idx == target {
			ans = total
			return
		}

		visit[idx] = true
		for _, v := range mm[idx] {
			if !visit[v] {
				travel(v, count+1, len(mm[idx])*total)
			}
		}
	}

	travel(1, 0, 1)
	return 1 / float64(ans)
}
