package leetcode

import "testing"

func generateMatrix(n int) [][]int {

	top, bottom, left, right := 0, n-1, 0, n-1

	ans := make([][]int, n)

	for i := range ans {
		ans[i] = make([]int, n)
	}

	counter := 1
	for top <= bottom && left <= right {

		for i := left; i <= right; i++ {
			ans[top][i] = counter
			counter++
		}
		top++
		for i := top; i <= bottom; i++ {
			ans[i][right] = counter
			counter++
		}

		right--

		for i := right; i >= left; i-- {
			ans[bottom][i] = counter
			counter++
		}

		bottom--

		for i := bottom; i >= top; i-- {
			ans[i][left] = counter
			counter++
		}

		left++

	}

	return ans
}

func TestGenerateMatrix(t *testing.T) {
	t.Log(generateMatrix(3))
}
