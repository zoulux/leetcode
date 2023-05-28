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
func isEvenOddTree(root *TreeNode) bool {

	var ans []int
	var travel func(root *TreeNode, level int) bool
	travel = func(root *TreeNode, level int) bool {
		if root == nil {
			return true
		}

		if level%2 == 0 {
			if root.Val%2 != 1 {
				// 奇偶逻辑不匹配
				return false
			}

			if len(ans) <= level {
				ans = append(ans, root.Val) // 为空直接塞进去
			} else {
				if root.Val <= ans[level] {
					return false // 单调性不满足
				}
				ans[level] = root.Val // 每一层只需要记录一个值
			}

		} else {
			if root.Val%2 != 0 {
				// 奇偶逻辑不匹配
				return false
			}

			if len(ans) <= level {
				ans = append(ans, root.Val)
			} else {
				if root.Val >= ans[level] {
					return false // 单调性不满足
				}
				ans[level] = root.Val // 每一层只需要记录一个值
			}
		}

		return travel(root.Left, level+1) && travel(root.Right, level+1) // 左右子树都要为真
	}
	return travel(root, 0)
}
