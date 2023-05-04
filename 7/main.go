package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(letterCombinations("9"))
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))

}

func letterCombinations(digits string) (ans []string) {
	if digits == "" {
		return nil
	}

	var mm = []byte{0, 3, 6, 9, 12, 15, 19, 22, 26, 26}
	var dfs func(idx int, arr []byte)
	dfs = func(idx int, arr []byte) {
		if idx == len(digits) {
			ans = append(ans, string(arr))
			return
		}

		for i := mm[digits[idx]-'2']; i < mm[digits[idx]-'1']; i++ {
			dfs(idx+1, append(arr, i+'a'))
		}
	}
	dfs(0, make([]byte, 0, len(digits)))
	return
}

func letterCombinations4(digits string) []string {
	if digits == "" {
		return nil
	}

	mm := [][]byte{
		nil, nil,
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}

	var ans []string
	var dfs func(idx int, arr []byte)
	dfs = func(idx int, arr []byte) {
		if idx == len(digits) {
			ans = append(ans, string(arr))
			return
		}

		for _, v := range mm[digits[idx]-'0'] {
			dfs(idx+1, append(arr, v))
		}
	}
	dfs(0, []byte{})
	return ans
}

func letterCombinations3(digits string) []string {
	if digits == "" {
		return nil
	}

	mm := [][]string{
		nil,
		nil,
		strings.Split("abc", ""),
		strings.Split("def", ""),
		strings.Split("ghi", ""),
		strings.Split("jkl", ""),
		strings.Split("mno", ""),
		strings.Split("pqrs", ""),
		strings.Split("tuv", ""),
		strings.Split("wxyz", ""),
	}

	var ans []string

	var dfs func(idx int, arr []string)
	dfs = func(idx int, arr []string) {
		if idx == len(digits) {
			ans = append(ans, strings.Join(arr, ""))
			return
		}

		for _, v := range mm[digits[idx]-'0'] {
			dfs(idx+1, append(arr, v))
		}
	}
	dfs(0, []string{})
	return ans
}

func letterCombinations2(digits string) []string {
	if digits == "" {
		return nil
	}

	mm := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var ans []string

	var dfs func(idx int, arr []byte)
	dfs = func(idx int, arr []byte) {
		if idx == len(digits) {
			ans = append(ans, string(arr))
			return
		}

		for _, v := range mm[digits[idx]] {
			dfs(idx+1, append(arr, v))
		}
	}
	dfs(0, []byte{})
	return ans
}
