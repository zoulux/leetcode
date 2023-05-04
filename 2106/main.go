package main

import "fmt"

func main() {

	//[[200000,10000]]

	// test 10
	fmt.Println(maxTotalFruits([][]int{
		{200000, 10000},
	}, 0, 200000))

	fmt.Println(maxTotalFruits([][]int{
		{200000, 10000},
	}, 200000, 200000))

	//输入：fruits = [[2,8],[6,3],[8,6]], startPos = 5, k = 4

	fmt.Println(9 == maxTotalFruits([][]int{
		{2, 8},
		{6, 3},
		{8, 6},
	}, 5, 4))

	//输入：fruits = [[0,9],[4,1],[5,7],[6,2],[7,4],[10,9]], startPos = 5, k = 4

	fmt.Println(maxTotalFruits([][]int{
		{0, 9},
		{4, 1},
		{5, 7},
		{6, 2},
		{7, 4},
		{10, 9},
	}, 5, 4))

}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	abs := func(a int) int {
		if a < 0 {
			a = -a
		}
		return a
	}

	var i, ans, win int

	for j, f := range fruits {
		win += f[1]

		for i <= j && f[0]-fruits[i][0]+min(abs(startPos-fruits[i][0]), abs(startPos-f[0])) > k {
			win -= fruits[i][1]
			i++
		}
		ans = max(ans, win)
	}

	return ans
}
func maxTotalFruits2(fruits [][]int, startPos int, k int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	fend := fruits[0][0]
	fstart := fruits[0][0]
	fmap := make(map[int]int)
	for _, f := range fruits {
		fmap[f[0]] = f[1]
		if f[0] > fend {
			fend = f[0]
		}

		if f[0] < fstart {
			fstart = f[0]
		}
	}

	//cache := make(map[[2]int]int)

	var dfs func(idx int, k int) int
	dfs = func(idx int, k int) int {
		if k < 0 {
			return 0
		}

		if idx < fstart || idx > fend {
			return 0
		}

		//if v, ok := cache[[2]int{idx, k}]; ok {
		//	return v
		//}
		//
		c, ok := fmap[idx]
		if ok && c != 0 {
			fmap[idx] = 0
		}
		ans := c + max(dfs(idx-1, k-1), dfs(idx+1, k-1))

		if ok {
			fmap[idx] = c
		}

		//cache[[2]int{idx, k}] = ans
		return ans
	}
	return dfs(startPos, k)
}
