package leetcode

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func numDifferentIntegers2(word string) int {

	mm := map[int]struct{}{}

	num := -1
	for _, v := range word {
		if v >= '0' && v <= '9' {
			if num == -1 {
				num = int(v - '0')
			} else {
				num = num*10 + int(v-'0')
			}

		} else {
			if num != -1 {
				mm[num] = struct{}{}
				num = -1
			}
		}
	}

	if num != -1 {
		mm[num] = struct{}{}
		num = -1
	}

	return len(mm)
}

func numDifferentIntegers(word string) int {
	fmt.Println(strings.FieldsFunc(word, unicode.IsLetter))

	return 0
}

func TestNumDifferentIntegers(t *testing.T) {
	t.Log(numDifferentIntegers("a123bc34d8ef34"))
	t.Log(numDifferentIntegers("leet1234code234"))
	t.Log(numDifferentIntegers("a1b01c001"))

}
