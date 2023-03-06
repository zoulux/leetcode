package main

import (
	"fmt"
	"strconv"
)

func main() {

	//fmt.Println(translateNum(12258))
	fmt.Println(translateNum(322))
	fmt.Println(translateNum(444))
	//fmt.Println(translateNum(25))
}

func translateNum(num int) int {
	nums := strconv.Itoa(num)
	dp := make([]int, len(nums)+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len(nums); i++ {
		if nums[i-2] == '1' || (nums[i-2] == '2' && nums[i-1] < '6') {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)]

}
