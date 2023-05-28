package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(largestValsFromLabels([]int{
		5, 4, 3, 2, 1,
	}, []int{
		1, 1, 2, 2, 3,
	}, 3, 1))

	fmt.Println(largestValsFromLabels([]int{
		5, 4, 3, 2, 1,
	}, []int{
		1, 3, 3, 3, 2,
	}, 3, 2))
}

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	type item struct {
		value int
		label int
	}
	var items []item
	for i, v := range values {
		items = append(items, item{
			value: v,
			label: labels[i],
		})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})

	var ans int
	var count int
	mm := make(map[int]int)
	for _, v := range items {
		if count < numWanted && mm[v.label] < useLimit {
			ans += v.value
			count++
			mm[v.label]++
		}
	}
	return ans
}
