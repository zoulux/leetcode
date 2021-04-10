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

func sumRootToLeaf(root *TreeNode) int {

	var dfs func(root *TreeNode, sum int) int
	dfs = func(root *TreeNode, sum int) int {

		if root == nil {
			return 0
		}

		sum = sum<<1 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}

		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}

	return dfs(root, 0)
}

func TestSumRootToLeaf(t *testing.T) {
	// [1,0,1,0,1,1,0,1,0,0,0,1,0,0,0,0,1,1,0,null,1,0,null,1,1,1,1,null,0,null,null,null,null,null,null,1,null,0,null,null,null,null,null,0,1,1,0,0,0,0,null,null,null,0,null,null,null,0,null,0,null,null,null,null,1,null,null,0,0,0,null,null,null,1,null,null,null,0,0,null,null,null,null,null,0,null,null,null,null,1,null,null,null,0,1,null,0]
	// tree := GenerateTrees(1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, -1, 1, 0, -1, 1, 1, 1, 1, -1, 0, -1, -1, -1, -1, -1, -1, 1, -1, 0, -1, -1, -1, -1, -1, 0, 1, 1, 0, 0, 0, 0, -1, -1, -1, 0, -1, -1, -1, 0, -1, 0, -1, -1, -1, -1, 1, -1, -1, 0, 0, 0, -1, -1, -1, 1, -1, -1, -1, 0, 0, -1, -1, -1, -1, -1, 0, -1, -1, -1, -1, 1, -1, -1, -1, 0, 1, -1, 0)
	tree := GenerateTrees(1, 0, 1, 0, 1, 0, 1)

	t.Log(sumRootToLeaf(tree))
}
