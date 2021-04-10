package leetcode

import (
	"testing"
)

func inArray(root *TreeNode, toDelete []int) bool {
	for _, d := range toDelete {
		if root.Val == d {
			return true
		}
	}
	return false
}

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {

	ans := []*TreeNode{}
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {

		if root == nil {
			return root
		}

		root.Left = dfs(root.Left)
		root.Right = dfs(root.Right)

		if inArray(root, toDelete) {
			if root.Left != nil {
				ans = append(ans, root.Left)
			}

			if root.Right != nil {
				ans = append(ans, root.Right)
			}
			root = nil
		}
		return root
	}

	r := dfs(root)
	if r != nil {
		ans = append(ans, root)
	}
	return ans
}

func TestDelNodes(t *testing.T) {
	tree := GenerateTrees(1, 2, 3, 4, 5, 6, 7)
	res := delNodes(tree, []int{3, 5})
	t.Log(res)
}
