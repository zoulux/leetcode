package main

func main() {

}

func insert(intervals [][]int, newInterval []int) [][]int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	newintervals := [][]int{newInterval}

	for i := 0; i < len(intervals); i++ {
		addx := intervals[i][0]
		addy := intervals[i][1]

		topx := newintervals[len(newintervals)-1][0]
		topy := newintervals[len(newintervals)-1][1]

		if addy < topx {
			// 需要单独添加，向前加
			newintervals = append(newintervals, newintervals[len(newintervals)-1])
			newintervals[len(newintervals)-2] = intervals[i]
		} else if addx > topy {
			// 需要单独添加，向后
			newintervals = append(newintervals, intervals[i])
		} else {
			// 有重复
			newintervals[len(newintervals)-1][0] = min(addx, topx)
			newintervals[len(newintervals)-1][1] = max(addy, topy)
		}

	}

	return newintervals
}
