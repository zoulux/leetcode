package main

import "fmt"

func main() {
	//[3,3,5,5,6,7]
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	dq := make([]int, 0)

	push := func(j int) {
		for len(dq) > 0 && nums[dq[len(dq)-1]] < nums[j] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, j)
	}

	pop := func() {
		dq = dq[1:]
	}

	var ans []int
	for right := 0; right < len(nums); right++ {
		push(right)
		left := right - k + 1
		if dq[0] < left {
			pop()
		}

		if right+1 >= k {
			ans = append(ans, nums[dq[0]])
		}
	}
	return ans
}
func maxSlidingWindow2(nums []int, k int) []int {
	ans := make([]int, 0)
	for i := 0; i <= len(nums)-k; i++ {
		v := nums[i]
		for j := i + 1; j < i+k; j++ {
			if nums[j] > v {
				v = nums[j]
			}
		}
		ans = append(ans, v)
	}
	return ans
}
