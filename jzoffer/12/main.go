package main

import "fmt"

func main() {

	fmt.Println(exist([][]byte{
		{'X', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'F'},
	}, "BCCED"))

	fmt.Println(exist([][]byte{
		{'X'},
	}, "Q"))
}

func exist(board [][]byte, word string) bool {

	visit := make([][]bool, len(board))
	for i := range visit {
		visit[i] = make([]bool, len(board[0]))
	}

	var travel func(i, j, w int) bool
	travel = func(i, j, w int) bool {
		if w >= len(word) {
			// 说明已经找到完整的字符了
			return true
		}

		// 不相等直接跳过不用管
		if board[i][j] != word[w] {
			return false
		}

		visit[i][j] = true

		// 上方是否还有空，是否没有访问过
		if i-1 >= 0 && !visit[i-1][j] {
			if travel(i-1, j, w+1) {
				return true
			}
		}

		// 左边是否还有空，是否没有访问过
		if j-1 >= 0 && !visit[i][j-1] {
			if travel(i, j-1, w+1) {
				return true
			}
		}

		// 下边是否还有空，是否没有访问过
		if i+1 < len(board) && !visit[i+1][j] {
			if travel(i+1, j, w+1) {
				return true
			}
		}

		// 右边是否还有空，是否没有访问过
		if j+1 < len(board[0]) && !visit[i][j+1] {
			if travel(i, j+1, w+1) {
				return true
			}
		}

		visit[i][j] = false
		return len(word) == w+1
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				// 先找到第一个字符相等的
				visit[i][j] = true
				if travel(i, j, 0) {
					return true
				}
				visit[i][j] = false
			}
		}
	}

	return false
}
