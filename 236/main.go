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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root

	}

	left := lowestCommonAncestor(root.Left, p, q)  // 找到 p，q，或者到底了
	right := lowestCommonAncestor(root.Left, p, q) // 找到 p，q，或者到底了
	if left == nil {
		// 左边到底了，那就返回右边的
		return right
	}

	if right == nil {
		// 右边到底了，那就返回左边
		return left
	}
	// 如果左右不为空，那就一个是 p，一个是 q，那当前 root 就是最近公众节点
	return root
}
