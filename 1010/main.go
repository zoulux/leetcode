package main

import "fmt"

func main() {

	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
	fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))

}

func numPairsDivisibleBy60(time []int) int {
	mm := make(map[int]int)
	for _, v := range time {
		t := v % 60
		mm[t]++
	}

	// 对小于 n 进行求和 1+2+3+。。。+n-1
	count := func(n int) int {
		return (n) * (n - 1) / 2
	}

	ans := count(mm[0]) // 因为 0 不参与 60-k 的计算，所以先计算
	delete(mm, 0)       // 算完删掉

	for k, v := range mm {
		if t, ok := mm[60-k]; ok {
			if 60-k != k {
				ans += t * v
				delete(mm, 60-k) // 算完删掉
			} else {
				ans += count(t)
			}
		}
	}
	return ans
}

//20, 30, 30, 40, 40
//
//20, 20, 30, 30, 40, 40
