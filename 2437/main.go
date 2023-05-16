package main

import (
	"fmt"
)

func main() {

	fmt.Println(countTime("?2:16") == 3)
	fmt.Println(countTime("?5:00") == 2)
	fmt.Println(countTime("0?:0?") == 100)
	fmt.Println(countTime("??:??") == 1440)

}

func countTime(time string) int {
	h := 1
	if time[0] == '?' {
		if time[1] == '?' {
			h *= 24
		} else if time[1] <= '3' {
			h *= 3
		} else {
			h *= 2
		}
	}

	if time[1] == '?' {
		if time[0] <= '1' {
			h *= 10
		} else if time[0] == '2' {
			h *= 4
		}
	}

	if time[3] == '?' {
		h *= 6
	}

	if time[4] == '?' {
		h *= 10
	}

	return h
}
