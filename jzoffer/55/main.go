package main

import . "leetcode"

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
func maxDepth(root *TreeNode) int {
	var travel func(root *TreeNode, level int)
	max := 0
	travel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level > max {
			max = level
		}
		travel(root.Left, level+1)
		travel(root.Right, level+1)
	}
	travel(root, 0)
	return max
}
