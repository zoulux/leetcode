package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(maximumScore([]int{
		1, 2, 3,
	}, []int{
		3, 2, 1,
	}))

	fmt.Println(maximumScore([]int{
		-5, -3, -3, -2, 7, 1,
	}, []int{
		-10, -5, 3, 4, 6,
	}))
}

func maximumScore(nums []int, multipliers []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	cache := make(map[[3]int]int) // 缓存一下之前是否计算过

	var dfs func(mIdx int, left, right int) int
	dfs = func(mIdx int, left, right int) int {

		// 如果 mIdx 到了最后，说明结束了
		if mIdx == len(multipliers) {
			return 0
		}

		if left > right {
			return 0
		}

		if v, ok := cache[[3]int{mIdx, left, right}]; ok {
			// 计算过，直接返回
			return v
		}

		res := max(
			multipliers[mIdx]*nums[left]+dfs(mIdx+1, left+1, right),  // 如果取 nums 左边的值，那接下来就从 nums 的  left+1 开始取
			multipliers[mIdx]*nums[right]+dfs(mIdx+1, left, right-1), // 同理
		)

		cache[[3]int{mIdx, left, right}] = res
		return res
	}

	return dfs(0, 0, len(nums)-1) // 从 multipliers 的第 0 位开始，nums 的头尾分别是 0 和  len(nums)-1
}

func maximumScore2(nums []int, multipliers []int) int {

	var ans int
	for len(multipliers) > 0 {
		first, last := 0, len(nums)-1

		firstm := 0
		ansFirst := math.MinInt / 2
		for i := 0; i < len(multipliers); i++ {
			if t := nums[first] * multipliers[i]; t > ansFirst {
				firstm = i
				ansFirst = t
			}
		}

		lastm := 0
		ansLast := math.MinInt / 2
		for i := 0; i < len(multipliers); i++ {
			if t := nums[last] * multipliers[i]; t > ansLast {
				lastm = i
				ansLast = t
			}
		}

		if ansFirst > ansLast {
			nums = nums[1:] //去头
			multipliers = append(multipliers[:firstm], multipliers[firstm+1:]...)

			ans += ansFirst

		} else {
			nums = nums[:len(nums)-1] //去尾巴
			multipliers = append(multipliers[:lastm], multipliers[lastm+1:]...)
			ans += ansLast
		}
	}

	return ans
}
