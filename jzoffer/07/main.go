package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	fmt.Println(buildTree([]int{
		3, 9, 20, 15, 7,
	}, []int{
		9, 3, 15, 20, 7,
	}))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	root := TreeNode{Val: preorder[0]}
	index := search(inorder, root.Val)

	// 前序 => 根 + 左子树 + 右子树
	// 中序 => 左子树 + 根 + 右子树
	// 所以 根 + 左子树  == 左子树 + 根
	// 所以找到 index，其实也就确定了在前序和中序中的两个子数组
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return &root
}

func search(nums []int, x int) int {
	for i, v := range nums {
		if v == x {
			return i
		}
	}
	return 0
}
