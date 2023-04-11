package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}) == 167)
}

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。
func maxCoins(nums []int) int {
	n := len(nums)

	points := append([]int{1}, nums...)
	points = append(points, 1)

	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := n; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(
					dp[i][j],
					dp[i][k]+dp[k][j]+points[i]*points[k]*points[j],
				)
			}
		}
	}

	return dp[0][n+1]
}

func maxCoins4(nums []int) int {
	n := len(nums)
	// 添加两侧的虚拟气球
	points := make([]int, 0, n+2)
	points = append([]int{1}, nums...)
	points = append(points, 1)
	// base case 已经都被初始化为 0
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	// 开始状态转移
	// i 应该从下往上
	for i := n; i >= 0; i-- {
		// j 应该从左往右
		for j := i + 1; j < n+2; j++ {
			// 最后戳破的气球是哪个？
			for k := i + 1; k < j; k++ {
				// 择优做选择
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}
	// 因为最终i=0，所以从最大值向下变小
	// j 等于
	return dp[0][n+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxCoins3(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var solve func(i, j int) int
	solve = func(l, r int) int {
		if l+2 > r {
			return 0
		} else if l+2 == r {
			return nums[l] * nums[l+1] * nums[l+2]
		}

		res := 0
		for k := l + 1; k <= r-1; k++ {
			left := solve(l, k)
			right := solve(k, r)

			sum := left + right + nums[k]*nums[l]*nums[r]
			res = max(res, sum)
		}
		return res
	}
	return solve(0, len(nums)-1)

}
func maxCoins2(nums []int) int {

	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// ln 表示开区间的长度
	for ln := 3; ln <= len(nums); ln++ {
		//l 表示开区间左端点,l + len - 1则表示开区间右端点
		for i := 0; i < len(nums)-ln; i++ {
			//k为开区间内的索引（代表区间最后一个被戳破的气球）
			for k := i + 1; k < i+ln-1; k++ {
				left := dp[i][k]
				right := dp[k][i+ln-1]
				sum := nums[i] * nums[k] * nums[i+ln-1]
				dp[i][i+ln-1] = max(dp[i][i+ln-1], left+sum+right)
			}
		}
	}
	return dp[0][len(nums)-2]
}
