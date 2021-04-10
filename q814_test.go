package leetcode

import (
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {

	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		if left == 0 {
			root.Left = nil
		}

		right := dfs(root.Right)

		if right == 0 {
			root.Right = nil
		}

		return left + right + root.Val
	}

	if dfs(root) != 0 {
		return root
	}

	return nil
}

func TestPruneTree(t *testing.T) {
	tree := GenerateTrees(1, 0, 1, 0, 0, 0, 1)
	t.Log(pruneTree(tree))
}
