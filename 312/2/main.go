package main

import "fmt"

func main() {

	arr := []int{2, 3}
	a2 := append(arr, 4, 5, 6, 7, 8)
	fmt.Println(cap(a2))

	fmt.Println(maxCoins([]int{
		1, 5,
	}))

	fmt.Println(maxCoins([]int{
		3, 1, 5, 8,
	}))
}
func maxCoins(nums []int) int {
	nums = append([]int{1}, append(nums, 1)...)

	n := len(nums)
	dp := make([][]int, n) // 表示 i j 开区间内获取的金币数量
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- { // 因为最终要从0到n
		for j := i + 2; j < n; j++ { //j其实可以从 i+1,但是这个时候，k 就没有位置了，i+2 给 k 罗位置
			for k := i + 1; k < j; k++ { // k 在 i 和 j 之间
				dp[i][j] = max(dp[i][j], dp[i][k]+nums[i]*nums[k]*nums[j]+dp[k][j])
			}
		}
	}
	return dp[0][n-1]
}

func maxCoins2(nums []int) int {
	var dfs func(i, j int) int // 表示开区间获得的金币数量
	dfs = func(i, j int) (ret int) {
		mx := 0
		for k := i + 1; k < j; k++ {
			left, right := dfs(i, k), dfs(k, j)
			mx = max(mx, left+nums[i]*nums[k]*nums[j]+right)
		}
		return mx
	}

	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	return dfs(0, len(nums)-1)
}

func maxCoins3(nums []int) int {
	mm := make(map[[2]int]int)

	var dfs func(i, j int) int // 表示开区间获得的金币数量
	dfs = func(i, j int) (ret int) {
		if v, ok := mm[[2]int{i, j}]; ok {
			return v
		}
		defer func() {
			mm[[2]int{i, j}] = ret
		}()

		mx := 0
		for k := i + 1; k < j; k++ {
			left, right := dfs(i, k), dfs(k, j)
			cnt := nums[i] * nums[k] * nums[j]
			mx = max(mx, left+cnt+right)
		}
		return mx
	}

	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	return dfs(0, len(nums)-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
