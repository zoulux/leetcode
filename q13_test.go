package leetcode

import (
	"testing"
)

// func romanToInt(s string) int {

// 	m1 := map[byte]int{
// 		'I': 1,
// 		'V': 5,
// 		'X': 10,
// 		'L': 50,
// 		'C': 100,
// 		'D': 500,
// 		'M': 1000,
// 	}

// 	// roman := map[string]int{
// 	// 	"I": 1,
// 	// 	"V": 5,
// 	// 	"X": 10,
// 	// 	"L": 50,
// 	// 	"C": 100,
// 	// 	"D": 500,
// 	// 	"M": 1000,
// 	// }

// 	// m2 := map[string]int{
// 	// 	"IV": 4,
// 	// 	"IX": 9,
// 	// 	"XL": 40,
// 	// 	"XC": 90,
// 	// 	"CD": 400,
// 	// 	"CM": 900,
// 	// }

// 	// n := len(s)

// 	// sum, lastnum := 0, 0
// 	// for i := 0; i < n; i++ {

// 	// 	num := roman[s[i]]
// 	// 	if num < lastnum {
// 	// 		sum -= num
// 	// 	} else {
// 	// 		sum += num
// 	// 	}
// 	// 	lastnum = num

// 	// 	// if i+1 < n {
// 	// 	// 	sii := s[i+1]
// 	// 	// 	ss := si &^ sii
// 	// 	// 	if value, ok := m1[ss]; ok {
// 	// 	// 		sum += value
// 	// 	// 		i++
// 	// 	// 	} else {
// 	// 	// 		sum += m1[si]
// 	// 	// 	}
// 	// 	// } else {
// 	// 	// 	sum += m1[si]
// 	// 	// }
// 	// }
// 	// return sum

// 	if s == "" {
// 		return 0
// 	}
// 	num, lastint, total := 0, 0, 0
// 	// for i := 0; i < len(s); i++ {
// 	// 	char := s[len(s)-(i+1) : len(s)-i]
// 	// 	// fmt.Println(char)

// 	// 	// char := s[i]
// 	// 	num = roman[char]
// 	// 	if num < lastint {
// 	// 		total = total - num
// 	// 	} else {
// 	// 		total = total + num
// 	// 	}
// 	// 	lastint = num
// 	// }

// 	for i := len(s) - 1; i >= 0; i-- {
// 		char := s[i]
// 		num = m1[char]
// 		if num < lastint {
// 			total = total - num
// 		} else {
// 			total = total + num
// 		}
// 		lastint = num
// 	}

// 	// for i := 0; i < len(s); i++ {
// 	// 	char := s[i]
// 	// 	num = m1[char]
// 	// 	if num < lastint {
// 	// 		total = total + num
// 	// 	} else {
// 	// 		total = total - num
// 	// 	}
// 	// 	lastint = num
// 	// }

// 	return total
// }

var m = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,

	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {

	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,

		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	n := len(s)

	sum := 0
	for i := 0; i < n; i++ {
		si := s[i : i+1]
		if i+1 < n {
			ss := s[i : i+2]
			if value, ok := m[ss]; ok {
				sum += value
				i++
			} else {
				sum += m[si]
			}
		} else {
			sum += m[si]
		}
	}
	return sum
}

func TestRomanToInt(t *testing.T) {

	// t.Log(romanToInt("III"))
	// t.Log(romanToInt("IV"))
	// t.Log(romanToInt("IX"))
	t.Log(romanToInt("IXL"))
	// t.Log(romanToInt("LVIII"))
	// t.Log(romanToInt("MCMXCIV"))
	// t.Log(romanToInt("DCXXI"))
}
