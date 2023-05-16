package main

import "fmt"

func main() {

	fmt.Println(findKthLargest([]int{
		3, 2, 1, 5, 6, 4,
	}, 2))

	fmt.Println(findKthLargest([]int{
		3, 2, 3, 1, 2, 4, 5, 5, 6,
	}, 4))
}

func findKthLargest(nums []int, k int) int {
	var quick func(low, right int)
	quick = func(low, height int) {
		if low >= height {
			return
		}

		left, right := low, height
		p := nums[right]

		for left < right {
			for left < right && nums[left] <= p {
				left++
			}
			nums[left], nums[right] = nums[right], nums[left]

			for left < right && nums[right] >= p {
				right--
			}
			nums[left], nums[right] = nums[right], nums[left]
		}

		if left < len(nums)-k {
			quick(left+1, height)
		} else if left > len(nums)-k {
			quick(low, left-1)
		} else {
			return
		}
	}

	quick(0, len(nums)-1)
	return nums[len(nums)-k]
}
