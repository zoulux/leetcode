package leetcode

import (
	"strings"
	"testing"
)

func generateParenthesis(n int) (ans []string) {

	var generate func(track []string, lCount, rCount int)
	generate = func(track []string, lCount, rCount int) {
		if lCount == n && rCount == n {
			ans = append(ans, strings.Join(track, ""))
			return
		}

		if lCount < n {
			generate(append(track, "("), lCount+1, rCount)
		}

		if rCount < lCount {
			generate(append(track, ")"), lCount, rCount+1)
		}
	}
	generate([]string{}, 0, 0)
	return
}

func TestGenerateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(1))
}

func generateParenthesis2(n int) (ans []string) {

	// 这里去掉一个路径参数
	// 这里额外加了一个参数，lCount, rCount ，否则每次要去统计 track 中的 左括号数量和右括号数量，直接取值更为方便
	var generate func(track []string, lCount, rCount int)

	generate = func(track []string, lCount, rCount int) {
		if lCount == n && rCount == n { // 满足条件加入结果
			ans = append(ans, strings.Join(track, "")) // 满足条件再将数组转成字符串，减少开销
			return
		}

		if lCount < n { // 左括号最多 n 个，不能再多了
			// 做选择
			track = append(track, "(")
			lCount++ // 左括号数量+1
			generate(track, lCount, rCount)
			// 撤销
			track = track[:len(track)-1]
			lCount-- // 左括号数量-1
		}

		if rCount < lCount { // 由括号最多和左括号数量一样，不能再多了
			// 做选择
			track = append(track, ")")
			rCount++ // 右括号数量+1
			generate(track, lCount, rCount)
			// 撤销
			track = track[:len(track)-1]
			rCount-- // 右括号数量-1
		}
	}

	generate([]string{}, 0, 0)
	return
}
