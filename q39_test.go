package leetcode

import "testing"

func combinationSum(candidates []int, target int) (ans [][]int) {

	var solve func(track []int, start int, target int)

	solve = func(track []int, start int, target int) {

		if target == 0 {
			ans = append(ans, append([]int{}, track...))
			return
		}

		if target < 0 {
			return
		}

		for idx, elm := range candidates {
			if idx < start {
				continue
			}

			solve(append(track, elm), idx, target-elm)
		}
	}
	solve([]int{}, 0, target)
	return
}

func TestCombinationSum(t *testing.T) {
	// t.Log(combinationSum([]int{2, 3, 6, 7}, 7))
	t.Log(combinationSum([]int{2, 3, 5}, 8))
}
