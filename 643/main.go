package main

import "fmt"

func main() {

	fmt.Println(findMaxAverage([]int{
		0, 1, 1, 3, 3,
	}, 4))

}

// 1,12,-5,-6,50,3

func findMaxAverage(nums []int, k int) float64 {

	//win := make([]int, 0, k)
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
		//win = append(win, nums[i])
	}

	mx := sum
	for i := 1; i+k <= len(nums); i++ {
		sum += nums[i+k-1]
		sum -= nums[i-1]

		if sum > mx {
			mx = sum
		}
	}
	return float64(mx) / float64(k)
}
