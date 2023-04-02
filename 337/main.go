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
func rob2(root *TreeNode) int {
	type Result struct {
		YES int // 偷
		NO  int // 不偷
	}
	var travel func(root *TreeNode) Result
	travel = func(root *TreeNode) Result {
		if root == nil {
			return Result{}
		}

		left := travel(root.Left)
		right := travel(root.Right)

		return Result{
			YES: root.Val + left.NO + right.NO,
			NO:  max(left.YES, left.NO) + max(right.YES, right.NO),
		}
	}
	r := travel(root)
	return max(r.YES, r.NO)
}

func rob(root *TreeNode) int {
	type Key struct {
		root *TreeNode
		flag int
	}
	memo := make(map[Key]int)

	var travel func(root *TreeNode, flag int) int
	travel = func(root *TreeNode, flag int) int {
		if root == nil {
			return 0
		}

		k := Key{
			root: root,
			flag: flag,
		}
		if result, ok := memo[k]; ok {
			return result
		}

		var result int

		if flag == 1 {
			// 偷
			result = root.Val + travel(root.Left, 0) + travel(root.Right, 0)
		} else {
			// 不偷
			result = max(travel(root.Left, 1), travel(root.Left, 0)) + max(travel(root.Right, 1), travel(root.Right, 0))
		}
		memo[k] = result
		return result
	}
	return max(travel(root, 0), travel(root, 1))
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
