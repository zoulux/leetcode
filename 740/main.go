package main

import (
	"fmt"
	"sort"
)

func main() {

	// test 30
	fmt.Println(18 == deleteAndEarn([]int{
		1, 1, 1, 2, 4, 5, 5, 5, 6,
	}))

	fmt.Println(deleteAndEarn([]int{
		2, 2, 3, 3, 3, 4,
	}))

	fmt.Println(deleteAndEarn([]int{
		3,
	}))

}

func deleteAndEarn(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	mm := make(map[int]int) // 存储字母对应的数字以及个数
	for _, n := range nums {
		mm[n]++
	}

	newnums := make([]int, 0, len(mm)) // 存储字母的key
	for k := range mm {
		newnums = append(newnums, k)
	}
	sort.Ints(newnums) // 排个序，可以直接和 i-1 作比较

	dp := make([]int, len(mm))

	dp[0] = newnums[0] * mm[newnums[0]] // 第 0 个默认值
	for i := 1; i < len(newnums); i++ {
		if newnums[i]-newnums[i-1] != 1 {
			// 不等于就等于不会触发警报系统，可以直接掠夺 dp[i-1] 的资源
			dp[i] = dp[i-1] + newnums[i]*mm[newnums[i]]
		} else {
			// 等于就等于会触发警报系统，要么偷取当前的+dp[i-2], 要么放弃当前的，偷取dp[i-1]
			if i-2 >= 0 {
				dp[i] = max(dp[i-2]+newnums[i]*mm[newnums[i]], dp[i-1])
			} else {
				dp[i] = max(newnums[i]*mm[newnums[i]], dp[i-1])
			}
		}
	}

	return dp[len(dp)-1]
}
