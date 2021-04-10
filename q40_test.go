package leetcode

import (
	"sort"
	"testing"
)

func combinationSum2(candidates []int, target int) (ans [][]int) {

	var solve func(track []int, start, target int)

	min := 0

	solve = func(track []int, start, target int) {

		if target < min {
			if target == 0 {
				tmp := make([]int, len(track))
				copy(tmp, track)
				ans = append(ans, tmp)
				return
			}
			return
		}

		for i := start; i < len(candidates); i++ {

			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			elm := candidates[i]
			solve(append(track, elm), i+1, target-elm)
		}
	}

	sort.Ints(candidates)
	min = candidates[0]

	solve([]int{}, 0, target)

	return
}

func TestCombinationSum2(t *testing.T) {
	// t.Log(combinationSum2([]int{2, 3, 6, 7}, 7))
	t.Log(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
