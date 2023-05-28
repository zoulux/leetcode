package main

import (
	. "leetcode"
	"sort"
)

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {

	mm := make(map[int][][2]int)
	var travel func(root *TreeNode, row, col int)
	travel = func(root *TreeNode, row, col int) {
		if root == nil {
			return
		}
		mm[col] = append(mm[col], [2]int{root.Val, row})
		travel(root.Left, row+1, col-1)
		travel(root.Right, row+1, col+1)
	}
	travel(root, 0, 0)

	keys := make([]int, 0)
	for k := range mm {
		keys = append(keys, k)

		sort.Slice(mm[k], func(i, j int) bool {
			if mm[k][i][1] == mm[k][j][1] {
				return mm[k][i][0] < mm[k][j][0] // 行相等，按照数值比较
			}
			return mm[k][i][1] < mm[k][j][1] // 按照按比较
		})
	}

	sort.Ints(keys)
	var ans [][]int
	for _, k := range keys {

		t := make([]int, 0, len(mm[k]))
		for _, v := range mm[k] {
			t = append(t, v[0])
		}

		ans = append(ans, t)
	}
	return ans
}
