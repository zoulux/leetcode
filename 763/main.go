package main

import "fmt"

func main() {

	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))

}

func partitionLabels(s string) []int {

	mm := make(map[byte]bool) // 处理过就不再处理
	var arr [][2]int          //拆分区间
	for i := 0; i < len(s); i++ {
		// 左边的坐标
		vi := s[i]

		if mm[vi] {
			continue
		}

		// 找右边的坐标
		var flag bool
		for j := len(s) - 1; j >= i; j-- {
			if s[j] == vi {
				flag = true
				arr = append(arr, [2]int{i, j})
			}
		}
		if !flag {
			arr = append(arr, [2]int{i, i})
		}

		// 当前字符处理过
		mm[vi] = true
	}

	// 下面就是线段合并区间的逻辑
	list := [][2]int{
		arr[0],
	}

	for i := 1; i < len(arr); i++ {
		top := list[len(list)-1]
		cur := arr[i]

		// 注意一下等于的情况，需要单独成段的
		if cur[0] < top[1] {
			if cur[1] > top[1] {
				list[len(list)-1][1] = cur[1]
			}
		} else {
			list = append(list, cur)
		}
	}

	// 右区间 - 左区间 +1 = 线段长度
	var ans []int
	for _, v := range list {
		ans = append(ans, v[1]-v[0]+1)
	}

	return ans
}
