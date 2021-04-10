package leetcode

import "testing"

func spiralOrder(matrix [][]int) []int {

	i, j := 0, 0

	rowLen := len(matrix)
	colLen := len(matrix[0])

	visited := make([][]bool, rowLen)

	for k := range visited {
		visited[k] = make([]bool, colLen)
	}

	ans := []int{}

	for true {

		breakFlag := true
		for ; j < colLen && !visited[i][j]; j++ {
			breakFlag = false
			ans = append(ans, matrix[i][j])
			visited[i][j] = true

		}

		j--
		i++

		for ; i < rowLen && !visited[i][j]; i++ {
			breakFlag = false
			ans = append(ans, matrix[i][j])
			visited[i][j] = true

		}
		i--
		j--

		for ; j >= 0 && !visited[i][j]; j-- {
			breakFlag = false
			ans = append(ans, matrix[i][j])
			visited[i][j] = true
		}
		j++
		i--

		for ; i >= 0 && !visited[i][j]; i-- {
			breakFlag = false
			ans = append(ans, matrix[i][j])
			visited[i][j] = true
		}

		i++
		j++

		if breakFlag {
			break
		}
	}

	return ans
}

func TestSpiralOrder(t *testing.T) {
	// m := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	m := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	t.Log(spiralOrder(m))
}
