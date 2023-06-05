package main

import "fmt"

func main() {
	//fmt.Println(3 == maxRepOpt1("ababa"))
	//fmt.Println(6 == maxRepOpt1("aaabaaa"))
	//fmt.Println(4 == maxRepOpt1("aaabbaaa"))
	fmt.Println(6 == maxRepOpt1("bbababaaaa")) //test15

}

func maxRepOpt1(text string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cnt := make(map[byte]int)
	for i := range text {
		// 统计各元素出现的次数
		cnt[text[i]]++
	}
	res := 0
	for cha, v := range cnt {
		// 各元素都去从头搜一遍

		// 经典双指针
		l, r := 0, 0
		not := 0
		for r < len(text) {
			if text[r] != cha {
				not++
			}
			r++
			for not >= 2 {
				if text[l] != cha {
					not--
				}
				l++
			}

			if r-l <= v {
				// 只有当前的长度小于总个数才进行计算
				res = max(r-l, res)
			}
		}
	}
	return res
}

func maxRepOpt12(text string) int {
	cnt := make([]int, 26)
	for i := range text {
		cnt[int(text[i]-'a')]++
	}
	res := 0
	for i := 0; i < 26; i++ {
		cha := byte('a' + i)
		l, r := 0, 0
		not := 0
		for r < len(text) {
			if text[r] != cha {
				not++
			}

			r++

			for not >= 2 {
				if text[l] != cha {
					not--
				}
				l++
			}

			if r-l <= cnt[i] {
				res = max(r-l, res)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
