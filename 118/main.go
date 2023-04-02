package main

import "fmt"

func main() {

	//fmt.Println(generate(1))
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
		left, right := 0, i
		dp[i][left] = 1
		dp[i][right] = 1

		left++
		right--
		for left <= right {

			dp[i][left] = dp[i-1][left]
			if left-1 >= 0 {
				dp[i][left] += dp[i-1][left-1]
			}

			if left != right {
				dp[i][right] = dp[i-1][right]
				if right-1 >= 0 {
					dp[i][right] += dp[i-1][right-1]
				}
			}

			left++
			right--
		}
	}

	return dp
}
