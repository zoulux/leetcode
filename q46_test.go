package leetcode

import "testing"

func permute2(nums []int) (ans [][]int) {

	var dfs func(nums []int, idx int)
	dfs = func(nums []int, idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, nums...))
			return
		}

		for i := idx; i < len(nums); i++ {
			nums[idx], nums[i] = nums[i], nums[idx]
			dfs(nums, idx+1)
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	dfs(nums, 0)
	return
}

func permute(nums []int) (ans [][]int) {

	trackcache := map[int]struct{}{}
	track := []int{}
	var dfs func()
	dfs = func() {
		if len(track) == len(nums) {
			ans = append(ans, append([]int{}, track...))
			return
		}

		for idx := 0; idx < len(nums); idx++ {
			num := nums[idx]
			if _, ok := trackcache[num]; ok {
				continue
			}
			trackcache[num] = struct{}{}
			track = append(track, num)
			dfs()
			delete(trackcache, num)
			track = track[:len(track)-1]
		}
	}
	dfs()
	return
}

func TestPermute(t *testing.T) {
	t.Log(permute([]int{1, 2, 3}))
	t.Log(permute2([]int{1, 2, 3}))
}
