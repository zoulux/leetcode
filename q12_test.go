package leetcode

import (
	"sort"
	"strings"
	"testing"
)

var roman = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",

	4:   "IV",
	9:   "IX",
	40:  "XL",
	90:  "XC",
	400: "CD",
	900: "CM",
}

var romanKeys = []int{
	1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
}

func intToRoman(num int) string {

	var sb strings.Builder
	for num > 0 {
		for _, key := range romanKeys {
			if num >= key {
				num -= key
				sb.WriteString(roman[key])
				break
			}
		}
	}

	return sb.String()
}

func TestIntToRoman(t *testing.T) {
	t.Log(intToRoman(3))
	t.Log(intToRoman(4))
	t.Log(intToRoman(9))
	t.Log(intToRoman(58))
	t.Log(intToRoman(1994))
}

func TestSearchInts(t *testing.T) {

	targets := []int{4, 9, 40, 90, 400, 900}
	// sort.SearchStrings()
	t.Log(sort.SearchInts(targets, 10))
	t.Log(sort.SearchInts(targets, 40))

}
