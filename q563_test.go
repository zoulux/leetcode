package leetcode

import (
	. "mutil"
	"testing"
)

var rootTilt *TreeNode

func init() {
	rootTilt = new(TreeNode)
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	return sum(root.Left) + sum(root.Right) + root.Val
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	tilt := sum(root.Left) - sum(root.Right)
	if tilt < 0 {
		tilt = -tilt
	}

	return tilt + findTilt(root.Left) + findTilt(root.Right)
}

func TestFindTilt(t *testing.T) {
	t.Log(findTilt(GenerateTrees(4, 2, 9, 3, 5, -1, 7)))
	t.Log(findTilt(GenerateTrees(21, 7, 14, 1, 1, 2, 2, 3, 3)))
}
