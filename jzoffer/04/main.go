package main

import "fmt"

func main() {
	//
	//fmt.Println(findNumberIn2DArray([][]int{
	//	{1, 4, 7, 11, 15},
	//	{2, 5, 8, 12, 19},
	//	{3, 6, 9, 16, 22},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30},
	//}, 5))
	//
	//fmt.Println(findNumberIn2DArray([][]int{
	//	{1, 4, 7, 11, 15},
	//	{2, 5, 8, 12, 19},
	//	{2, 5, 8, 12, 19},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30},
	//}, 20))
	//fmt.Println(findNumberIn2DArray([][]int{
	//	{1, 4, 7, 11, 15},
	//}, 11))
	//
	//fmt.Println(findNumberIn2DArray([][]int{
	//	{1},
	//	{4},
	//	{7},
	//	{11},
	//	{15},
	//}, 11))

	//fmt.Println(findNumberIn2DArray([][]int{
	//	{-5},
	//}, -10))

	fmt.Println(findNumberIn2DArray([][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}, 19))
}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		if k := matrix[row][col]; k > target {
			col--
		} else {
			row++
		}
	}

	return false
}

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	leftCol, rightCol := 0, len(matrix[0])-1
	for leftCol <= rightCol {
		mid := (leftCol + rightCol) >> 1
		if target > matrix[0][mid] {
			leftCol = mid + 1
		} else if target < matrix[0][mid] {
			rightCol = mid - 1
		} else {
			return true
		}
	}
	if rightCol < 0 {
		return false
	}

	leftRow, rightRow := 0, len(matrix)-1

	for leftRow <= rightRow {
		mid := (leftRow + rightRow) >> 1
		if target > matrix[mid][rightCol] {
			leftRow = mid + 1
		} else if target < matrix[mid][rightCol] {
			rightRow = mid - 1
		} else {
			return true
		}
	}
	return false
}
