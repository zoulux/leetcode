package leetcode

func singleNumber(nums []int) int {

	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0
}

func singleNumber2(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}

	return a

}
