package leetcode

import (
	. "mutil"
	"testing"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func TestLowestCommonAncestor(t *testing.T) {
	tree := GenerateTrees(3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4)
	lowestCommonAncestor(tree, tree.Left, tree.Right)
}
