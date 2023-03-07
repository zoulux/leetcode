package main

import (
	"fmt"
	. "leetcode"
)

//给定一棵二叉搜索树，请找出其中第 k 大的节点的值。

func main() {
	fmt.Println(kthLargest(NewTreeNode(3, 1, 4, 0, 2), 1))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	arr := make([]int, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		arr = append(arr, root.Val)
		travel(root.Right)
	}

	travel(root)
	return arr[len(arr)-k]
}
