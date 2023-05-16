package main

import (
	"fmt"
	"sort"
)

func main() {

	//fmt.Println(rearrangeBarcodes([]int{
	//	51, 83, 51, 40, 51, 40, 51, 40, 83, 40, 83, 83, 51, 40, 40, 51, 51, 51, 40, 40, 40, 83, 51, 51, 40, 51, 51, 40, 40, 51, 51, 40, 51, 51, 51, 40, 83, 40, 40, 83, 51, 51, 51, 40, 40, 40, 51, 51, 83, 83, 40, 51, 51, 40, 40, 40, 51, 40, 83, 40, 83, 40, 83, 40, 51, 51, 40, 51, 51, 51, 51, 40, 51, 83, 51, 83, 51, 51, 40, 51, 40, 51, 40, 51, 40, 40, 51, 51, 51, 40, 51, 83, 51, 51, 51, 40, 51, 51, 40, 40,
	//}))
	//
	//fmt.Println(rearrangeBarcodes([]int{
	//	1, 1, 1, 2, 2, 2,
	//}))

	fmt.Println(rearrangeBarcodes([]int{
		2, 1, 1,
	}))

	fmt.Println(rearrangeBarcodes([]int{
		1, 1, 1, 1, 2, 2, 3, 3, 3,
	}))
}

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)
	if n < 3 {
		return barcodes
	}
	counts := map[int]int{}
	for _, v := range barcodes {
		counts[(v)]++
	}

	evenIdx, oddIdx := 0, 1
	res := make([]int, n)
	halfN := n / 2

	for i, c := range counts {
		v := int(i)

		for c > 0 && oddIdx < n && c <= halfN {
			res[oddIdx] = v
			oddIdx += 2
			c--
		}

		for c > 0 && evenIdx < n {
			res[evenIdx] = v
			evenIdx += 2
			c--
		}

	}
	return res
}

func rearrangeBarcodes2(barcodes []int) []int {
	mm := make(map[int]int)
	for _, b := range barcodes {
		mm[b]++
	}

	sort.Slice(barcodes, func(i, j int) bool {
		a, b := barcodes[i], barcodes[j]
		if mm[a] == mm[b] {
			return a < b
		}
		return mm[a] > mm[b]
	})

	newbarcodes := make([]int, len(barcodes))

	j := 0
	for i := 0; i < len(barcodes); i += 2 {
		newbarcodes[i] = barcodes[j]
		j++
	}

	for i := 1; i < len(barcodes); i += 2 {
		newbarcodes[i] = barcodes[j]
		j++
	}

	return newbarcodes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func rearrangeBarcodesErr(barcodes []int) []int {

	mm := make(map[int]int)
	for _, b := range barcodes {
		mm[b]++
	}

	type item struct {
		k     int
		count int
	}
	var items []item

	for k, v := range mm {
		if len(items) == 0 {
			items = append(items, item{
				k:     k,
				count: v,
			})
		} else {
			j := sort.Search(len(items), func(i int) bool {
				return items[i].count >= v
			})
			items = append(items[:j], append([]item{{
				k:     k,
				count: v,
			}}, items[j:]...)...)
		}
	}

	n := len(barcodes)

	barcodes = make([]int, 0, len(barcodes))

	for len(barcodes) < n && len(items) > 0 {
		top := items[len(items)-1]
		barcodes = append(barcodes, top.k)
		items[len(items)-1].count--

		if len(items)-2 >= 0 {
			next := items[len(items)-2]
			barcodes = append(barcodes, next.k)
			items[len(items)-2].count--
		}

		for items[len(items)-1].count == 0 {
			items = items[:len(items)-1]
		}

	}

	//fmt.Println(items)

	//1,4
	// 2,2
	// 3,2

	//112

	return barcodes
}
