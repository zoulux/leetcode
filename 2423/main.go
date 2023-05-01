package main

import "fmt"

func main() {

	fmt.Println(equalFrequency("abcc"))
	fmt.Println(equalFrequency("bac"))
	fmt.Println(equalFrequency("ddaccb"))
	fmt.Println(equalFrequency("aca"))
}

func equalFrequency(word string) bool {
	bs := []byte(word)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		m := make(map[byte]int)
		for i := 0; i < len(bs); i++ {
			if i != idx {
				m[bs[i]]++
			}
		}

		mm := make(map[int]int)

		for _, v := range m {
			mm[v]++
		}

		if len(mm) == 1 {
			return true
		}

		return false
	}

	for i := 0; i < len(bs); i++ {
		if dfs(i) {
			return true
		}
	}

	return false
}
