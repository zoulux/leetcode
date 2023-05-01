package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(successfulPairs([]int{
		15, 8, 19,
	}, []int{
		38, 36, 23,
	}, 328))
	//
	//fmt.Println(successfulPairs([]int{
	//	5, 1, 3,
	//}, []int{
	//	1, 2, 3, 4, 5,
	//}, 7))

}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	suc := int(success)
	ln := len(potions)
	newPotions := make([]int, len(potions))

	type item struct {
		idx int
		val int
	}
	spellList := make([]item, len(spells))
	for i, v := range spells {
		spellList[i] = item{
			idx: i,
			val: v,
		}
	}

	sort.Slice(spellList, func(i, j int) bool {
		return spellList[i].val < spellList[j].val
	})

	ans := make([]int, len(spells))
	idx := 0
	for i, v := range spellList {
		idx = i
		for i, p := range potions {
			newPotions[i] = v.val * p
		}
		idx := sort.SearchInts(newPotions, suc)
		ans[v.idx] = ln - idx
		if ans[v.idx] == ln {
			break
		}
	}

	if spellList[idx].idx == ln {
		for i := idx + 1; i < len(spellList); i++ {
			ans[spellList[i].idx] = ln
		}
	}

	return ans
}
