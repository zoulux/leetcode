package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(1))
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(7))

	// 4
	// 16
	// 1+36=37
	//9 +49 =58
	// 25+64 =89
	// 64+ 81 = 145
	// 1+ 16 +25 = 42
}

func isHappy(n int) bool {

	mm := make(map[int]bool)
	var _isHappy func(n int) bool
	_isHappy = func(n int) bool {
		if n == 1 {
			return true
		}

		if mm[n] {
			return false
		}

		mm[n] = true

		ns := strconv.Itoa(n)

		var ans int
		for _, v := range ns {
			ans += int(v-'0') * int(v-'0')
		}

		return _isHappy(ans)
	}
	return _isHappy(n)
}
