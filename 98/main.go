package main

import (
	"fmt"
	. "leetcode"
	"math"
)

//98. 验证二叉搜索树

//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//有效 二叉搜索树定义如下：
//节点的左子树只包含 小于 当前节点的数。
//节点的右子树只包含 大于 当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。

func main() {
	n := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(isValidBST(n))

	n = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(isValidBST(n))

}

func isValidBST(root *TreeNode) bool {
	var _isValidBST func(root *TreeNode, min, max int) bool
	_isValidBST = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}

		// 当前值不在限定访问内
		if min >= root.Val || root.Val >= max {
			return false
		}

		// 左子树不为空，进一步判断左子树，并且可以限定左子树的范围
		if root.Left != nil && !_isValidBST(root.Left, min, root.Val) {
			return false
		}

		// 右子树同理
		if root.Right != nil && !_isValidBST(root.Right, root.Val, max) {
			return false
		}
		return true
	}

	return _isValidBST(root, math.MinInt, math.MaxInt)
}
