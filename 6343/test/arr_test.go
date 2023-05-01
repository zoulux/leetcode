package main

import (
	"fmt"
	"testing"
)

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

func TestArr(t *testing.T) {

	arr := [][]int{
		{1, 5, 3, 5, 5},
		{3, 4, 1, 5, 4},
		{3, 3, 2, 5, 5},
		{1, 5, 3, 4, 5},
	}

	//for _, v := range specialRoads {
	//	curCost := abs(v[2]-v[0]) + abs(v[3]-v[1])
	//	if v[4] < curCost {
	//		specialMap[[2]int{v[0], v[1]}] = append(specialMap[[2]int{v[0], v[1]}], [...]int{v[2], v[3], v[4]})
	//	}
	//}

	for {
		flag := true
		for i, v := range arr {
			p := Path{v[0], v[1], v[2], v[3]}
			if v[4] > p.cost() {
				flag = false
				arr = append(arr[:i], arr[i+1:]...)
				break
			}
		}
		if flag {
			break
		}
	}

	//for i, v := range arr {
	//	p := Path{v[0], v[1], v[2], v[3]}
	//	if p.cost() < v[0] {
	//		arr = append(arr[:i], arr[i+1:]...)
	//	}
	//}
	fmt.Println(arr)
}
