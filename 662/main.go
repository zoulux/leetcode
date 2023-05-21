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
func widthOfBinaryTree(root *TreeNode) int {

	var ans int
	mm := make(map[int]int)
	var travel func(root *TreeNode, u int, level int)
	travel = func(root *TreeNode, u int, level int) {
		if root == nil {
			return
		}
		if _, ok := mm[level]; !ok {
			mm[level] = u // 最小
		}
		if d := u - mm[level] + 1; d > ans {
			ans = d
		}
		travel(root.Left, u<<1, level+1)
		travel(root.Right, u<<1+1, level+1)
	}

	travel(root, 1, 0)
	return ans
}
