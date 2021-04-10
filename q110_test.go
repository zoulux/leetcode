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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var height func(root *TreeNode) int

	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		return Max(height(root.Left), height(root.Right)) + 1
	}

	lheight, rheight := height(root.Left), height(root.Right)

	return abs(lheight-rheight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func TestIsBalanced(t *testing.T) {
	tree := GenerateTrees(1, 2, 2, 3, -1, -1, 3, 4, -1, -1, 4)
	// tree := GenerateTrees(1, 2)
	// tree := GenerateTrees(1, 2, 2, 3, 3, -1, -1, 4, 4)
	// tree := GenerateTrees(3, 9, 20, -1, -1, 15, 7)
	t.Log(isBalanced2(tree))
	t.Log(isBalanced(tree))
}

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
