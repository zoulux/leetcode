package leetcode

import (
	"sort"
	"testing"
)

// [-2 -1 -1 1 1 2 2]

// [[-2 -1 1 2] [-2 -1 1 2] [-1 -1 1 1]]
func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)

	n := len(nums)

	ans := [][]int{}

	// [L R left right  ]
	for L := 0; L < n-3; L++ {
		// 左区间相等忽略
		if L > 0 && nums[L] == nums[L-1] {
			continue
		}

		if min := nums[L] + nums[L+1] + nums[L+2] + nums[L+3]; min > target {
			break
		}

		if max := nums[L] + nums[n-1] + nums[n-2] + nums[n-3]; max < target {
			continue
		}

		for R := L + 1; R < n-2; R++ {

			// 右区间相等忽略
			if R > L+1 && nums[R] == nums[R-1] {
				continue
			}

			left, right := R+1, n-1

			if min := nums[L] + nums[R] + nums[left] + nums[left+1]; min > target {
				break
			}

			if max := nums[L] + nums[R] + nums[right] + nums[right-1]; max < target {
				continue
			}

			for left < right {

				if sum := nums[L] + nums[left] + nums[right] + nums[R]; sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					ans = append(ans, []int{
						nums[L],
						nums[R],
						nums[left],
						nums[right],
					})

					for left < right && nums[left] == nums[left+1] {
						left++
					}

					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				}
			}
		}
	}
	return ans
}

func TestFourSum(t *testing.T) {

	t.Log(fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))

	t.Log(fourSum([]int{2, 1, 0, -1}, 2))
	t.Log(fourSum([]int{0, 0, 0, 0}, 0))
	t.Log(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	t.Log(fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
	// [[-5,-4,-3,1]]

	// 	[]
	// -11

}
