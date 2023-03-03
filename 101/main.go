package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	n := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}

	fmt.Println(isSymmetric(n))

	n = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 2},
		},
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 2},
		},
	}

	fmt.Println(isSymmetric(n))
	n = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
	}

	fmt.Println(isSymmetric(n))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var travel func(left, right *TreeNode) bool
	travel = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		return travel(left.Left, right.Right) && travel(left.Right, right.Left)
	}
	return travel(root.Left, root.Right)
}
