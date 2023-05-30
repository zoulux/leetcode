package main

import (
	"github.com/davecgh/go-spew/spew"
	. "leetcode"
)

func main() {

	spew.Dump(delNodes(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}, []int{
		3,
	}))
	//
	//spew.Dump(delNodes(NewTreeNode(1, 2, 3, 4, 5, 6, 7), []int{
	//	3, 5,
	//}))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	mm := make(map[int]bool)
	for _, v := range toDelete {
		mm[v] = true
	}

	var ans []*TreeNode

	var travel func(root *TreeNode) *TreeNode
	travel = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		root.Left = travel(root.Left)
		root.Right = travel(root.Right)

		if mm[root.Val] {
			if root.Left != nil {
				ans = append(ans, root.Left)
			}

			if root.Right != nil {
				ans = append(ans, root.Right)
			}
			return nil
		}
		return root
	}

	if travel(root) != nil {
		ans = append(ans, root)
	}
	return ans
}
