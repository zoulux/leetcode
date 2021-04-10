package leetcode

import (
	"strings"
	"testing"
)

func convert(s string, numRows int) string {
	with := calcWith(s, numRows)

	content := make([][]uint8, with)
	ln := len(s)
	for i := 0; i < with; i++ {
		content[i] = make([]uint8, numRows)
	}

	strIdx := 0

	i, j := 0, 0

	for strIdx < ln {
		for t := 0; t < numRows && strIdx < ln; t++ {
			content[i][j] = 1
			content[i][j] = s[strIdx]
			strIdx++
			j++
		}

		i++
		j--

		if j > 0 {
			j--
			for j > 0 && strIdx < ln {
				content[i][j] = s[strIdx]
				i++
				j--
				strIdx++
			}
		}

	}

	var ret strings.Builder
	for j := 0; j < numRows; j++ {
		for i := 0; i < with; i++ {
			c := content[i][j]
			if c != 0 {
				ret.WriteString(string(rune(c)))
			}
		}
	}

	return ret.String()
}

func calcWith(s string, numRows int) int {
	with, n := 0, len(s)
	for n > 0 {
		n = n - numRows
		with++
		for i := 0; i < numRows-2; i++ {
			if n <= 0 {
				break
			}
			n = n - 1
			with++
		}
	}
	return with
}

func TestConvert(t *testing.T) {
	// t.Log(convert("PAYPALISHIRING", 3))
	// t.Log(convert("PAYPALISHIRING", 3))
	// t.Log(calcWith("AB", 1))

	// t.Log(convert("AB", 1))
	t.Log(convert("PAYPALISHIRING", 3))
	// t.Log(convert("A", 1))
	// t.Log("-----")

	// t.Log("AB"[0], reflect.TypeOf("AB"[0]))
	// t.Log(",AB"[0], reflect.TypeOf("AB"[0]))

}
