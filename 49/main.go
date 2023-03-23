package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println(groupAnagrams([]string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}))
}

func sortStr(str string) string {
	a := strings.Split(str, "")
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	return strings.Join(a, "")
}

func groupAnagrams(strs []string) [][]string {
	mm := make(map[string][]string)
	for _, str := range strs {
		s := sortStr(str)
		mm[s] = append(mm[s], str)
	}

	var ans [][]string
	for _, v := range mm {
		ans = append(ans, v)
	}
	return ans
}
