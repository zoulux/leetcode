package main

import . "leetcode"

func main() {

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var travel func(l, r *TreeNode) bool
	travel = func(l, r *TreeNode) bool {
		if l == nil || r == nil {
			return true
		}

		if l.Val != r.Val {
			return false
		}
		return travel(l.Left, r.Right) && travel(l.Right, r.Left)
	}

	return travel(root.Left, root.Right)
}
