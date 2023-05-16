package main

import (
	"fmt"
	"math"
)

func main() {

	// test 5
	fmt.Println(15 == minDifficulty([]int{
		7, 1, 7, 1, 7, 1,
	}, 3))

	fmt.Println(843 == minDifficulty([]int{
		11, 111, 22, 222, 33, 333, 44, 444,
	}, 6))
}

func minDifficulty3(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	dp := make([][]int, len(jobDifficulty))
	for i := range dp {
		dp[i] = make([]int, d)
	}

	//for i, v := range jobDifficulty {
	//
	//	for k := i; k < i+d; k++ {
	//
	//	}
	//}

	var dfs func(idx, day int) int
	dfs = func(idx, day int) int {
		if day == 0 {
			mx := 0
			for _, v := range jobDifficulty[:idx+1] {
				if v > mx {
					mx = v
				}
			}
			return mx
		}

		min := math.MaxInt / 2
		mx := 0
		for k := idx; k >= day; k-- { // 需要保留 day 天，因为前面还需要保留
			if jobDifficulty[k] > mx {
				mx = jobDifficulty[k]
			}
			if t := mx + dfs(k-1, day-1); t < min {
				// 从 k-1 的索引向前找，天数减 1
				min = t
			}
		}
		return min
	}
	return dfs(len(jobDifficulty)-1, d-1)
}

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	mm := make(map[[2]int]int)

	var dfs func(idx, day int) int
	dfs = func(idx, day int) (ret int) {
		if v, ok := mm[[2]int{idx, day}]; ok {
			return v
		}
		defer func() {
			mm[[2]int{idx, day}] = ret
		}()

		if day == 0 {
			mx := 0
			for _, v := range jobDifficulty[:idx+1] {
				if v > mx {
					mx = v
				}
			}
			//mm[[2]int{idx, day}] = mx
			return mx
		}

		min := math.MaxInt / 2
		mx := 0
		for k := idx; k >= day; k-- { // 需要保留 day 天，因为前面还需要保留
			if jobDifficulty[k] > mx {
				mx = jobDifficulty[k]
			}
			if t := mx + dfs(k-1, day-1); t < min {
				// 从 k-1 的索引向前找，天数减 1
				min = t
			}
		}
		//mm[[2]int{idx, day}] = min
		return min
	}
	return dfs(len(jobDifficulty)-1, d-1)
}
