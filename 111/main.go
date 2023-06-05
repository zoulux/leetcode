package main

import (
	. "leetcode"
)

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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type item struct {
		root  *TreeNode
		level int
	}
	var q []item

	q = append(q, item{
		root:  root,
		level: 1,
	})

	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		node, level := top.root, top.level
		if node.Left == nil && node.Right == nil {
			return level
		}

		if node.Left != nil {
			q = append(q, item{
				root:  node.Left,
				level: level + 1,
			})
		}

		if node.Right != nil {
			q = append(q, item{
				root:  node.Right,
				level: level + 1,
			})
		}
	}

	return 0
}
