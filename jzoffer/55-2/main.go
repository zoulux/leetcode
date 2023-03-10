package main

import (
	"fmt"
	. "leetcode"
)

//输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
//如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

func main() {
	fmt.Println(isBalanced(NewTreeNode(1, 0, 2, 0, 3)))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)

	if abs(ld-rd) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
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
	travel(root, 1)
	return max
}
