package main

import "fmt"

func main() {

	//fmt.Println(circularGameLosers(5, 2))
	//fmt.Println(circularGameLosers(4, 4))
	fmt.Println(circularGameLosers(5, 3))

}

func circularGameLosers(n int, k int) []int {

	j := 1
	i := 1
	mm := make(map[int]bool)
	for {
		if mm[i] {
			break
		}
		mm[i] = true
		i += k * j
		j++
		if i > n {
			i %= n
			if i == 0 {
				i = n
			}
		}
	}

	var ans []int
	for i := 1; i <= n; i++ {
		if !mm[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
