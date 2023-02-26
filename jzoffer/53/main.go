package main

func main() {

}

func search1(nums []int, target int) int {
	ans := 0
	for _, n := range nums {
		if n == target {
			ans++

		}
	}
	return ans
}

// 5,8,8,8,8,10
// 8
func search(nums []int, target int) int {

	// 最左边的 index
	li := func() int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)>>1 // 7
			if target > nums[mid] {
				left = mid + 1
			} else if target < nums[mid] {
				right = mid - 1
			} else {
				if mid-1 >= 0 && nums[mid-1] == target {
					right = mid - 1
				} else {
					return mid
				}
			}
		}
		return -1
	}()

	// 最右边的 index
	ri := func() int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)>>1 // 7
			if target > nums[mid] {
				left = mid + 1
			} else if target < nums[mid] {
				right = mid - 1
			} else {
				if mid+1 < len(nums) && nums[mid+1] == target {
					left = mid + 1
				} else {
					return mid
				}
			}
		}
		return -1
	}()

	if ri != -1 && li != -1 {
		return ri - li + 1
	}
	return 0
}

func search2(nums []int, target int) int {

	leftBinarySearch := func(nums []int, target int) int {
		left := 0
		right := len(nums) - 1
		for left <= right {
			mid := (left + right) / 2

			if target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			mid = (left + right) / 2
		}
		if left < len(nums) && nums[left] == target {
			return left
		}
		return -1
	}

	rightBinarySearch := func(nums []int, target int) int {
		left := 0
		right := len(nums) - 1
		for left <= right {
			mid := (left + right) / 2
			if target >= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
			mid = (left + right) / 2
		}
		if right >= 0 && nums[right] == target {
			return right
		}
		return -1
	}

	l := leftBinarySearch(nums, target)
	r := rightBinarySearch(nums, target)

	if l != -1 && r != -1 {
		return r - l + 1
	}
	return 0
}
