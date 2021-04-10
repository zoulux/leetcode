package leetcode

import (
	"testing"
)

func inorderTraversal2(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	left := inorderTraversal(root.Left)
	left = append(left, root.Val)
	right := inorderTraversal(root.Right)
	return append(left, right...)
}

func inorderTraversal(root *TreeNode) (ans []int) {

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			dfs(root.Left)
		}

		ans = append(ans, root.Val)

		if root.Right != nil {
			dfs(root.Right)
		}
	}

	dfs(root)

	return ans
}

func TestInorderTraversal(t *testing.T) {

	tree := GenerateTrees(1, -1, 2, 3)

	t.Log(inorderTraversal(tree))

}

