package main

import (
	. "leetcode"
)

func main() {
	pathSum(NewTreeNode(
		5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1,
	), 22)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {

	ans := make([][]int, 0)
	var travel func(root *TreeNode, target int, arr []int)

	travel = func(root *TreeNode, target int, arr []int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			if target == root.Val {
				ans = append(ans, append([]int{}, append(arr, root.Val)...))
			}
			return
		}

		travel(root.Left, target-root.Val, append(arr, root.Val))
		travel(root.Right, target-root.Val, append(arr, root.Val))
		return
	}

	travel(root, target, []int{})
	return ans
}
