package main

import (
	"fmt"
)

func main() {

	fmt.Println(hardestWorker(70, [][]int{{36, 3}, {1, 5}, {12, 8}, {25, 9}, {53, 11}, {29, 12}, {52, 14}}))
}

func hardestWorker(n int, logs [][]int) int {

	if len(logs) == 1 {
		return logs[0][0]
	}

	ans := logs[0][1] // 取第一个 耗时
	idx := logs[0][0] // 取第一个 id
	for i := 1; i < len(logs); i++ {

		if d := logs[i][1] - logs[i-1][1]; d > ans || (d == ans && logs[i][0] < idx) {
			ans = d
			idx = logs[i][0] // 如果找到新的则用新的
		}
	}

	return idx // 最小的 id
}
