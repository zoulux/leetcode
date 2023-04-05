// leetcode submit region begin(Prohibit modification and deletion)
package main

import "fmt"

func main() {
	fmt.Println(sortArray([]int{
		//5, 1, 1, 2, 0, 0,
		-4, 0, 7, 4, 9, -5, -1, 0, -7, -1,
	}))
}

func sortArray(nums []int) []int {

	var sort func(left, right int)
	sort = func(low, height int) {
		left, right := low, height
		if left >= right {
			return
		}

		provit := nums[(left+right)/2]

		for left <= right {
			for nums[left] < provit {
				left++
			}

			for nums[right] > provit {
				right--
			}

			if left <= right {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right--
			}
		}

		sort(low, right)
		sort(left, height)
	}

	sort(0, len(nums)-1)

	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
