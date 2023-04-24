package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	//fmt.Println(fieldOfGreatestBlessing([][]int{{565, 34, 242}, {299, 628, 870}, {724, 673, 221}, {128, 267, 917}, {561, 488, 696}, {341, 71, 604}, {835, 92, 846}, {945, 696, 973}, {554, 776, 430}, {209, 134, 594}, {987, 898, 282}, {819, 154, 462}, {618, 946, 505}, {561, 40, 677}, {602, 863, 657}, {295, 83, 718}, {456, 920, 939}, {94, 717, 813}, {611, 946, 366}, {818, 850, 85}, {395, 554, 333}, {325, 700, 628}, {590, 401, 119}, {248, 858, 499}, {298, 723, 602}, {189, 538, 646}, {194, 283, 344}, {842, 535, 866}, {192, 832, 195}}))

	//fmt.Println(fieldOfGreatestBlessing([][]int{{932, 566, 342}, {546, 489, 250}, {723, 454, 748}, {830, 887, 334}, {617, 534, 721}, {924, 267, 892}, {151, 64, 65}, {318, 825, 196}, {102, 941, 940}, {748, 562, 582}, {76, 938, 228}, {921, 15, 245}, {871, 96, 823}, {701, 737, 991}, {339, 861, 146}, {484, 409, 823}, {574, 728, 557}, {104, 845, 459}, {363, 804, 94}, {445, 685, 83}, {324, 641, 328}, {626, 2, 897}, {656, 627, 521}, {935, 506, 956}, {210, 848, 502}, {990, 889, 112}}))

	fmt.Println(fieldOfGreatestBlessing([][]int{{932, 566, 342}, {546, 489, 250}, {723, 454, 748}, {830, 887, 334}, {617, 534, 721}, {924, 267, 892}, {151, 64, 65}, {318, 825, 196}, {102, 941, 940}, {748, 562, 582}, {76, 938, 228}, {921, 15, 245}, {871, 96, 823}, {701, 737, 991}, {339, 861, 146}, {484, 409, 823}, {574, 728, 557}, {104, 845, 459}, {363, 804, 94}, {445, 685, 83}, {324, 641, 328}, {626, 2, 897}, {656, 627, 521}, {935, 506, 956}, {210, 848, 502}, {990, 889, 112}}))
	fmt.Println(fieldOfGreatestBlessing([][]int{
		{0, 0, 1},
		{1, 0, 1},
	}))

	fmt.Println(fieldOfGreatestBlessing([][]int{
		{4, 4, 6},
		{7, 5, 3},
		{1, 6, 2},
		{5, 6, 3},
	}))

}

func fieldOfGreatestBlessing(forceField [][]int) int {
	xset := make(map[int]struct{})
	yset := make(map[int]struct{})
	for _, field := range forceField {
		x, y, side := field[0]*2, field[1]*2, field[2]
		xset[x-side] = struct{}{}
		xset[x+side] = struct{}{}
		yset[y-side] = struct{}{}
		yset[y+side] = struct{}{}
	}
	xlist := make([]int, 0, len(xset))
	ylist := make([]int, 0, len(yset))

	for i := range xset {
		xlist = append(xlist, i)
	}

	for i := range yset {
		ylist = append(ylist, i)
	}

	sort.Ints(xlist)
	sort.Ints(ylist)

	diff := make([][]int, len(xlist)+2)
	for i := range diff {
		diff[i] = make([]int, len(ylist)+2)
	}

	for _, field := range forceField {
		x, y, side := field[0]*2, field[1]*2, field[2]

		// x1,y1 为差分数组的左上角
		// x2，y2 为差分数组的右下角

		x1 := sort.SearchInts(xlist, x-side)
		x2 := sort.SearchInts(xlist, x+side)
		y1 := sort.SearchInts(ylist, y-side)
		y2 := sort.SearchInts(ylist, y+side)

		// x1+1 是为了下面求和方便运算
		// x2+2 => (x2+1)+1 为了排除x2以外的数，x2 本身的边是包含在内的
		diff[x1+1][y1+1]++
		diff[x2+2][y2+2]++
		diff[x1+1][y2+2]--
		diff[x2+2][y1+1]--
	}

	var ans int
	for i := 1; i < len(diff); i++ {
		for j := 1; j < len(diff[i]); j++ {
			diff[i][j] += diff[i-1][j] + diff[i][j-1] - diff[i-1][j-1]
			// 差分求和
			// 当前的值 + 左边数组的值 + 右边数组的值 - 重叠的值
			if diff[i][j] > ans {
				ans = diff[i][j]
			}
		}
	}

	return ans
}

func fieldOfGreatestBlessing3(forceField [][]int) int {
	covered := make(map[[2]int]int)
	for _, field := range forceField {
		x, y, side := field[0]*2, field[1]*2, field[2]
		for i := x; i <= x+side; i++ {
			for j := y; j <= y+side; j++ {
				covered[[2]int{i, j}]++
			}
		}
	}
	maxCount := 0
	for _, count := range covered {
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func fieldOfGreatestBlessing2(forceField [][]int) int {

	left, right := float64(math.MaxInt/2), float64(math.MinInt/2)
	top, down := float64(math.MinInt/2), float64(math.MaxInt/2)

	arr := make([][4]float64, 0, len(forceField)) // 预处理
	for _, v := range forceField {
		x1 := float64(v[0]) - float64(v[2])/2.0
		if x1 < left {
			left = x1
		}

		x2 := float64(v[0]) + float64(v[2])/2.0
		if x2 > right {
			right = x2
		}

		y1 := float64(v[1]) + float64(v[2])/2.0
		if y1 > top {
			top = y1
		}

		y2 := float64(v[1]) - float64(v[2])/2.0
		if y2 < down {
			down = y2
		}

		arr = append(arr, [...]float64{x1, x2, y1, y2})
	}

	if d := right - left; float64(int(d)) != d {
		// 是正数

		//0.5 5

		for i := range arr {
			arr[i][0] *= 10
			arr[i][1] *= 10
		}
		right *= 10
		left *= 10
	}

	if d := top - down; float64(int(d)) != d {
		// 是正数
		//0.5 5

		for i := range arr {
			arr[i][3] *= 10
			arr[i][2] *= 10
		}
		top *= 10
		down *= 10
	}

	if left != 0 {
		for i := range arr {
			arr[i][0] -= left
			arr[i][1] -= left
		}
		right -= left
		left = 0
	}

	if down != 0 {
		for i := range arr {
			arr[i][3] -= down
			arr[i][2] -= down
		}
		top -= down
		down = 0
	}

	// 将分布拖到第一象限
	area := make([][]int, int(right)+1)
	for i := range area {
		area[i] = make([]int, int(top)+1)
	}

	for _, v := range arr {
		for i := v[0]; i <= v[1]; i++ {
			for j := v[3]; j <= v[2]; j++ {
				area[int(i)][int(j)]++
			}
		}
	}

	var ans int
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[i]); j++ {
			if area[i][j] > ans {
				ans = area[i][j]
			}
		}
	}
	return ans
}
