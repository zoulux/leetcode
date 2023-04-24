package main

import "sync"

func main() {

}

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)

	for _, v := range nums1 {
		m1[v] = true
	}

	for _, v := range nums2 {
		m2[v] = true
	}

	ans := make([][]int, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for k := range m1 {
			if !m2[k] {
				ans[0] = append(ans[0], k)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for k := range m2 {
			if !m1[k] {
				ans[1] = append(ans[1], k)
			}
		}
	}()
	wg.Wait()

	return ans
}
