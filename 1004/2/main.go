package main

func main() {

}

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	zeroCount := 0
	mx := 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}
		right++

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		if t := right - left; t > mx {
			mx = t
		}
	}
	return mx
}
