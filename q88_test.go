package leetcode

import (
	"fmt"
	"testing"
)

func merge3(nums1 []int, m int, nums2 []int, n int) {
	left := 0
	for i := 0; i < n; i++ {
		target := nums2[i]

		right := m - 1
		for left <= right {
			mid := (right-left)/2 + left
			if nums1[mid] == target {
				// 插入位置
				left = mid + 1
			} else if nums1[mid] > target {
				right = mid - 1
			} else if nums1[mid] < target {
				left = mid + 1
			}
		}

		copy(nums1[left+1:], nums1[left:])
		nums1[left] = target
		m++
	}
}

func TestMerge3(t *testing.T) {

	nums1 := []int{1, 1, 3, 0, 0, 0}
	merge3(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1)
}
