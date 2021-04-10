package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	tmp := 0

	if root.Val >= low && root.Val <= high {
		tmp = root.Val
	}

	return tmp + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
