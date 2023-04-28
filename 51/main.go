package main

import "fmt"

func main() {

	//solveNQueens(1)
	fmt.Println(solveNQueens(4))
	//solveNQueens(4)

}

func solveNQueens(n int) [][]string {
	table := make([][]byte, n)
	for i := range table {
		table[i] = make([]byte, n)
		for j := range table[i] {
			table[i][j] = '.'
		}
	}

	check := func(r, c int) bool {

		if table[r][c] == 'Q' {
			return false
		}

		for j := 0; j < n; j++ {
			if table[r][j] == 'Q' {
				return false
			}
		}

		for i := 0; i < n; i++ {
			if table[i][c] == 'Q' {
				return false
			}
		}

		// 左边的斜边
		for i := 1; r-i >= 0 && c-i >= 0; i++ {
			if table[r-i][c-i] == 'Q' {
				return false
			}
		}

		// 右边的斜边
		for i := 1; r-i >= 0 && c+i < n; i++ {
			if table[r-i][c+i] == 'Q' {
				return false
			}
		}
		return true
	}

	var ans [][]string

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			// 到达最后一行了

			t := make([]string, 0, n)
			for i := range table {
				t = append(t, string(table[i]))
			}
			ans = append(ans, t)
			return
		}

		for col := 0; col < n; col++ {
			if check(row, col) {
				table[row][col] = 'Q'
				dfs(row + 1)
				table[row][col] = '.'
			}
		}
	}
	dfs(0)
	return ans
}

func solveNQueens2(n int) [][]string {

	table := make([][]byte, n)
	for i := range table {
		table[i] = make([]byte, n)
		for j := range table[i] {
			table[i][j] = '.'
		}
	}

	rowMap := make(map[int]bool)
	colMap := make(map[int]bool)

	var dfs func(ii, jj, q int)
	dfs = func(ii, jj, q int) {
		if q == n {
			for i := range table {
				fmt.Println(string(table[i]))
			}
			fmt.Println("---")
			return
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i < n && j < n && table[i][j] == '.' && !rowMap[i] && !colMap[j] {
					table[i][j] = 'Q'
					rowMap[i] = true
					colMap[j] = true
					dfs(i, j+1, q+1)
					dfs(i+1, j, q+1)
					table[i][j] = '.'
					rowMap[i] = false
					colMap[j] = false
				}
			}
		}
	}

	dfs(0, 0, 0)

	return nil

}
