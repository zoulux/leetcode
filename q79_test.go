package leetcode

import "testing"

func exist(board [][]byte, word string) bool {

	visted := make([][]bool, len(board))
	mm := len(board)
	nn := len(board[0])

	for i := range board {
		visted[i] = make([]bool, len(board[i]))
	}

	var dfs func(track, i, j int) bool
	dfs = func(track, i, j int) bool {

		if track == len(word) {
			return true
		}

		if i < 0 || i >= mm || j < 0 || j >= nn {
			return false
		}

		if track == 0 {
			// 长度为0 的时候 在全表进行扫描
			for m := 0; m < mm; m++ {
				for n := 0; n < nn; n++ {
					if board[m][n] == word[track] && !visted[m][n] {
						visted[m][n] = true
						// 然后找到了 在这个边界的上下左右进行扫描
						res := dfs(track+1, m+1, n) || dfs(track+1, m, n+1) || dfs(track+1, m-1, n) || dfs(track+1, m, n-1)
						visted[m][n] = false
						if res {
							// 为真 则找到了 不继续去找
							return true
						}
					}
				}
			}
			return false // 第一轮都没找到 后面也就不用找了，直接 false
		}

		if board[i][j] != word[track] || visted[i][j] {
			// 两个字符不相等 或者 已经访问过
			return false
		}

		visted[i][j] = true // visted 访问过
		defer func() {
			visted[i][j] = false // 重置 visted
		}()

		// 继续搜索下个字母
		return dfs(track+1, i+1, j) || dfs(track+1, i, j+1) || dfs(track+1, i-1, j) || dfs(track+1, i, j-1)
	}
	return dfs(0, 0, 0)
}

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	t.Log(exist(board, "ABCCED"))
	t.Log(exist(board, "SEE"))
	t.Log(exist(board, "ABCB"))
	t.Log(exist([][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}, "abcd"))

}
