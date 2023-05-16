package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(getNandResult(3, []int{1, 2}, [][]int{
		{1, 2, 3},
		{0, 0, 3},
		{1, 2, 2},
	}))
	fmt.Println(getNandResult(4, []int{4, 6, 4, 7, 10, 9, 11}, [][]int{
		{1, 5, 7},
		{1, 7, 14},
		{0, 6, 7},
		{1, 6, 5},
	}))

}

func getNandResult(k int, arr []int, operations [][]int) int {
	n := len(arr)

	var xor []int
	for _, v := range operations {
		t, x, y := v[0], v[1], v[2]
		if t == 0 {
			arr[x] = y
		} else {
			//若 type = 1，表示运算操作，将数字 y 进行 x*n 次「与非」操作，第 i 次与非操作为 y = y NAND arr[i%n]；
			//运算操作结果即：y NAND arr[0%n] NAND arr[1%n] NAND arr[2%n] ... NAND arr[(x*n-1)%n]

			//k 位 2进制
			//bits.Len()

			for i := 1; i <= int(math.Pow(float64(x), float64(n))); i++ {
				y = ^(n & arr[i%n])
			}
			xor = append(xor, y)
		}
	}
	fmt.Println(xor)

	var ans int
	for _, v := range xor {
		ans = ^v
	}
	return ans
}
