package main

func main() {

}

func findRepeatNumber(nums []int) int {
	arr := make([]int, len(nums)+1)
	for _, n := range nums {
		arr[n]++
		if arr[n] > 1 {
			return n
		}
	}

	return 0
}
