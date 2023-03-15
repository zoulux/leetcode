package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	ans := diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	})

	//ans := diameterOfBinaryTree(&TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//	},
	//})
	fmt.Println(ans)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(nums ...int) int {
		m := nums[0]
		for _, v := range nums {
			if v > m {
				m = v
			}
		}
		return m
	}

	ans := 0
	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		if root.Left == nil && root.Right == nil {
			return 0
		}
		var l, r int
		if root.Left != nil {
			l = travel(root.Left) + 1
		}

		if root.Right != nil {
			r = travel(root.Right) + 1
		}

		ans = max(ans, l+r)
		return max(l, r)
	}

	travel(root)
	return ans
}
