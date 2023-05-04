package main

func main() {

}

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func findPeakElement2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}

	if nums[0] > nums[1] {
		return 0
	}

	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	return -1
}
