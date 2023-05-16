package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(rampartDefensiveLine([][]int{{1, 2}, {7, 8}, {10, 13}, {18, 19}, {24, 26}, {29, 31}, {34, 39}, {43, 51}, {52, 57}, {63, 68}, {69, 73}, {75, 76}, {77, 78}, {80, 82}, {85, 91}, {105, 109}, {118, 123}, {130, 133}, {134, 146}, {154, 160}, {164, 165}, {178, 181}, {186, 188}, {192, 193}, {198, 200}, {201, 202}, {205, 206}, {212, 214}, {216, 218}, {221, 222}, {228, 234}, {239, 244}, {245, 246}, {247, 248}, {253, 256}, {259, 260}, {262, 268}, {270, 277}, {281, 282}, {287, 292}, {294, 295}, {302, 303}, {305, 306}, {307, 308}, {313, 316}, {317, 321}, {322, 329}, {330, 331}, {348, 354}, {357, 358}, {360, 364}, {365, 366}, {368, 371}, {375, 378}, {379, 381}, {384, 386}, {387, 388}, {407, 411}, {413, 418}, {419, 420}, {422, 424}, {425, 429}, {441, 443}, {444, 467}, {471, 472}, {476, 477}, {478, 481}, {482, 483}, {488, 489}, {496, 500}, {502, 504}, {506, 507}}))

	//fmt.Println(rampartDefensiveLine([][]int{
	//	{2, 3},
	//	{7, 17},
	//	{29, 33},
	//}))
	//
	//fmt.Println(rampartDefensiveLine([][]int{
	//	{0, 3},
	//	{4, 5},
	//	{7, 9},
	//}))
	//
	fmt.Println(rampartDefensiveLine([][]int{
		{1, 2},
		{5, 8},
		{11, 15},
		{18, 25},
	}))

}

func rampartDefensiveLine(rampart [][]int) int {
	right := math.MaxInt / 2
	for i := 1; i < len(rampart)-1; i++ {
		t := (rampart[i][0] - rampart[i-1][1]) + (rampart[i+1][0] - rampart[i][1])
		if t < right {
			right = t
		}
	}

	var dfs func(mm map[[2]int]bool, k int, idx int, left int) bool

	dfs = func(mm map[[2]int]bool, k, idx int, left int) (result bool) {
		if left > k {
			return false
		}

		if rampart[idx-1][1] > rampart[idx][0] {
			return false
		}

		if idx == len(rampart)-1 {
			return true
		}

		if v, ok := mm[[2]int{idx, left}]; ok {
			return v
		}

		defer func() {
			if result {
				mm[[2]int{idx, left}] = result
			}
		}()

		if rampart[idx-1][1] <= rampart[idx][0]-left {

			rampart[idx][0] -= left
			rampart[idx][1] += k - left
			ans := dfs(mm, k, idx+1, 0)
			rampart[idx][0] += left
			rampart[idx][1] -= k - left
			if ans {
				return true
			}
		}

		if left+1 <= k && dfs(mm, k, idx, left+1) {
			return true
		}

		return false
	}

	left := 1
	for left <= right {
		mid := (left + right) / 2
		if dfs(map[[2]int]bool{}, mid, 1, 0) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}
