package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 1}))
	fmt.Println(permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1}))
	//fmt.Println(permute([]int{1}))
}

func permuteUnique(nums []int) [][]int {

	contains := func(idx []int, k int) bool {
		for _, v := range idx {
			if v == k {
				return true
			}
		}
		return false
	}

	ans := make([][]int, 0)
	var travel func(selectIdxs, selectNums []int)
	travel = func(selectIdxs, selectNums []int) {
		if len(selectNums) == len(nums) {
			ans = append(ans, append([]int{}, selectNums...))
			return
		}

		for i, v := range nums {

			// 当前的路径用过就跳过
			if contains(selectIdxs, i) {
				continue
			}

			// 因为排过序，这个值和前面的值相同
			// ！！！ 如果前面的那个路径用到，说明可以将当前的节点追加上去，加长枝干
			//  ！！！如果前面的那个路径没用到，说明当前层的逻辑前面已经计算过
			if i > 0 && v == nums[i-1] && !contains(selectIdxs, i-1) {
				continue
			}

			travel(append(selectIdxs, i), append(selectNums, v))
		}
	}

	sort.Ints(nums)
	travel([]int{}, []int{})
	return ans
}
