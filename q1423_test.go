package leetcode

import (
	"math"
	"testing"
)

func maxScore3(cardPoints []int, k int) int {
	n := len(cardPoints)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b

	}

	winSize := n - k
	sum := 0             // 数组总和
	winSum := 0          // 窗口内和
	ret := math.MaxInt64 // 最小和
	for i, v := range cardPoints {
		sum += v
		winSum += v

		if i >= winSize {
			winSum -= cardPoints[i-winSize]
		}

		if i >= winSize-1 {
			ret = min(winSum, ret)
		}
	}

	return sum - ret
}

func maxScore2(cardPoints []int, k int) int {
	n := len(cardPoints)
	preSum := make([]int, n+1)
	for i, v := range cardPoints {
		preSum[i+1] = preSum[i] + v
	}

	winSize := n - k

	ret := math.MaxInt64

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b

	}

	for i := 0; i+winSize <= n; i++ { // 窗口是左闭右闭 可以取到 n
		winSum := preSum[i+winSize] - preSum[i] // 窗口内和最小，那么两边和相加就是最大
		ret = min(winSum, ret)
	}

	return preSum[n] - ret
}

func maxScore(cardPoints []int, k int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(cardPoints)
	mm := map[[2]int]int{}

	var dfs func(i, j, track, k int) int
	dfs = func(i, j, track, k int) int {
		if k == 0 {
			return track
		}

		if v, ok := mm[[2]int{i, j}]; ok {
			return v
		}

		ret := max(dfs(i, j-1, track+cardPoints[j], k-1), dfs(i+1, j, track+cardPoints[i], k-1))
		mm[[2]int{i, j}] = ret
		return ret
	}

	return dfs(0, n-1, 0, k)
}

func TestMaxScore(t *testing.T) {

	// [2,2,2]
	// 2
	// [9,7,7,9,7,7,9]
	// 7
	// [1,1000,1]
	// 1
	// [1,79,80,1,1,1,200,1]
	// 3
	// t.Log(maxScore([]int{96, 90, 41, 82, 39, 74, 64, 50, 30}, 8))
	// t.Log(maxScore2([]int{96, 90, 41, 82, 39, 74, 64, 50, 30}, 8))
	t.Log(maxScore2([]int{1, 2, 3, 4, 5, 6, 1}, 3))
	t.Log(maxScore3([]int{1, 2, 3, 4, 5, 6, 1}, 3))

}
