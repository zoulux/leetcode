package leetcode

import (
	"sort"
	"testing"
)

func threeSum2(nums []int) [][]int {
	left, right := 0, len(nums)-1

	sort.Ints(nums)
	cache := map[int]struct{}{}
	cacheadd := map[int]struct{}{}

	ans := [][]int{}
	for left < right {
		cur := (left + right) / 2
		temp := 0

		for cur > left && cur < right {

			k := nums[left] &^ nums[cur] &^ nums[right]
			if _, ok := cacheadd[k]; ok {
				continue
			}

			temp = nums[left] + nums[cur] + nums[right]
			if temp == 0 {
				break
			} else if temp > 0 {
				cur--
			} else if temp < 0 {
				cur++
			}
			cacheadd[k] = struct{}{}
		}

		if temp == 0 {
			if cur != left && cur != right {
				k := nums[left] &^ nums[cur] &^ nums[right]
				if _, ok := cache[k]; !ok {
					ans = append(ans, []int{nums[left], nums[cur], nums[right]})
					cache[k] = struct{}{}
				}
			}
			left++
		} else if temp > 0 {
			right--
		} else if temp < 0 {
			left++
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	ans := [][]int{}

	for i := 0; i < len(nums); i++ {

		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			//去除重复元素
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			tmp := nums[i] + nums[left] + nums[right]

			if tmp == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--

			} else if tmp < 0 {
				left++
			} else {
				right--
			}
		}

	}
	return ans
}

func threeSum3(nums []int) [][]int {

	sort.Ints(nums)
	ans := [][]int{}

	left, right := 0, len(nums)-1

	for left < right {
		cur := (left + right) / 2

		tmp := 0
		for cur > left && cur < right {
			tmp = nums[left] + nums[cur] + nums[right]
			if tmp == 0 {
				ans = append(ans, []int{nums[left], nums[cur], nums[right]})
				break
			} else if tmp < 0 {
				cur = (cur + right) / 2
			} else {
				cur = (cur + left) / 2
			}
		}

		if tmp < 0 {
			left++
		} else {
			right--
		}

	}

	return ans
}

func TestThreeSum(t *testing.T) {
	// t.Log(threeSum3([]int{1, 2, -2, 1}))
	t.Log(threeSum3([]int{-1, 0, 1, 2, -1, -4}))
}
