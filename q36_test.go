package leetcode

import (
	"testing"
)

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		mm := [3][9]bool{}

		for j := 0; j < 9; j++ {
			if board[i][j] >= '1' && board[i][j] <= '9' {
				if mm[0][board[i][j]-'1'] {
					return false
				}
				mm[0][board[i][j]-'1'] = true
			}

			if board[j][i] >= '1' && board[j][i] <= '9' {
				if mm[1][board[j][i]-'1'] {
					return false
				}
				mm[1][board[j][i]-'1'] = true
			}

			celli, cellj := j/3+(i/3)*3, j%3+(i%3)*3
			if board[celli][cellj] >= '1' && board[celli][cellj] <= '9' {
				if mm[2][board[celli][cellj]-'1'] {
					return false
				}
				mm[2][board[celli][cellj]-'1'] = true
			}

		}
	}

	return true
}

func TestIsValidSudoku(t *testing.T) {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			celli, cellj := j/3+(i/3)*3, j%3+(i%3)*3

			t.Log(celli, cellj)
		}
	}
}
