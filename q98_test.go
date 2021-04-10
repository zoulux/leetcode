package leetcode

import (
	. "mutil"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {

	// 用时记录左子树的最大值和右子树的最小值
	var helper func(root, min, max *TreeNode) bool
	helper = func(root, min, max *TreeNode) bool {

		if root == nil {
			return true
		}

		// 当前节点 必须在 min 和 max 之前
		if min != nil && root.Val <= min.Val {
			return false
		}

		if max != nil && root.Val >= max.Val {
			return false
		}

		return helper(root.Left, min, root) && helper(root.Left, root, max)

	}
	return helper(root, nil, nil)
}
