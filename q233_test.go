package leetcode

func countDigitOne(n int) int {
	// 23
	// 21
	// 13
	// 12
	// 11
	// 10
	// 01

	if n < 10 {
		return 1
	} else {
		return countDigitOne(n / 10)
	}

}
