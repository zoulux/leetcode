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
func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		//叶子节点
		return targetSum == root.Val
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func TestHasPathSum(t *testing.T) {
	// [5,4,8,11,null,13,4,7,2,null,null,null,1]
	// 22
	// [1,2,3]
	// 5
	// [1,2]
	// 0

	// []
	// -5
	tree := GenerateTrees(-2, -1, -3)
	t.Log(hasPathSum(tree, -5))
}
