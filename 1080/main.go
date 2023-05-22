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
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	var travel func(root *TreeNode, sum int) bool // 是否要删除
	travel = func(root *TreeNode, sum int) bool {
		if root == nil {
			return false
		}
		if root.Left == nil && root.Right == nil {
			// 叶子节点
			return sum+root.Val < limit
		}

		// 观察左节点是否要被删除
		lb := travel(root.Left, sum+root.Val)

		// 观察右节点是否要被删除
		rb := travel(root.Right, sum+root.Val)

		if lb {
			root.Left = nil //需要删除左节点
		}

		if rb {
			root.Right = nil // 需要删除右节点
		}

		return lb && rb // 左右节点如果都删除，当前节点就不需要了
	}
	if travel(root, 0) { // 根节点如果也需要删除，那就返回 nil
		return nil
	}
	return root // 返回原来的根节点
}

func sufficientSubset2(root *TreeNode, limit int) *TreeNode {
	var travel func(root *TreeNode, sum int) bool // 是否要删除
	travel = func(root *TreeNode, sum int) bool {
		if root == nil {
			return true
		}

		if root.Left == nil && root.Right == nil {
			// 叶子节点
			return sum+root.Val < limit
		}

		lb, rb := true, true // 当有一个节点不用被删除的时候,父节点就不用被删除
		if root.Left != nil {
			lb = travel(root.Left, sum+root.Val) // 观察左节点是否要被删除
			if lb {
				root.Left = nil //需要删除左节点
			}
		}

		if root.Right != nil {
			rb = travel(root.Right, sum+root.Val) // 观察右节点是否要被删除
			if rb {
				root.Right = nil // 需要删除右节点
			}
		}

		return lb && rb // 左右节点如果都删除，当前节点就不需要了
	}
	if travel(root, 0) { // 根节点如果也需要删除，那就返回 nil
		return nil
	}
	return root // 返回原来的根节点
}
