package main

import (
	"fmt"
)

func main() {

	fmt.Println(predictPartyVictory("RD"))
	//fmt.Println(predictPartyVictory("DDRRR"))

}

func predictPartyVictory(senate string) string {
	bs := []byte(senate)
	r, d := true, true
	flag := 0
	for r && d {
		r, d = false, false
		for i := 0; i < len(bs); i++ {
			switch bs[i] {
			case 'D':
				if flag < 0 {
					bs[i] = 0
				} else {
					d = true
				}
				flag++
			case 'R':
				if flag > 0 {
					bs[i] = 0
				} else {
					r = true
				}
				flag--
			}
		}
	}
	if r {
		return "Radiant"
	}
	return "Dire"
}
