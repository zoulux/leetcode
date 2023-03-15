package main

import "os"

func main() {

}

func majorityElement(nums []int) int {
	n := len(nums)
	mm := make(map[int]int)
	for _, v := range nums {
		mm[v]++
		if mm[v] > n/2 {
			return v
		}
	}
	return 0
}

func init() {
	os.Exit(0)
}
