package leetcode

import (
	"math"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ll, lll, isEven := len(nums1)+len(nums2), 0, false
	if ll%2 == 1 {
		lll = (ll + 1) / 2
		isEven = false
	} else {
		lll = ll/2 + 1
		isEven = true
	}

	current, left, right := 0, 0, 0
	for current < lll {
		v1, v2 := math.MaxInt64, math.MaxInt64
		if len(nums1) != 0 {
			v1 = nums1[0]
		}

		if len(nums2) != 0 {
			v2 = nums2[0]
		}

		if v1 < v2 {
			left, right, nums1 = right, v1, nums1[1:]
		} else {
			left, right, nums2 = right, v2, nums2[1:]
		}
		current++
	}

	if isEven {
		return float64(left) + float64(right-left)/2.0
	}

	return float64(right)
}

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	t.Log(findMedianSortedArrays([]int{1, 2}, []int{3, 4, 5}))
	t.Log(findMedianSortedArrays([]int{0, 0, 0, 0, 0}, []int{-1, 0, 0, 0, 0, 0, 1}))
	t.Log(findMedianSortedArrays([]int{1}, []int{}))
	t.Log(findMedianSortedArrays([]int{0}, []int{}))
}
