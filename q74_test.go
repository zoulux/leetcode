package leetcode

import (
	"testing"
)

func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	lefti, righti := 0, row

	for lefti < row && righti >= 0 && lefti <= righti {
		mid := (lefti + righti) / 2
		if target == matrix[mid][0] {
			return true
		} else if target < matrix[mid][0] {
			righti = mid - 1

		} else {
			lefti = mid + 1
		}
	}

	if righti < 0 || righti > row {
		return false
	}

	if righti == row {
		righti--
	}

	leftj, rightj := 0, col

	for rightj >= 0 && leftj < col && leftj <= rightj {
		mid := (leftj + rightj) / 2
		if target == matrix[righti][mid] {
			return true
		} else if target < matrix[righti][mid] {
			rightj = mid - 1

		} else {
			leftj = mid + 1
		}
	}

	return false
}

func TestSearchMatrix(t *testing.T) {
	// m := [][]int{
	// 	{1, 3, 5, 7},
	// 	{10, 11, 16, 20},
	// 	{23, 30, 34, 60},
	// 	{23, 30, 34, 60},
	// }

	m := [][]int{
		{1},
	}
	t.Log(searchMatrix(m, 10))
}
