package leetcode

import "testing"

func maxArea(height []int) int {
	ans := -1
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > i; j-- {
			temp := 0
			if height[i] > height[j] {
				temp = height[j] * (j - i)
			} else {
				temp = height[i] * (j - i)
			}
			if temp > ans {
				ans = temp
			}
		}
	}
	return ans
}

func maxArea2(height []int) int {
	ans, i, j := -1, 0, len(height)-1
	for i < j {
		temp := 0
		if height[i] > height[j] {
			temp = height[j] * (j - i)
			j--
		} else {
			temp = height[i] * (j - i)
			i++
		}
		if temp > ans {
			ans = temp
		}
	}
	return ans
}

func TestMaxArea(t *testing.T) {

	t.Log(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	t.Log(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
