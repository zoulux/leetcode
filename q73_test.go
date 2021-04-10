package leetcode

import "testing"



func setZeroes2(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])

	zeroRow := make([]bool, row)
	zeroCol := make([]bool, col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				zeroRow[i] = true
				zeroCol[j] = true
			}
		}
	}

	for i := 0; i < row; i++ {

		if !zeroRow[i] {
			continue
		}
		for j := 0; j < col; j++ {
			if !zeroCol[j] {
				continue
			}

			for m := 0; m < row; m++ {
				matrix[m][j] = 0
			}

			for n := 0; n < col; n++ {
				matrix[i][n] = 0
			}
		}
	}
}

func setZeroes(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])

	zeroMatrix := make([][]bool, row)

	for i := 0; i < row; i++ {
		zeroMatrix[i] = make([]bool, col)
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				zeroMatrix[i][j] = true
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if zeroMatrix[i][j] {
				for m := 0; m < row; m++ {
					matrix[m][j] = 0
				}

				for n := 0; n < col; n++ {
					matrix[i][n] = 0
				}
			}
		}
	}
}

func TestSetZeroes(t *testing.T) {

	m := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes2(m)

	for i := 0; i < len(m); i++ {
		t.Log(m[i])
	}
}
