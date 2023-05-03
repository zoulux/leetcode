package main

import (
	"fmt"
)

func main() {
	fmt.Println(minFlips(8, 3, 5))
	//fmt.Println(minFlips(2, 6, 5))
	//fmt.Println(minFlips(4, 2, 7))
	//fmt.Println(minFlips(1, 2, 3))
	//fmt.Println(minFlips(10e9, 2, 3*10e9))

}

//1000
//0011
//0101

func minFlips(a int, b int, c int) int {
	var ans int
	for c != 0 || a != 0 || b != 0 {
		ai, bi, ci := a&1, b&1, c&1

		if ai|bi != ci {
			// 需要变化
			//if c
			if ci == 1 {
				// ai=0
				// bi=0
				ans++
			} else {
				//	ai =1 bi=1
				//	ai =1 bi=0
				//	ai =0 bi=1
				if ai == 1 {
					ans++
				}
				if bi == 1 {
					ans++
				}
			}
		}
		a, b, c = a>>1, b>>1, c>>1
	}
	return ans
}
