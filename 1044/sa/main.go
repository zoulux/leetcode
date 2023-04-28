package main

import (
	"fmt"
	"index/suffixarray"
	"reflect"
	"unsafe"
)

func main() {
	//fmt.Println(longestDupSubstring("aa"))
	fmt.Println(longestDupSubstring("banana"))
	//fmt.Println(longestDupSubstring("abcd"))
}

func longestDupSubstring(s string) string {
	// 求出后缀数组和高度数组
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	// 高度数组中的最大值对应的就是最长的重复子串
	idx, maxH := 0, 0
	for i, h := range height {
		if h > maxH {
			idx, maxH = i, h
		}
	}
	return s[sa[idx] : int(sa[idx])+maxH]
}

func longestDupSubstring2(s string) string {

	sa := suffixarray.New([]byte(s))

	check := func(ln int) (string, bool) {
		for i := 0; i < len(s)-ln; i++ {
			sub := s[i : i+ln+1]
			idxs := sa.Lookup([]byte(sub), 2)
			if len(idxs) >= 2 {
				return sub, true
			}

		}
		return "", false
	}

	left, right := 0, len(s)-1
	var ans string
	for left <= right {
		mid := (left + right) >> 1
		if t, ok := check(mid); ok {
			left = mid + 1 // 找到了，那就以mid +1 去试试，看看有没有更长一些的
			ans = t
		} else {
			right = mid - 1 // 如果没找到，那就找一个更短一点的字符串
		}
	}
	return ans
}
