package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	fmt.Println(merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}))

	fmt.Println(merge([][]int{
		{1, 4},
		{4, 5},
	}))

	fmt.Println(merge([][]int{
		{1, 4},
		//{5, 6},
	}))
}

func merge(intervals [][]int) [][]int {
	// 	第一种情况
	//	------
	// 	----

	// 第二种情况
	// -----
	// -------

	// 第三种情况
	//  ----
	//        ----

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{
		intervals[0],
	}

	for i := 1; i < len(intervals); i++ {
		top := ans[len(ans)-1]
		cur := intervals[i]

		if cur[0] <= top[1] {
			if cur[1] > top[1] {
				ans[len(ans)-1][1] = cur[1]
			}
		} else {
			ans = append(ans, cur)
		}
	}

	return ans
}

// 判断当前数组是否不需要合并
func isNoMerge(intervals [][]int) bool {

	// 长度为 0，1 ，是不需要合并的
	if len(intervals) <= 1 {
		return true
	}

	// 获取这个数组里面，最小值和最大值
	min, max := math.MaxInt, math.MinInt
	for _, v := range intervals {
		for _, vv := range v {
			if vv > max {
				max = vv
			}

			if vv < min {
				min = vv
			}
		}
	}

	// 开辟一个一维数组空间，可以存下 max 就好
	arr := make([]int, max+1)

	for _, v := range intervals {
		for i := v[0]; i <= v[1]; i++ {
			// 看看有多少点落在上面
			arr[i]++
		}
	}

	// 默认不需要合并
	flag := true
	for i := min; i <= max; i++ {
		if arr[i] > 1 {
			// 一旦大于 1 ，说明还有重叠数组
			flag = false
		}
	}

	return flag
}

func merge2(intervals [][]int) [][]int {
	if isNoMerge(intervals) {
		return intervals
	}

	// 先排个序，好比较，因为按左端点排序，所以需要判断的合并条件就两种
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		a := intervals[i]
		j := i + 1
		if j < len(intervals) {
			b := intervals[j]
			// 拿笔画画，两个线段的关系， b[0] <= a[1] 差点丢掉了
			if a[0] <= b[0] && a[1] <= b[1] && b[0] <= a[1] {
				// 线段 b    ------
				// 线段 a  ------
				ans = append(ans, []int{a[0], b[1]})
				i++ // 两个线段合并了，所以前进 1 位

				continue
			} else if a[0] <= b[0] && a[1] >= b[1] {
				// 线段 b    ------
				// 线段 a   ----------
				ans = append(ans, []int{a[0], a[1]})
				i++ // 两个线段合并了，所以前进 1 位
				continue
			}
		}

		// 其他情况
		//默认把当前线段加进去
		ans = append(ans, []int{a[0], a[1]})
	}
	// 圆哥说 合并的线段还可能有重叠部分，递归就完事了
	return merge(ans)
}
