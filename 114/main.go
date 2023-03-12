package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	r := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	flatten(r)
	fmt.Println(r)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	var xx *TreeNode
	var travel func(root *TreeNode, flag int) *TreeNode
	travel = func(root *TreeNode, flag int) *TreeNode {
		if root == nil || (root.Left == nil && root.Right == nil) {
			if flag == 1 {
				xx = root
			}
			return root
		}

		l := travel(root.Left, 1)  // 拿到左子树的返回值
		r := travel(root.Right, 2) // // 拿到右子树的返回值
		if l != nil {
			//t := l // 主要是获取左子树的最后一个节点
			//for t.Right != nil {
			//	t = t.Right
			//}
			xx.Right = r   // 左子树的最后一个节点等于右子树
			root.Right = l // 当前根的右子树等于左子树的返回值
		} else {
			root.Right = r // 可能存在左子树为空的情况，那就直接把右子树拿过来拼接
		}
		root.Left = nil // 左子树重置掉
		return root
	}

	travel(root, 0)
}
