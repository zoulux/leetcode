package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(countSeniors([]string{
		"9751302862F0693", "3888560693F7262",
		"5485983835F0649", "2580974299F6042",
		"9976672161M6561", "0234451011F8013",
		"4294552179O6482",
	}))

}

func countSeniors(details []string) int {

	var ans int
	for _, d := range details {
		if v, err := strconv.Atoi(d[11:13]); err == nil && v > 60 {
			ans++
		}
	}

	return ans
}
