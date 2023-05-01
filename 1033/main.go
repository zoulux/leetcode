package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(numMovesStones(3, 5, 1))

}

func numMovesStones(a int, b int, c int) []int {
	s := []int{a, b, c}
	sort.Ints(s)
	a, b, c = s[0], s[1], s[2]
	min := 0
	max := c - a - 2

	if c-a == 2 {
		min = 0 // 有序无需移动
	} else if b-a == 1 || c-b == 1 {
		min = 1
		// 一步移动两边
	} else if b-a == 2 || c-b == 2 {
		min = 1
		// 一步移动到中间
	} else {
		// 那就左右两边向中间各移动1步
		min = 2
	}

	return []int{min, max}
}
