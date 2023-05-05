package main

import "sort"

func main() {

}

func eraseOverlapIntervals(intervals [][]int) int {

	// 按左端点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	right := intervals[0][1]
	ans := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= right {
			ans++
			right = intervals[i][1]
		}
	}

	return len(intervals) - ans
}

// 会超时
func eraseOverlapIntervals2(intervals [][]int) int {

	// 按左端点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 以 i 左端点结尾的区间，最多多少个连续的
	dp := make([]int, len(intervals))
	for i := range dp {
		dp[i] = 1
	}

	max := func(arr []int) int {
		x := arr[0]
		for _, v := range arr[1:] {
			if v > x {
				x = v
			}
		}
		return x
	}

	for i := 1; i < len(intervals); i++ {
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				// j 的右端点小于 i 的左端点
				dp[i] = max([]int{dp[i], dp[j] + 1})
			}
		}
	}

	// 去掉的= 总长度 - 最长连续的
	return len(intervals) - max(dp)
}
