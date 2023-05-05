package main

import "sort"

func main() {

}

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	right := points[0][1]
	ans := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > right {
			ans++
			right = points[i][1]
		}
	}
	return ans
}
