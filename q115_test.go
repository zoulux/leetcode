package leetcode

import "testing"

func numDistinct(s string, t string) int {

	mem := make([][]int, len(s))
	for i := range mem {
		mem[i] = make([]int, len(t))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var dp func(i, j int) int

	dp = func(i, j int) int {
		if j < 0 {
			return 1
		}
		if i < 0 {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}

		if s[i] == t[j] {
			mem[i][j] = dp(i-1, j) + dp(i-1, j-1)
		} else {
			mem[i][j] = dp(i-1, j)
		}

		return mem[i][j]
	}

	return dp(len(s)-1, len(t)-1)
}

func TestNumDistinct(t *testing.T) {
	t.Log(numDistinct("rabbbit", "rabbit"))
	t.Log(numDistinct("babgbag", "bag"))

}

//详细注释见下方 记忆化递归 代码
func numDistinct2(s string, t string) int {
	sLen, tLen := len(s), len(t)

	var helper func(i, j int) int
	helper = func(i, j int) int {
		if j < 0 { // base case
			return 1
		}
		if i < 0 { // 这两个base case 的顺序不能调换！因为 i<0 且 j<0 时 应该返回1
			return 0
		}
		if s[i] == t[j] {
			return helper(i-1, j) + helper(i-1, j-1)
		} else {
			return helper(i-1, j)
		}
	}
	return helper(sLen-1, tLen-1)
}
