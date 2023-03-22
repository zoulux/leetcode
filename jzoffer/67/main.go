package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(strToInt("  1"))
	//fmt.Println(strToInt("  12"))
	//fmt.Println(strToInt("42"))
	//fmt.Println(strToInt("   -42"))
	//fmt.Println(strToInt("4193 with words"))
	//fmt.Println(strToInt("words and 987"))
	fmt.Println(strToInt("91283472332"))
	fmt.Println(strToInt("+-2"))
	fmt.Println(strToInt("2147483648"))
	fmt.Println(strToInt("0-1"))
	fmt.Println(strToInt("01-"))
}

func strToInt(str string) int {

	for len(str) != 0 {
		if str[0] != ' ' {
			break
		}
		str = str[1:]
	}

	var signFlag int = 1
	var num int = 0
	hasSignFlag := false
	hasNumFlag := false

	if len(str) > 0 {
		for _, v := range str {

			if v == '-' {
				if hasNumFlag {
					break
				}

				if hasSignFlag {
					return 0
				}

				signFlag = -1
				hasSignFlag = true
			} else if v == '+' {
				if hasNumFlag {
					break
				}
				if hasSignFlag {
					return 0
				}

				signFlag = 1
				hasSignFlag = true

			} else if v >= '0' && v <= '9' {
				hasNumFlag = true
				num = num*10 + int(v-'0')

				if num*signFlag >= math.MaxInt32 {
					return math.MaxInt32
				}
				if num*signFlag <= math.MinInt32 {
					return math.MinInt32
				}
			} else {
				break
			}
		}
		return num * signFlag
	}
	return 0
}
