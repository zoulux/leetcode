package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))

	fmt.Println(spiralOrder([][]int{
		{3},
		{2},
	}))

}

func spiralOrder(matrix [][]int) []int {

	ans := make([]int, 0)

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}

	rl, rh := 0, len(matrix)-1
	cl, ch := 0, len(matrix[0])-1
	for {

		// 向右走
		for c := cl; c <= ch; c++ {
			ans = append(ans, matrix[rl][c])
		}
		rl++
		if rl > rh {
			break
		}

		// 向下走
		for r := rl; r <= rh; r++ {
			ans = append(ans, matrix[r][ch])
		}

		ch--
		if ch < 0 {
			break
		}

		// 向左走
		for c := ch; c >= cl; c-- {
			ans = append(ans, matrix[rh][c])
		}

		rh--
		if rh < 0 {
			break
		}

		// 向上走
		for r := rh; r >= rl; r-- {
			ans = append(ans, matrix[r][cl])
		}

		cl++
		if cl > ch {
			break
		}
	}

	return ans
}
