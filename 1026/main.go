package main

import (
	. "leetcode"
)

func main() {

}

func maxAncestorDiff(root *TreeNode) int {
	ans := 0
	var travel func(root *TreeNode, min, max int)
	travel = func(root *TreeNode, min, max int) {
		if root == nil {
			return
		}

		if root.Val < min {
			min = root.Val
		}
		if root.Val > max {
			max = root.Val
		}

		if root.Left == nil && root.Right == nil {
			if d := max - min; d > ans {
				ans = d
			}
		}

		if root.Left != nil {
			travel(root.Left, min, max)
		}

		if root.Right != nil {
			travel(root.Right, min, max)
		}

	}
	travel(root, root.Val, root.Val)
	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiffErr(root *TreeNode) int {

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	maxnum := func(arr ...int) int {
		x := arr[0]
		for i := range arr[1:] {
			if arr[i] > x {
				x = arr[i]
			}
		}
		return x
	}

	var max func(v int, root *TreeNode) int
	max = func(v int, root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			return abs(root.Val - v)
		}

		return maxnum(
			abs(root.Val-v),
			max(v, root.Right),
			max(v, root.Left),
		)
	}

	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		return maxnum(
			max(root.Val, root),
			travel(root.Left),
			travel(root.Right),
		)
	}

	return travel(root)

}
