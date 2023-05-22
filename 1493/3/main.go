package main

func main() {

}

func longestSubarray(nums []int) int {

	max := 0
	left, right := 0, 0
	zero := 0
	for right < len(nums) {
		if nums[right] == 0 {
			zero++
		}
		right++

		for zero > 1 {
			if nums[left] == 0 {
				zero--
			}
			left++
		}

		if t := right - left; t > max {
			max = t
		}
	}
	return max
}
