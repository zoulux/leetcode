package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//：arriveAlice = "10-01", leaveAlice = "10-31", arriveBob = "11-01", leaveBob = "12-31"

	//arriveAlice = "08-15", leaveAlice = "08-18", arriveBob = "08-16", leaveBob = "08-19"
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19"))
	fmt.Println(countDaysTogether("10-01", "10-31", "11-01", "12-31"))
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	var days = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if leaveAlice < arriveBob {
		// 无重合区间
		return 0
	}

	if leaveBob < arriveAlice {
		// 无重合区间
		return 0
	}

	max := func(a, b string) string {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b string) string {
		if a < b {
			return a
		}
		return b
	}

	l := max(arriveAlice, arriveBob) // 重叠左区间
	r := min(leaveAlice, leaveBob)   //重叠右区间
	la := strings.Split(l, "-")
	ra := strings.Split(r, "-")

	lm, _ := strconv.Atoi(la[0]) // 左边界月份
	ld, _ := strconv.Atoi(la[1]) // 左边界日
	rm, _ := strconv.Atoi(ra[0]) // 右边界月份
	rd, _ := strconv.Atoi(ra[1]) // 右边界日

	// 月份相加
	var ans int
	for m := lm; m <= rm; m++ {
		ans += days[m]
	}

	//减去左边界的日
	ans -= ld

	// 减去右边界的日
	ans -= days[rm] - rd

	// 16号到18号，相差 2 天，其实重叠是 2+1
	return ans + 1
}
