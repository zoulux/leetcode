package main

import "fmt"

func main() {

	//fmt.Println(asteroidCollision([]int{
	//	2, 1, 2, -3,
	//}))
	//
	//fmt.Println(asteroidCollision([]int{
	//	-2, 1, 1, -1,
	//}))
	//
	//fmt.Println(asteroidCollision([]int{
	//	-2, 1, -1, 1,
	//}))

	fmt.Println(asteroidCollision([]int{
		5, 10, -5,
	}))

	fmt.Println(asteroidCollision([]int{
		10, 2, -5,
	}))

	fmt.Println(asteroidCollision([]int{
		8, -8,
	}))

	fmt.Println(asteroidCollision([]int{
		-8, 8,
	}))

}

func asteroidCollision(asteroids []int) []int {

	star := make([]int, 0)

	for _, v := range asteroids {
		flag := true

		for v < 0 && flag && len(star) > 0 && star[len(star)-1] > 0 {
			//2 < 2
			// 1 < 2
			// 2 > 1

			if star[len(star)-1] >= -v {
				flag = false // 都会爆炸，不用加新的星星
			}

			if star[len(star)-1] <= -v {
				star = star[:len(star)-1] // 左边删除
			}
		}

		if flag {
			star = append(star, v)
		}
	}
	return star
}

func asteroidCollision2(asteroids []int) []int {

	x := make([]int, 0) // 负数
	s := make([]int, 0) //正数
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			s = append(s, i)
		} else {
			if len(s) == 0 {
				x = append(x, i)
			} else {

				flag := false
				for len(s) > 0 {
					top := asteroids[s[len(s)-1]]
					v := -asteroids[i]
					if top > v {
						flag = true
						break
					} else if top < v {
						s = s[:len(s)-1] // 左边删除
					} else {
						s = s[:len(s)-1] // 左边删除
						flag = true
						break
					}
				}
				if !flag {
					x = append(x, i)
				}
			}
		}
	}

	var ans []int
	for _, v := range append(x, s...) {
		ans = append(ans, asteroids[v])
	}

	return ans
}
