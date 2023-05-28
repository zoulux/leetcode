package main

import "fmt"

func main() {

	t := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(t)

	for i := range t {
		fmt.Println(string(t[i]))
	}
}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	var travel func(i, j int)
	travel = func(i, j int) {
		if i < 0 || i >= m {
			return
		}
		if j < 0 || j >= n {
			return
		}

		if board[i][j] == 'O' {
			board[i][j] = '-'
			travel(i+1, j)
			travel(i-1, j)
			travel(i, j+1)
			travel(i, j-1)
		}
	}

	// 从边出发，将 o变成-
	for i := 0; i < m; i++ {
		travel(i, 0)
		travel(i, n-1)
	}
	for j := 0; j < n; j++ {
		travel(0, j)
		travel(m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				// 没有覆盖的到 o 变成 x
				board[i][j] = 'X'
			} else if board[i][j] == '-' {
				// 覆盖到的 o 变回去 o
				board[i][j] = 'O'
			}
		}
	}
}
