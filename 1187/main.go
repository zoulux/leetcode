package main

import (
	"fmt"
	"math"
	"sort"
)

func lengthOfLIS(arr []int) int {
	dp := make([]int, len(arr))
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	mx := 0
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > mx {
			mx = dp[i]
		}
	}

	return mx

}
func main() {
	//
	//fmt.Println(lengthOfLIS([]int{
	//	10, 9, 2, 5, 3, 7, 101, 18,
	//}))

	//fmt.Println(makeArrayIncreasing(
	//	[]int{
	//		1, 5, 3, 6, 7,
	//	}, []int{
	//		1, 3, 2, 4,
	//	},
	//))

	fmt.Println(makeArrayIncreasing(
		[]int{
			1, 28, 27, 14, 5, 16, 31, 2, 9, 4, 3, 22, 13, 24, 7, 10,
		}, []int{
			12, 11, 30, 21, 0, 15, 18, 25, 20, 19, 6, 29, 8, 23, 26, 1, 28,
		},
	))

	//fmt.Println(makeArrayIncreasing(
	//	[]int{
	//		1, 5, 3, 6, 7,
	//	}, []int{
	//		1, 6, 3, 3,
	//	},
	//))
}

func makeArrayIncreasing(a, b []int) int {
	b = unique(b) // 对数组进行去重

	cache := make(map[[2]int]int) // 记忆化搜索 ，go 里面数组可以作为 key
	var dfs func(idx int, pre int) int
	dfs = func(idx int, pre int) int {
		if idx < 0 {
			return 0
		}

		if v, ok := cache[[...]int{idx, pre}]; ok {
			return v
		}

		res := math.MaxInt / 2
		if a[idx] < pre {
			// 如果小 那就往左搜一下试试
			// 其实这里，可以更换数字，也可以不替换，需要取最小值
			res = min(res, dfs(idx-1, a[idx]))
		}

		//否则肯定要替换
		k := sort.SearchInts(b, pre) - 1 // 去 b 中找 pre，但是我们需要找一个比 pre 小的值
		if k >= 0 {
			res = min(res, dfs(idx-1, b[k])+1) // 找到了，继续往左搜
		}
		// 否则未找到
		cache[[...]int{idx, pre}] = res
		return res
	}

	t := dfs(len(a)-1, math.MaxInt/2)
	if t >= math.MaxInt/2 {
		return -1
	}
	return t
}

func unique(arr []int) []int {
	idx := 0
	for _, v := range arr[1:] {
		if arr[idx] != v {
			idx++
			arr[idx] = v
		}
	}
	arr = arr[:idx+1]
	sort.Ints(arr)
	return arr
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func makeArrayIncreasing222(a, b []int) int {
	a = append(a, math.MaxInt) // 简化代码逻辑
	sort.Ints(b)
	m := 0
	for _, x := range b[1:] {
		if b[m] != x {
			m++
			b[m] = x // 原地去重
		}
	}
	b = b[:m+1]

	n := len(a)
	f := make([]int, n)
	for i, x := range a {
		k := sort.SearchInts(b, x)
		res := 0 // 小于 a[i] 的数全部替换
		if k < i {
			res = math.MinInt
		}
		if i > 0 && a[i-1] < x { // 无替换
			res = max(res, f[i-1])
		}
		for j := i - 2; j > i-k-1 && j >= 0; j-- {
			if b[k-(i-j-1)] > a[j] {
				// a[j+1] 到 a[i-1] 替换成 b[k-(i-j-1)] 到 b[k-1]
				res = max(res, f[j])
			}
		}
		f[i] = res + 1 // 把 +1 移到这里，表示 a[i] 不替换
	}
	if f[n-1] < 0 { // 注意 a 已经添加了一个元素
		return -1
	}
	return n - f[n-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func has(idx int, arr1, arr2 []int, m map[int]bool) bool {
	for j := 0; j < len(arr2); j++ {
		if !m[j] && arr2[j] > arr1[idx-1] && arr2[j] <= arr1[idx+1] {
			arr1[idx] = arr2[j]
			m[j] = true
			return true
		}
	}

	return false
}

func makeArrayIncreasing2(arr1 []int, arr2 []int) int {

	lol := lengthOfLIS(arr1)
	if lol == len(arr1) {
		return 0
	}

	sort.Ints(arr2)

	visitArr2 := make(map[int]bool) //
	var ans int

	for l := 1; l < len(arr1); l++ {
		for ; l < len(arr1); l++ {
			if arr1[l-1] >= arr1[l] {
				break
			}
		}

		if l < len(arr1) {
			if has(l-1, arr1, arr2, visitArr2) {
				ans++
				l--
			} else {
				return -1
			}
		}

	}

	return ans

}
