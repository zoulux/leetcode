package leetcode

import (
	"testing"
)

func partition3(s string) (ans [][]string) {
	n := len(s)

	splits := []string{}

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string{}, splits...))
			return
		}

		for j := i; j < n; j++ {
			if isPartition(s[i : j+1]) {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[0 : len(splits)-1]
			}
		}
	}

	dfs(0)
	return
}

func partition2(s string) (ans [][]string) {
	n := len(s)

	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	var dfs func(i int, splits []string)
	dfs = func(i int, splits []string) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}

		for j := i; j < n; j++ {
			if f[i][j] {
				dfs(j+1, append(splits, s[i:j+1]))
			}
		}
	}

	dfs(0, []string{})
	return
}

func TestPartition2(t *testing.T) {
	t.Log(partition3("aab"))
	// t.Log(partition2("aab"))
	// t.Log(isPartition("aab"))
	// t.Log(isPartition("a"))
	// t.Log(partition2("cbbbcc"))
	// t.Log(partition2("cbbbca"))

}

func partition(s string) [][]string {

	ans := [][]string{}
	n := len(s)
	var solve func(track []string, start int)

	solve = func(track []string, start int) {

		if start >= n {
			m := 0
			for _, ss := range track {
				m += len(ss)
			}
			if m == n {
				ans = append(ans, append([]string{}, track...))
			}
			return
		}

		for i := start; i < len(s); i++ {
			for j := i + 1; j <= len(s); j++ {
				sub := s[i:j] // 左闭右开
				if isPartition(sub) {
					solve(append(track, sub), j)
				}
			}
		}
	}

	solve([]string{}, 0)
	return ans
}

func isPartition(s string) bool {

	if len(s) == 0 {
		return false
	}

	left, right := 0, len(s)-1

	for left <= right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func TestPartition(t *testing.T) {
	// t.Log(partition("aab"))
	t.Log(partition("cbbbcc"))

}

func TestIsPartition(t *testing.T) {
	t.Log(isPartition("aab"))
	t.Log(isPartition("aabaa"))
	t.Log(isPartition("aa"))
	t.Log(isPartition("a"))
	t.Log(isPartition("aaba"))
}
