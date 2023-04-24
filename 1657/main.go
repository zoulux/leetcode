package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println(closeStrings("cabbba", "aabbss"))
	fmt.Println(closeStrings("abc", "bca"))
	fmt.Println(closeStrings("a", "aa"))
	fmt.Println(closeStrings("cabbba", "abbccc"))
	fmt.Println(closeStrings("abbzzca", "babzzcz"))
	fmt.Println(closeStrings("cabbba", "abbccc"))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	// 不同字符出现的个数
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, v := range word1 {
		m1[v]++
	}

	for _, v := range word2 {
		m2[v]++
	}

	s1 := make([]int, 0)
	s2 := make([]int, 0)
	for k, v1 := range m1 {
		s1 = append(s1, v1)
		if _, ok := m2[k]; !ok {
			// 只能出现现有字符
			return false
		}
	}

	for k, v1 := range m2 {
		s2 = append(s2, v1)
		if _, ok := m1[k]; !ok {
			// 只能出现现有字符
			return false
		}
	}

	// 出现次数进行排序
	sort.Ints(s1)
	sort.Ints(s2)
	return reflect.DeepEqual(s1, s2) // 不同字符出现的次数必须相同
}
