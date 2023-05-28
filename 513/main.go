package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	fmt.Println(findBottomLeftValue(&TreeNode{Val: 0}))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {

	var ans []int
	var travel func(root *TreeNode, level int)
	travel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(ans) <= level {
			ans = append(ans, root.Val)
		}

		travel(root.Left, level+1)
		travel(root.Right, level+1)
	}

	travel(root, 0)

	return ans[len(ans)-1]
}
