package leetcode

import (
	. "mutil"
	"testing"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}
	preorder = preorder[1:]

	if len(preorder) > 0 {

		targetJ := 0
		for i, v := range inorder {
			targetJ = i
			if root.Val == v {
				break
			}
		}

		leftInorder, rightInorder := inorder[:targetJ], inorder[targetJ+1:]
		leftPreorder, rightPreorder := preorder[:len(leftInorder)], preorder[len(leftInorder):]

		root.Left = buildTree(leftPreorder, leftInorder)
		root.Right = buildTree(rightPreorder, rightInorder)
	}

	return root
}

func TestBuildTree(t *testing.T) {
	t.Log(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
