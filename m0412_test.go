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

func dfs(root *TreeNode, sum int, ans *int) {
	if root == nil {
		return
	}
	sum -= root.Val
	if sum == 0 {
		*ans++
	}
	dfs(root.Left, sum, ans)
	dfs(root.Right, sum, ans)
}

func helper(root *TreeNode, sum int, ans *int) {
	if root == nil {
		return
	}
	dfs(root, sum, ans)
	helper(root.Left, sum, ans)
	helper(root.Right, sum, ans)
}

func pathSum2(root *TreeNode, sum int) int {

	ans := 0
    helper(root, sum, &ans)
    return ans
}

func pathSum(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	ans := 0
	var dfs func(root *TreeNode) []int

	dfs = func(root *TreeNode) []int {

		if root == nil {
			return []int{}
		}

		if root.Val == sum {
			ans++
		}

		if root.Left == nil && root.Right == nil {
			return []int{root.Val}
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		for k, v := range left {

			left[k] = v + root.Val

			if left[k] == sum {
				ans++
			}
		}

		for k, v := range right {
			right[k] = v + root.Val

			if right[k] == sum {
				ans++
			}
		}

		return append(append(left, right...), root.Val)
	}
	dfs(root)
	return ans
}

func TestPathSum(t *testing.T) {
	// tree := GenerateTrees(1, 2)
	// tree := GenerateTrees(5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, 5, 1)
	tree := &TreeNode{
		Val: 1,
	}
	t.Log(pathSum(GenerateTrees(1, 2), 2))
	t.Log(pathSum(GenerateTrees(1, 2), 1))
	t.Log(pathSum(tree, 1))
}
