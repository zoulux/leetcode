package main

import (
	. "leetcode"
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
func replaceValueInTree(root *TreeNode) *TreeNode {

	var values []int
	var travel func(root *TreeNode, level int)
	travel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(values) <= level {
			values = append(values, 0)
		}

		values[level] += root.Val
		travel(root.Left, level+1)
		travel(root.Right, level+1)
	}
	travel(root, 0)

	var replace func(root *TreeNode, level int)
	replace = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			return
		}

		t := values[level+1] // 这一层的和
		if root.Left != nil {
			t -= root.Left.Val
		}

		if root.Right != nil {
			t -= root.Right.Val
		}

		if root.Left != nil {
			root.Left.Val = t
		}

		if root.Right != nil {
			root.Right.Val = t
		}

		replace(root.Left, level+1)
		replace(root.Right, level+1)
	}

	replace(root, 0)
	root.Val = 0
	return root
}
