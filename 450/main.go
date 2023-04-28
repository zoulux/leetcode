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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	} else {
		// 右子树的最左节点是最小的
		// 左子树上比当前数值小
		// 将当前的左子树挂到右子树的最小值上面去
		// 返回右子树的头，因为她巧好比 root 大一点点
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		cur.Left = root.Left
		return root.Right
	}
	return root
}
