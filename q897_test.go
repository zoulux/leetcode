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
func increasingBST(root *TreeNode) *TreeNode {

	list := []int{}

	var orderTree func(root *TreeNode)
	orderTree = func(root *TreeNode) {
		if root == nil {
			return
		}
		orderTree(root.Left)
		list = append(list, root.Val)
		orderTree(root.Right)
	}

	orderTree(root)

	if len(list) == 0 {
		return nil
	}

	dummyNode := new(TreeNode)
	cur := dummyNode

	for len(list) > 0 {
		cur.Right = &TreeNode{
			Val: list[0],
		}
		list = list[1:]
		cur = cur.Right
	}

	return dummyNode.Right
}

func TestIncreasingBST(t *testing.T) {
	tree := GenerateTrees(5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9)
	t.Log(increasingBST(tree))
}
