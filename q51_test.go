package leetcode

import (
	"fmt"
	"testing"
)

func solveNQueens(n int) (ans [][]string) {

	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	var isVaild func(i, j int) bool

	// r=>0
	// c=>1
	isVaild = func(r, c int) bool {

		// 当前不等于 。 就退出
		if board[r][c] != '.' {
			return false
		}

		// 行中没有 Q
		for i := 0; i < r; i++ {
			if board[i][c] == 'Q' {
				return false
			}

		}

		// 列中没有 Q
		for i := 0; i < c; i++ {
			if board[r][i] == 'Q' {
				return false
			}
		}

		// 右斜边没有 Q
		for i := 0; r-i >= 0 && c+i < n; i++ {
			if board[r-i][c+i] == 'Q' {
				return false
			}
		}

		// 左斜边没有 Q
		for i := 0; r-i >= 0 && c-i >= 0; i++ {
			if board[r-i][c-i] == 'Q' {
				return false
			}
		}

		return true
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			_board := make([]string, n)
			for j := 0; j < n; j++ {
				_board[j] = string(board[j])
			}
			ans = append(ans, _board)
			return
		}

		// 找 对应的 列
		for j := 0; j < n; j++ {

			if !isVaild(i, j) {
				continue
			}

			board[i][j] = 'Q'
			dfs(i + 1) // 找下一列
			board[i][j] = '.'
		}
	}
	// 才第 0 行开始迭代
	dfs(0)
	return ans
}

func TestSolveNQueens(t *testing.T) {
	t.Log(solveNQueens(4))
	t.Log(solveNQueens(3))
	t.Log(solveNQueens(2))
	t.Log(solveNQueens(1))
	res := solveNQueens(8)

	fmt.Println(len(res))
	for _, v1 := range res {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
		fmt.Println("---")
	}

}
