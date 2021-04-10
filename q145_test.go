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
func postorderTraversal(root *TreeNode) (ans []int) {

	if root == nil {
		return
	}

	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)

	ans = append(ans, root.Val)

	return append(append(left, right...), ans...)
}
