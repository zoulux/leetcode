package main

import (
	"fmt"
	"index/suffixarray"
	"sort"
	"unsafe"
)

func main2() {
	fmt.Println(lastSubstring("cacacb"))

	//fmt.Println(lastSubstring("aaaaa"))
	fmt.Println(lastSubstring("abab"))
	fmt.Println(lastSubstring("leetcode"))
	fmt.Println(lastSubstring("cacacb"))
	fmt.Println(lastSubstring("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}

func lastSubstringxx(s string) string {
	n := len(s)

	// 初始化指针i和j
	i, j, k := 0, 1, 0

	// 双指针遍历字符串
	for j < n && j+k < n {
		if s[i+k] == s[j+k] {
			// 如果相等，则继续比较后面的字符
			k++
		} else if s[i+k] > s[j+k] {
			// 如果s[i+k]大于s[j+k]，则最大子串在s[i+1:j+1]中
			j = j + k + 1
			k = 0
		} else {
			// 如果s[i+k]小于s[j+k]，则最大子串在s[j:n]中
			i = j
			j = j + 1
			k = 0
		}
	}

	// 返回最大子串
	return s[i:]
}

func main() {
	s := "banana"
	last := lastSubstringxx(s)
	fmt.Println(last) // 输出：nana
}

func lastSubstring(s string) string {
	check := func(ln int, t string) (string, bool) {
		x := ""
		for i := 0; i < len(s)-ln; i++ {
			sub := s[i : i+ln+1]
			if sub > x {
				x = sub
			}
		}

		if x > t {
			return x, true
		}
		return "", false
	}

	left, right := 0, len(s)-1
	var ans string
	for left <= right {
		mid := (left + right) >> 1
		if t, ok := check(mid, ans); ok {
			left = mid + 1 // 找到了，那就以mid +1 去试试，看看有没有更长一些的
			ans = t
		} else {
			right = mid - 1 // 如果没找到，那就找一个更短一点的字符串
		}
	}
	return ans
}

func lastSubstring6(s string) string {
	sa := (*struct {
		_  []byte
		sa []int32
	})(unsafe.Pointer(suffixarray.New([]byte(s)))).sa
	return s[sa[len(s)-1]:]
}

func lastSubstring5(s string) string {
	sa := (*struct {
		_  []byte
		sa []int32
	})(unsafe.Pointer(suffixarray.New([]byte(s)))).sa
	return s[sa[len(s)-1]:]
}

func lastSubstring4(s string) string {
	bs := []byte(s)

	var ans []byte
	for i := 0; i < len(bs); i++ {
		if sub := bs[i:]; string(sub) > string(ans) {
			ans = sub
		}
	}
	return string(ans)
}

func lastSubstring3(s string) string {

	maxIdx := 0
	equalIdx := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] > s[maxIdx] {
			maxIdx = i
			equalIdx = []int{i}
		} else if s[i] == s[maxIdx] {
			equalIdx = append(equalIdx, i)
		}
	}

	mm := make(map[string]bool)
	for i := 0; i < len(equalIdx); i++ {
		mm[s[equalIdx[i]:]] = true
	}
	sl := make([]string, 0, len(mm))
	for k := range mm {
		sl = append(sl, k)
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] > sl[j]
	})
	return sl[0]
}

// 超时
func lastSubstring2(s string) string {
	m := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			m[s[i:j]] = true
		}
	}

	sl := make([]string, 0, len(m))
	for k := range m {
		sl = append(sl, k)
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] > sl[j]
	})
	return sl[0]
}
