package leetcode

import (
	"testing"
)

var res []string

var panel = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func findCombination(digits string, index int, s string) {

	if index == len(digits) {
		res = append(res, s)
		return
	}

	d := digits[index]
	panelArr := panel[d]

	for i := 0; i < len(panelArr); i++ {
		findCombination(digits, index+1, s+panelArr[i])
	}
}

func letterCombinations(digits string) (ans []string) {
	if len(digits) == 0 {
		return
	}

	res = []string{}
	findCombination(digits, 0, "")

	return res
}

func TestLetterCombinations(t *testing.T) {
	t.Log(letterCombinations2("23"))
}

func letterCombinations2(digits string) (ret []string) {
	if digits == "" {
		return nil
	}
	n := len(digits)
	var s = make(map[string][]string)
	s["2"] = []string{"a", "b", "c"}
	s["3"] = []string{"d", "e", "f"}
	s["4"] = []string{"g", "h", "i"}
	s["5"] = []string{"j", "k", "l"}
	s["6"] = []string{"m", "n", "o"}
	s["7"] = []string{"p", "q", "r", "s"}
	s["8"] = []string{"t", "u", "v"}
	s["9"] = []string{"w", "x", "y", "z"}

	var generator func(int, string)
	generator = func(i int, cur string) {
		if i == n {
			ret = append(ret, cur)
			return
		}
		for _, a := range s[string(digits[i])] {
			generator(i+1, cur+a)
		}
	}
	generator(0, "")
	return ret
}
