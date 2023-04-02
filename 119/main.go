package main

import "fmt"

func main() {

	//fmt.Println(generate(1))
	fmt.Println(getRow(3))
}

//你可以优化你的算法到 O(rowIndex) 空间复杂度吗？

func getRow(numRows int) []int {

	//$row = [];
	//$row[0] = 1;
	//for($i=1;$i<=$rowIndex;$i++){
	//$row[$i] = $row[$i-1] * ($rowIndex-$i +1)/$i;
	//}
	//return $row;

	dp := make([]int, numRows)

	dp[0] = 1
	dp[len(dp)-1] = 1

	for i := 1; i < numRows; i++ {
		dp[i] = dp[i-1] * (numRows - i + 1) / i
	}

	return dp
}
