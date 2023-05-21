package main

import "math"

func main() {

}

func maxProfit(prices []int) int {
	ans := 0
	min := math.MaxInt / 2
	for _, v := range prices {
		if v < min {
			min = v
		}

		if t := v - min; t > ans {
			ans = t

		}
	}

	return ans
}
