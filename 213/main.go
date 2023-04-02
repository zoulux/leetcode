package main

import "fmt"

func main() {

	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))

}

func rob(nums []int) int {

	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	_rob := func(nums []int) int {
		dp := make([]int, len(nums))
		// 在第 i 间屋子能获得最大的收益
		//1,2,3,1
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			dp[i] = max(
				dp[i-2]+nums[i], // 偷当前房间
				dp[i-1],         // 不偷当前房间
			)
		}
		return dp[len(nums)-1]
	}

	t1 := _rob(nums[:len(nums)-1])
	t2 := _rob(nums[1:])

	return max(t1, t2)
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
