package leetcode

import (
	"testing"
)

func trap3(height []int) int {
	n := len(height)

	if n == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	l, r := height[0], height[n-1]
	lefth, righth := make([]int, n), make([]int, n)

	for i := range height {
		l = max(l, height[i])     // 左边最大值
		r = max(r, height[n-1-i]) // 右边最大值

		lefth[i] = l      // 当前位置左边柱子最大值
		righth[n-1-i] = r // 当前位置右边柱子最大值
	}

	ans := 0
	for i, v := range height {
		tmp := min(lefth[i], righth[i]) - v // 含水量是左右两边的最大值的最小值 - 当前格子的高度

		if tmp > 0 {
			ans += tmp // 如果大于 0 表示含水 就相加
		}
	}
	return ans
}

func TestTrap3(t *testing.T) {
	trap3(nil)
}
