package leetcode

import . "mutil"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) (ans []int) {

	if root == nil {
		return
	}

	ans = append(ans, root.Val)

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	return append(append(ans, left...), right...)
}
