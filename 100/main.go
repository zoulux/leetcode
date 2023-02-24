package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	n := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isSameTree(n, n))

	n1 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	n2 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isSameTree(n1, n2))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
