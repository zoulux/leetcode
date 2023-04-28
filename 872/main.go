package main

import . "leetcode"

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	var travel func(root *TreeNode) []int
	travel = func(root *TreeNode) []int {
		if root == nil {
			return nil
		}
		if root.Left == nil && root.Right == nil {

			return []int{root.Val}
		}

		return append(travel(root.Left), travel(root.Right)...)
	}

	t1 := travel(root1)
	t2 := travel(root2)
	if len(t1) != len(t2) {
		return false
	}

	for i := 0; i < len(t1); i++ {
		if t1[i] != t2[i] {
			return false
		}
	}

	return true
}
