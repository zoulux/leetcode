package main

import "fmt"

func main() {

	fmt.Println(numTilePossibilities("AAB") == 8)

}

func numTilePossibilities(tiles string) int {
	mm := make(map[byte]int)
	for i := 0; i < len(tiles); i++ {
		mm[tiles[i]]++
	}
	var dfs func(mm map[byte]int) int

	dfs = func(mm map[byte]int) int {

		var res int
		for k, v := range mm {
			if v == 0 {
				continue
			}
			res++
			mm[k]--
			res += dfs(mm)
			mm[k]++
		}
		return res
	}

	return dfs(mm)
}
