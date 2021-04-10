package leetcode

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func merge2(intervals [][]int) [][]int {

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		interval:=intervals[i]

		top := ans[len(ans)-1]

		if interval[0]>top[1]{
			ans=append(ans,interval)
		}else{
			mi := min(interval[0], top[0])
			ma := max(interval[1], top[1])
			ans[len(ans)-1] = []int{mi, ma}
		}
	}
	return ans
}

func TestMerge2(t *testing.T) {
	var a uint = 1
	var b uint = 2

	fmt.Println(a-b, math.MaxInt32)

	// [[2,3],[2,2],[3,3],[1,3],[5,7],[2,2],[4,6]]
	t.Log(merge2([][]int{
		// {2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10},
		// {2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6},
		{4, 5}, {2, 4}, {4, 6}, {3, 4}, {0, 0}, {1, 1}, {3, 5}, {2, 2},
	}))
}
