package leetcode

import (
	"fmt"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(string(rune(board[i][j])))
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func solveSudoku(board [][]byte) {
	var solve func(index int) bool

	solve = func(index int) bool {
		if index%9 >= 9 {
			return solve(index + 1)
		}

		if index == 80 {
			return true
		}

		if board[index/9][index%9] != '.' {
			return solve(index + 1)
		}

		for ch := byte('1'); ch <= '9'; ch++ {
			if !isValidSudoku2(board, index/9, index%9, ch) {
				continue
			}

			board[index/9][index%9] = ch
			if solve(index + 1) {
				return true
			}

			board[index/9][index%9] = '.'
		}
		return false
	}
	solve(0)
}

func isValidSudoku2(board [][]byte, r, c int, num byte) bool {

	for i := 0; i < 9; i++ {
		if tmp := board[r][i]; tmp == num {
			return false
		}

		if tmp := board[i][c]; tmp == num {
			return false
		}

		if tmp := board[i/3+(r/3)*3][i%3+(c/3)*3]; tmp == num {
			return false
		}
	}

	return true
}
