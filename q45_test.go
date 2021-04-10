package leetcode

import "testing"

func jump(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	faster := 0
	step := 0
	end := 0
	for i, elm := range nums {

		if i+elm > faster {
			faster = i + elm
		}

		if faster >= n-1 {
			return step + 1
		}

		if end == i {
			step++
			end = faster
		}

	}
	return step
}

func TestJump(t *testing.T) {
	t.Log(jump([]int{2, 3, 1, 1, 4}))
	t.Log(jump([]int{0}))
}
