package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{
		//100, 4, 200, 1, 3, 2,
		0, 3, 7, 2, 5, 8, 4, 6, 0, 1,
	}))

}

func longestConsecutive(nums []int) int {
	//nums = append(nums, nums...)
	mm := make(map[int]int)

	max := 0
	for _, v := range nums {
		if _, ok := mm[v]; ok {
			continue
		}

		j := 1
		l, r := mm[v-1], mm[v+1]

		if c, ok := mm[v]; ok {
			j = c
		}

		j = 1 + l + r

		if j > max {
			max = j
		}
		mm[v] = j
		mm[v-l] = j
		mm[v+r] = j

	}

	return max
}
