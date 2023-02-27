package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	//输入：root = [4,2,7,1,3,6,9]

	t1 := NewTreeNode(4, 2, 7, 1, 3, 6)
	t2 := mirrorTree(t1)
	fmt.Println(
		t2,
	)
}

var mm map[*TreeNode]bool = make(map[*TreeNode]bool)

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}
