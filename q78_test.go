package leetcode

import (
	"sort"
	"testing"
)

func subsets(nums []int) (ans [][]int) {
	var generate func(start int, track []int)
	generate = func(start int, track []int) {
		ans = append(ans, track)

		for i := start; i < len(nums); i++ {
			generate(i+1, append(track, nums[i]))
		}
	}

	generate(0, []int{})
	return
}

func subsets2(nums []int) (ans [][]int) {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	n := nums[len(nums)-1]
	res := subsets2(nums[0 : len(nums)-1])

	for i := 0; i < len(res); i++ {
		res = append(res, res[i])
		res[len(res)-1] = append(res[len(res)-1], n)
	}

	return res
}

func subsets3(nums []int) (ans [][]int) {

	ans = append(ans, []int{})
	

	for i := 0; i < len(nums); i++ {
		lenAns := len(ans)
		for j := 0; j < lenAns; j++ {
			tmp := ans[j]
			tmp = append(tmp, nums[i])
			sort.Ints(tmp)
			ans = append(ans, tmp)
		}
	}
	return
}

func TestSubsets(t *testing.T) {
	// t.Log(subsets2([]int{1, 2, 3}))
	// t.Log(subsets([]int{1, 2, 3}))
	t.Log(subsets3([]int{1, 2, 3}))
}
