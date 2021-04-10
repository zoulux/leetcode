package leetcode

import (
	"sort"
	"strings"
	"testing"
)

func permuteUnique2(nums []int) (ans [][]int) {

	cache := map[string]struct{}{}

	var dfs func(nums []int, idx int)
	dfs = func(nums []int, idx int) {
		if idx == len(nums) {

			key := hash(nums)
			if _, ok := cache[key]; ok {
				// return
			}

			cache[key] = struct{}{}

			ans = append(ans, append([]int{}, nums...))
			return
		}

		for i := idx; i < len(nums); i++ {
			if idx != i && nums[idx] == nums[i] {
				continue
			}
			// if idx+1 < len(nums) && nums[idx] == nums[idx+1] {
			// 	continue
			// }
			nums[idx], nums[i] = nums[i], nums[idx]
			dfs(nums, idx+1)
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	sort.Ints(nums)
	dfs(nums, 0)
	return
}

func permuteUnique(nums []int) (ans [][]int) {
	visited := make([]bool, len(nums))

	var dfs func(track []int)
	dfs = func(track []int) {
		if len(track) == len(nums) {
			ans = append(ans, append([]int{}, track...))
			return
		}

		for idx := 0; idx < len(nums); idx++ {
			num := nums[idx]
			if visited[idx] {
				continue
			}

			if idx > 0 && nums[idx] == nums[idx-1] && !visited[idx-1] {
				// 核心的一句话
				continue
			}

			visited[idx] = true
			dfs(append(track, num))
			visited[idx] = false
		}
	}

	sort.Ints(nums)
	dfs([]int{})
	return
}

func hash(nums []int) string {
	var sb strings.Builder
	// sum := 0
	for _, num := range nums {
		// sum = sum >> num
		sb.WriteRune(rune(num))
	}

	return sb.String()
}

func TestPermuteUnique(t *testing.T) {
	// t.Log(permuteUnique([]int{1, 1, 3}))
	t.Log(permuteUnique([]int{2, 2, 1, 1}))
}
