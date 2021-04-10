package leetcode

import "testing"

type NumMatrix struct {
	matrixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	matrixSum := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		m := len(matrix[i])
		matrixSum[i] = make([]int, m)
		sum := 0
		for j := 0; j < m; j++ {
			sum += matrix[i][j]
			if i > 0 {
				matrixSum[i][j] = matrixSum[i-1][j] + sum
			} else {
				matrixSum[i][j] = sum
			}
		}
	}

	return NumMatrix{matrixSum}
}

func TestNumMatrixConstructor(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	numMatrix := Constructor(matrix)
	t.Log(numMatrix)
	t.Log(numMatrix.SumRegion(2, 1, 4, 3))
	t.Log(numMatrix.SumRegion(1, 1, 2, 2))
	t.Log(numMatrix.SumRegion(1, 2, 2, 4))
}

func (numMatrix *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	sum := numMatrix.matrixSum[row2][col2]

	if row1-1 >= 0 {
		sum -= numMatrix.matrixSum[row1-1][col2]
	}

	if col1-1 >= 0 {
		sum -= numMatrix.matrixSum[row2][col1-1]
	}

	if row1-1 >= 0 && col1-1 >= 0 {
		sum += numMatrix.matrixSum[row1-1][col1-1]
	}

	return sum

	// return   numMatrix.matrixSum[row2][col2]- numMatrix.matrixSum[row1-1][col2] - (numMatrix.matrixSum[row2][col1-1] - numMatrix.matrixSum[row1-1][col1-1])
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
