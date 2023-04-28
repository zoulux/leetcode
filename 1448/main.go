package main

import (
	"fmt"
	. "leetcode"
	"math"
)

func main() {
	fmt.Println(goodNodes(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 5},
		},
	}))

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	var ans int

	var travel func(root *TreeNode, maxValue int)
	travel = func(root *TreeNode, maxValue int) {
		if root == nil {
			return
		}

		// 只要当前的值比最大值大/等于 就算好节点
		if root.Val >= maxValue {
			maxValue = root.Val
			ans++
		}

		travel(root.Left, maxValue)
		travel(root.Right, maxValue)
	}

	travel(root, math.MinInt)
	return ans
}
