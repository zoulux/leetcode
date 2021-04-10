package leetcode

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

const targetNum = 11

func copyArr(track []int) [targetNum]int {
	arr := [targetNum]int{}

	nn := len(track)

	for i := 0; i < targetNum; i++ {
		if i < nn {
			arr[i] = track[i]
		} else {
			arr[i] = targetNum
		}
	}

	return arr
}

func subsetsWithDup1(nums []int) (ret [][]int) {
	sort.Ints(nums)

	n := len(nums)

	ans := map[[11]int]struct{}{}

	var dfs func(track []int, start int)

	dfs = func(track []int, start int) {
		ans[copyArr(track)] = struct{}{}

		for i := start; i < n; i++ {
			dfs(append(track, nums[i]), i+1)
		}
	}

	dfs([]int{}, 0)

	for k := range ans {

		retRow := []int{}
		for _, n := range k {
			if n != targetNum {
				retRow = append(retRow, n)
			}

		}
		ret = append(ret, retRow)

	}

	return
}

func subsetsWithDup(nums []int) (ret [][]int) {
	sort.Ints(nums)

	n := len(nums)

	ans := map[string]struct{}{}

	var dfs func(track []int, start int)

	dfs = func(track []int, start int) {

		var sb strings.Builder
		for _, v := range track {
			sb.WriteString(strconv.Itoa(v))
		}

		ans[sb.String()] = struct{}{}

		for i := start; i < n; i++ {

			dfs(append(track, nums[i]), i+1)
		}
	}

	dfs([]int{}, 0)

	for k := range ans {
		splitNums := strings.Split(k, "")

		retRow := []int{}
		for _, n := range splitNums {
			num, e := strconv.Atoi(n)
			if e == nil {
				retRow = append(retRow, num)
			}
		}
		ret = append(ret, retRow)
	}
	return
}

func TestSubsetsWithDup(t *testing.T) {
	// nums := []int{-1, 1, 2}
	nums := []int{2, 1, 2, 1, 3}
	// nums := []int{1, 2, 2}

	t.Log(subsetsWithDup2(nums))
}

func subsetsWithDup2(nums []int) (ret [][]int) {
	sort.Ints(nums)

	n := len(nums)

	ans := [][]int{}
	var dfs func(track []int, start int)

	dfs = func(track []int, start int) {
		t := make([]int, len(track))
		copy(t, track)
		ans = append(ans, t)
		// ans = append(ans, append([]int(nil), track...))

		for i := start; i < n; i++ {
			if i > start && nums[i] == nums[i-1] {
				// 当前值和上次用的值是一样的 表示用过了
				continue
			}

			dfs(append(track, nums[i]), i+1)
		}
	}

	dfs([]int{}, 0)

	return ans
}
