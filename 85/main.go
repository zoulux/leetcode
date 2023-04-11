package main

import "fmt"

func main() {

	fmt.Println(maximalRectangle([][]byte{
		{'1', '0', '1', '1', '1'},
		{'0', '1', '0', '1', '0'},
		{'1', '1', '0', '1', '1'},
		{'1', '1', '0', '1', '1'},
		{'0', '1', '1', '1', '1'},
	}))

}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	width := make([][]int, m)

	for i := range width {
		width[i] = make([]int, n)
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	area := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				width[i][j] = 1
				if j-1 >= 0 {
					width[i][j] += width[i][j-1]
				}
			}

			minWidth := width[i][j]
			for row := i; row >= 0; row-- {
				minWidth = min(minWidth, width[row][j])

				if minWidth == 0 {
					break
				}

				height := i - row + 1
				area = max(area, minWidth*height)

				//fmt.Printf("area %d , row %d minWidth:%d height:%d \n", area, row, minWidth, height)
			}

		}
	}

	fmt.Println(width)

	return area
}
