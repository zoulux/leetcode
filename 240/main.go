package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	for row, col := 0, len(matrix[0])-1; row < len(matrix) && col >= 0; {
		if target > matrix[row][col] {
			row++
		} else if target < matrix[row][col] {
			col--
		} else {
			return true
		}
	}
	return false
}
