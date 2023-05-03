package main

import (
	"fmt"
	"math"
)

func main() {

	yb := int(math.Ceil(math.Log(float64(10e6)) / math.Log(float64(2))))
	fmt.Println(yb)
	//
	//fmt.Println(math.Log(1))

	//fmt.Println(powerfulIntegers(2, 1, 10))
	fmt.Println(powerfulIntegers(2, 2, 10e6))
	//fmt.Println(powerfulIntegers(2, 2, 400000))
	//fmt.Println(powerfulIntegers(2, 3, 10))
	//fmt.Println(powerfulIntegers(3, 5000000, 15))

}

func powerfulIntegers(x int, y int, bound int) []int {
	pow := func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}

	xb, yb := 1, 1 // 1 需要特殊处理一下
	if x != 1 {
		xb = int(math.Ceil(math.Log(float64(bound)) / math.Log(float64(x))))
	}

	if y != 1 {
		yb = int(math.Ceil(math.Log(float64(bound)) / math.Log(float64(y))))
	}

	mm := make(map[int]struct{})

	visit := make(map[[2]int]bool)
	var dfs func(l, r int)
	dfs = func(l, r int) {
		if l > xb || r > yb {
			// 递归出口
			return
		}

		// 访问过就无需再计算了
		if visit[[2]int{l, r}] {
			return
		}

		// 表示当前访问过
		visit[[2]int{l, r}] = true

		// 计算 t
		t := pow(x, l) + pow(y, r)
		if t <= bound {
			mm[t] = struct{}{} // 存起来 hash 表去重
			dfs(l+1, r)        // 算算 x
			dfs(l, r+1)        // 算算 y
		}

	}
	dfs(0, 0) // 从 0 开始算

	// 将哈希表转为列表
	ans := make([]int, 0, len(mm))
	for k := range mm {
		ans = append(ans, k)
	}
	return ans
}
