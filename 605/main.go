package main

import "fmt"

func main() {

	fmt.Println(canPlaceFlowers([]int{
		0, 0, 0, 0, 1,
	}, 2))

	fmt.Println(canPlaceFlowers([]int{
		1, 0, 0, 0, 0, 0, 1,
	}, 2))

	fmt.Println(canPlaceFlowers([]int{
		0, 0, 1, 0, 1,
	}, 1))

	fmt.Println(canPlaceFlowers([]int{
		0,
	}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	ln := len(flowerbed)

	for i := 0; i < ln; i++ {
		v := flowerbed[i]
		if v == 0 {
			if i == 0 || i == ln-2 {
				if i+1 < ln && flowerbed[i+1] == 0 {
					// 左右两侧只要连续 2 个 0
					n--
				} else if ln == 1 {
					// 单独一个 【0】，特殊情况
					n--
				}
			} else {
				if i+2 < ln && flowerbed[i+1] == 0 && flowerbed[i+2] == 0 {
					// 中间连续 3 个 0
					n--
					i++
				}
			}
		}

		if n <= 0 {
			// 减到 0 就代表消耗完毕
			return true
		}
	}
	return n <= 0
}
